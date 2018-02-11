package drawer

import (
	"image"
	"image/color"
	"image/draw"
	"math"
	"sync"
)

const (
	FillupNearestNeighbor FillupMode = iota
	FillupNearest         FillupMode = iota
	FillupGausian         FillupMode = iota
)

type FillupMode uint8
type Fillup struct {
	src  *image.RGBA
	Mode FillupMode
	//

}

func NewFillup(img image.Image, mode FillupMode) *Fillup {
	var src *image.RGBA
	var ok bool
	if src, ok = img.(*image.RGBA); !ok {
		src = image.NewRGBA(img.Bounds())
		draw.Draw(src, src.Rect, img, image.ZP, draw.Src)
	}
	return &Fillup{
		src:  src,
		Mode: mode,
	}
}
func (s Fillup) ColorModel() color.Model {
	return s.src.ColorModel()
}

func (s Fillup) Bounds() image.Rectangle {
	return s.src.Bounds()
}

func (s Fillup) At(x, y int) color.Color {
	return s.src.At(x, y)
}

func (s Fillup) Draw(dst draw.Image) {
	var dstSize = dst.Bounds().Size()
	var pix, stride, startx, starty = startEditPix(dst)
	defer endEditPix(dst, pix, stride, startx, starty)
	//
	switch s.Mode {
	case FillupNearestNeighbor:
		resizeNearestNeighbor(pix, stride, s.src.Pix, s.src.Stride, dstSize.X, dstSize.Y)
	case FillupNearest:
		resize(pix, stride, s.src.Pix, s.src.Stride, 1., nearest)
	default:
		fallthrough
	case FillupGausian:
		resize(pix, stride, s.src.Pix, s.src.Stride, 1., gausian)
		//resizeminGoroutine(pix, stride, s.src.Pix, s.src.Stride, 1., gausian)
	}
}
func gausian(x float64) float64 {
	x = math.Abs(x)
	if x < 1.0 {
		exp := 2.0
		x *= 2.0
		y := math.Pow(0.5, math.Pow(x, exp))
		base := math.Pow(0.5, math.Pow(2, exp))
		return (y - base) / (1 - base)
	}
	return 0
}
func nearest(x float64) float64 {
	x = math.Abs(x)
	if x < 1.0 {
		return 1.0 - x
	}
	return 0
}
func filterRange(v, length int, delta, radius float64) (val float64, start, end int) {
	val = float64(v) * delta
	if val < 0 {
		val = 0
	} else if val >= float64(length) {
		val = float64(length)
	}
	start, end = int(val-radius+0.5), int(val+radius)
	if start < 0 {
		start = 0
	}
	if end >= length {
		end = length
	}
	return
}
func resize(dst []uint8, dstStride int, src []uint8, srcStride int, support float64, fn func(x float64) float64) {
	var dstw, dsth = sizeEidtPix(len(dst), dstStride)
	var srcw, srch = sizeEidtPix(len(src), srcStride)
	var deltaH = float64(srcw) / float64(dstw)
	var scaleH = math.Max(deltaH, 1.0)
	var radiusH = math.Ceil(scaleH * support)
	var deltaV = float64(srch) / float64(dsth)
	var scaleV = math.Max(deltaV, 1.0)
	var radiusV = math.Ceil(scaleV * support)
	var wg = new(sync.WaitGroup)
	//
	var memorizationW = make([]struct {
		val        float64
		start, end int
	}, dstw)
	for x := 0; x < dstw; x++ {
		xsrc, hstart, hend := filterRange(x, srcw, deltaH, radiusH)
		memorizationW[x] = struct {
			val        float64
			start, end int
		}{
			val:   xsrc,
			start: hstart,
			end:   hend,
		}
	}
	var memorizationH = make([]struct {
		val        float64
		start, end int
	}, dstw)
	for y := 0; y < dsth; y++ {
		ysrc, vstart, vend := filterRange(y, srch, deltaV, radiusV)
		memorizationH[y] = struct {
			val        float64
			start, end int
		}{
			val:   ysrc,
			start: vstart,
			end:   vend,
		}
	}
	//
	for x := 0; x < dstw; x++ {
		wg.Add(1)
		go func(innerX int) {
			for y := 0; y < dsth; y++ {
				var r, g, b, a float64
				var sum float64
				//
				h, v := memorizationW[innerX], memorizationH[y]
				// H : pixel evaluate
				for kx := h.start; kx < h.end; kx++ {
					srcoffset := offsetEditPix(kx, int(v.val), srcStride)
					normal := (float64(kx) - h.val) / scaleH
					res := fn(normal)
					// normalized r, g, b, a and sum
					r += float64(src[srcoffset+R]) * res
					g += float64(src[srcoffset+G]) * res
					b += float64(src[srcoffset+B]) * res
					a += float64(src[srcoffset+A]) * res
					sum += res
				}
				// V : pixel evaluate
				for ky := v.start; ky < v.end; ky++ {
					srcoffset := offsetEditPix(int(h.val), ky, srcStride)
					normal := (float64(ky) - v.val) / scaleV
					res := fn(normal)
					// normalized r, g, b, a and sum
					r += float64(src[srcoffset+R]) * res
					g += float64(src[srcoffset+G]) * res
					b += float64(src[srcoffset+B]) * res
					a += float64(src[srcoffset+A]) * res
					sum += res
				}
				// pixel set
				dstoffset := offsetEditPix(innerX, y, dstStride)
				dst[dstoffset+R] = clamp((r/sum)+0.5, 0, 255)
				dst[dstoffset+G] = clamp((g/sum)+0.5, 0, 255)
				dst[dstoffset+B] = clamp((b/sum)+0.5, 0, 255)
				dst[dstoffset+A] = clamp((a/sum)+0.5, 0, 255)
			}
			wg.Done()
		}(x)
	}
	wg.Wait()
}
func resizeNearestNeighbor(dst []uint8, dstStride int, src []uint8, srcStride int, width, height int) {
	var srcw, srch = sizeEidtPix(len(src), srcStride)
	var dx = float32(srcw) / float32(width)
	var dy = float32(srch) / float32(height)
	var wg = new(sync.WaitGroup)

	for y := 0; y < height; y++ {
		wg.Add(1)
		go func(yin int) {
			for x := 0; x < width; x++ {
				dstIdx := offsetEditPix(
					x,
					yin,
					dstStride,
				)
				srcIdx := offsetEditPix(
					int((float32(x)+0.5)*dx),
					int((float32(yin)+0.5)*dy),
					srcStride,
				)
				// R, G, B, A
				dst[dstIdx+R] = src[srcIdx+0]
				dst[dstIdx+G] = src[srcIdx+1]
				dst[dstIdx+B] = src[srcIdx+2]
				dst[dstIdx+A] = src[srcIdx+3]
			}
			wg.Done()
		}(y)
	}
	wg.Wait()
}