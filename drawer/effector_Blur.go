package drawer

import (
	"image/draw"
	"math"
	"sync"
)

const (
	BlurBox BlurMode = iota
	BlurGaussian BlurMode = iota
)
type BlurMode uint8

type Blur struct {
	Radius float64
	KeepAlpha bool
	Mode BlurMode
}

func NewBlur(Radius float64, KeepAlpha bool, Mode BlurMode) *Blur {

	return &Blur{
		Radius:Radius,
		KeepAlpha:KeepAlpha,
		Mode:Mode,
	}
}
func (s Blur) Draw(target draw.Image) {
	var targetPix, targetStride, targetX, targetY = startEditPix(target)
	defer endEditPix(target, targetPix, targetStride, targetX, targetY)
	switch s.Mode {
	default:
		fallthrough
	case BlurBox:
		if s.KeepAlpha{
			boxBlurKeepAlpha(targetPix, targetStride, s.Radius)
		}else {
			boxBlur(targetPix, targetStride, s.Radius)
		}
	case BlurGaussian:
		if s.KeepAlpha{
			boxGaussianKeepAlpha(targetPix, targetStride, s.Radius)
		}else {
			boxGaussian(targetPix, targetStride, s.Radius)
		}

	}
}

func boxBlur(src []uint8, srcStride int, radius float64)  {
	var length = int(math.Ceil(radius)) + 1
	var rad = length/2
	var mat = NewMatrix(length, length)
	mat.Clear(1)
	mat = mat.Normal()
	var pad, padStride = paddingExtend(src, srcStride,rad, rad)
	var wg = new(sync.WaitGroup)
	srcW, srcH := sizeEidtPix(len(src), srcStride)

	for y := 0; y < srcH; y++ {
		wg.Add(1)
		go func(innerY int) {
			for x := 0; x < srcW; x++ {
				var r, g, b, a float64
				for ky := 0; ky < length; ky++ {
					iy := innerY + ky
					for kx := 0; kx < length; kx++ {
						ix := x + kx
						kval := mat[kx][ky]
						offset := offsetEditPix(ix, iy, padStride)
						r += float64(pad[offset+R]) * kval
						g += float64(pad[offset+G]) * kval
						b += float64(pad[offset+B]) * kval
						a += float64(pad[offset+A]) * kval
					}
				}
				offset := offsetEditPix(x, innerY, srcStride)
				src[offset+R] = uint8(clamp(r, 0, 255))
				src[offset+G] = uint8(clamp(g, 0, 255))
				src[offset+B] = uint8(clamp(b, 0, 255))
				src[offset+A] = uint8(clamp(a, 0, 255))
			}
			wg.Done()
		}(y)
	}
	wg.Wait()
}
func boxBlurKeepAlpha(src []uint8, srcStride int, radius float64)  {
	var length = int(math.Ceil(radius)) + 1
	var rad = length/2
	var mat = NewMatrix(length, length)
	mat.Clear(1)
	mat = mat.Normal()
	var pad, padStride = paddingExtend(src, srcStride,rad, rad)
	var wg = new(sync.WaitGroup)
	srcW, srcH := sizeEidtPix(len(src), srcStride)

	for y := 0; y < srcH; y++ {
		wg.Add(1)
		go func(innerY int) {
			for x := 0; x < srcW; x++ {
				var r, g, b float64
				for ky := 0; ky < length; ky++ {
					iy := innerY + ky
					for kx := 0; kx < length; kx++ {
						ix := x + kx
						kval := mat[kx][ky]
						offset := offsetEditPix(ix, iy, padStride)
						r += float64(pad[offset+R]) * kval
						g += float64(pad[offset+G]) * kval
						b += float64(pad[offset+B]) * kval
					}
				}
				offset := offsetEditPix(x, innerY, srcStride)
				src[offset+R] = uint8(clamp(r, 0, 255))
				src[offset+G] = uint8(clamp(g, 0, 255))
				src[offset+B] = uint8(clamp(b, 0, 255))
			}
			wg.Done()
		}(y)
	}
	wg.Wait()
}
func gaussian(x, y, sigma float64) float64 {
	return math.Exp(-(x*x/sigma + y*y/sigma))
}
func boxGaussian(src []uint8, srcStride int, radius float64)  {
	var length = int(math.Ceil(radius)) + 1
	var rad = length/2
	var mat = NewMatrix(length, length)
	for x := 0; x < length; x++ {
		for y := 0; y < length; y++ {
			mat[x][y] = gaussian(float64(x)-radius, float64(y)-radius, 4*radius)
		}
	}
	var pad, padStride = paddingExtend(src, srcStride,rad, rad)
	var wg = new(sync.WaitGroup)
	srcW, srcH := sizeEidtPix(len(src), srcStride)

	for y := 0; y < srcH; y++ {
		wg.Add(1)
		go func(innerY int) {
			for x := 0; x < srcW; x++ {
				var r, g, b, a float64
				for ky := 0; ky < length; ky++ {
					iy := innerY + ky
					for kx := 0; kx < length; kx++ {
						ix := x + kx
						kval := mat[kx][ky]
						offset := offsetEditPix(ix, iy, padStride)
						r += float64(pad[offset+R]) * kval
						g += float64(pad[offset+G]) * kval
						b += float64(pad[offset+B]) * kval
						a += float64(pad[offset+A]) * kval
					}
				}
				offset := offsetEditPix(x, innerY, srcStride)
				src[offset+R] = uint8(clamp(r, 0, 255))
				src[offset+G] = uint8(clamp(g, 0, 255))
				src[offset+B] = uint8(clamp(b, 0, 255))
				src[offset+A] = uint8(clamp(a, 0, 255))
			}
			wg.Done()
		}(y)
	}
	wg.Wait()
}
func boxGaussianKeepAlpha(src []uint8, srcStride int, radius float64)  {
	var length = int(math.Ceil(radius)) + 1
	var rad = length/2
	var mat = NewMatrix(length, length)
	for x := 0; x < length; x++ {
		for y := 0; y < length; y++ {
			mat[x][y] = gaussian(float64(x)-radius, float64(y)-radius, 4*radius)
		}
	}
	var pad, padStride = paddingExtend(src, srcStride,rad, rad)
	var wg = new(sync.WaitGroup)
	srcW, srcH := sizeEidtPix(len(src), srcStride)

	for y := 0; y < srcH; y++ {
		wg.Add(1)
		go func(innerY int) {
			for x := 0; x < srcW; x++ {
				var r, g, b float64
				for ky := 0; ky < length; ky++ {
					iy := innerY + ky
					for kx := 0; kx < length; kx++ {
						ix := x + kx
						kval := mat[kx][ky]
						offset := offsetEditPix(ix, iy, padStride)
						r += float64(pad[offset+R]) * kval
						g += float64(pad[offset+G]) * kval
						b += float64(pad[offset+B]) * kval
					}
				}
				offset := offsetEditPix(x, innerY, srcStride)
				src[offset+R] = uint8(clamp(r, 0, 255))
				src[offset+G] = uint8(clamp(g, 0, 255))
				src[offset+B] = uint8(clamp(b, 0, 255))
			}
			wg.Done()
		}(y)
	}
	wg.Wait()
}