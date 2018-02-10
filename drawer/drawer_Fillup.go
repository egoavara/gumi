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
	var pix, stride = startEditPix(dst)
	defer endEditPix(dst, pix, stride)
	//
	switch s.Mode {
	case FillupNearestNeighbor:
		nearestNeighbor(pix, stride, s.src.Pix, s.src.Stride, dstSize.X, dstSize.Y)
	case FillupNearest:
		resizeHorizontal(pix, stride, s.src.Pix, s.src.Stride, 1., nearest)
		resizeVertical(pix, stride, s.src.Pix, s.src.Stride, 1., nearest)
	default:
		fallthrough
	case FillupGausian:
		resizeHorizontal(pix, stride, s.src.Pix, s.src.Stride, 1., gausian)
		resizeVertical(pix, stride, s.src.Pix, s.src.Stride, 1., gausian)

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
	val = (float64(v)+0.5)*delta - 0.5
	start, end = int(val - radius + 0.5), int(val + radius)
	if start < 0 {
		start = 0
	}
	if end >= length {
		end = length - 1
	}
	return
}
func resizeHorizontal(dst []uint8, dstStride int, src []uint8, srcStride int, support float64, fn func(x float64) float64) {
	var dstw, dsth = sizeEidtPix(len(dst), dstStride)
	var srcw, _ = sizeEidtPix(len(src), srcStride)
	var deltaH = float64(srcw) / float64(dstw)
	var scaleH = math.Max(deltaH, 1.0)
	var radiusH = math.Ceil(scaleH * support)
	//
	for x := 0; x < dstw; x++ {
		// H : index range setup
		hval, hstart, hend := filterRange(x, dstw, deltaH, radiusH)
		for y := 0; y < dsth; y++ {
			// H : pixel evaluate
			var r, g, b, a float64
			var sum float64
			for kx := hstart; kx <= hend; kx++ {
				srcoffset := offsetEditPix(kx, y, srcStride)
				normal := (float64(kx) - hval) / scaleH
				res := fn(normal)
				// normalized r, g, b, a and sum
				r += float64(src[srcoffset+RED]) * res
				g += float64(src[srcoffset+GREEN]) * res
				b += float64(src[srcoffset+BLUE]) * res
				a += float64(src[srcoffset+ALPHA]) * res
				sum += res
			}
			// H : pixel set
			dstoffset := offsetEditPix(x, y, dstStride)
			dst[dstoffset+RED] = clamp((r/sum)+0.5, 0, 255)
			dst[dstoffset+GREEN] = clamp((g/sum)+0.5, 0, 255)
			dst[dstoffset+BLUE] = clamp((b/sum)+0.5, 0, 255)
			dst[dstoffset+ALPHA] = clamp((a/sum)+0.5, 0, 255)
		}
	}

}
func resizeVertical(dst []uint8, dstStride int, src []uint8, srcStride int, support float64, fn func(x float64) float64) {
	var dstw, dsth = sizeEidtPix(len(dst), dstStride)
	var _, srch = sizeEidtPix(len(src), srcStride)
	var deltaV = float64(srch) / float64(dsth)
	var scaleV = math.Max(deltaV, 1.0)
	var radiusV = math.Ceil(scaleV * support)
	//
	for y := 0; y < dsth; y++ {
		// V : index range setup
		vval, vstart, vend := filterRange(y, dsth, deltaV, radiusV)
		for x := 0; x < dstw; x++ {
			var r, g, b, a float64
			var sum float64
			// V : pixel evaluate
			for ky := vstart; ky <= vend; ky++ {
				srcoffset := offsetEditPix(ky, y, srcStride)
				normal := (float64(ky) - vval) / scaleV
				res := fn(normal)
				// normalized r, g, b, a and sum
				r += float64(src[srcoffset+RED]) * res
				g += float64(src[srcoffset+GREEN]) * res
				b += float64(src[srcoffset+BLUE]) * res
				a += float64(src[srcoffset+ALPHA]) * res
				sum += res
			}
			// V : pixel set
			dstoffset := offsetEditPix(x, y, dstStride)
			dst[dstoffset+RED] = clamp((r/sum)+0.5, 0, 255)
			dst[dstoffset+GREEN] = clamp((g/sum)+0.5, 0, 255)
			dst[dstoffset+BLUE] = clamp((b/sum)+0.5, 0, 255)
			dst[dstoffset+ALPHA] = clamp((a/sum)+0.5, 0, 255)
		}

	}
}
func nearestNeighbor(dst []uint8, dstStride int, src []uint8, srcStride int, width, height int) {
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
				dst[dstIdx+RED] = src[srcIdx+0]
				dst[dstIdx+GREEN] = src[srcIdx+1]
				dst[dstIdx+BLUE] = src[srcIdx+2]
				dst[dstIdx+ALPHA] = src[srcIdx+3]
			}
			wg.Done()
		}(y)
	}
	wg.Wait()
}
