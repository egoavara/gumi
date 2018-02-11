package drawer

import (
	"image"
	"image/color"
	"image/draw"
)

const (
	R           = 0
	G           = 1
	B           = 2
	A           = 3
	STRIDECOUNT = 4
)

func startEditPix(src image.Image) (pix []uint8, stride, startX, startY int) {
	if v, ok := src.(*image.RGBA); ok {
		return v.Pix, v.Stride, v.Rect.Min.X, v.Rect.Min.Y
	}
	var imageSize = src.Bounds().Size()
	stride = imageSize.X * STRIDECOUNT
	var pixlen = stride * imageSize.Y
	startX = src.Bounds().Min.X
	startY = src.Bounds().Min.Y

	pix = make([]uint8, pixlen)
	// src Copy to pix
	for x := 0; x < imageSize.X; x++ {
		for y := 0; y < imageSize.Y; y++ {
			r, g, b, a := src.At(x, y).RGBA()
			offset := offsetEditPix(x, y, stride)
			pix[offset+R] = uint8(r >> 8)
			pix[offset+G] = uint8(g >> 8)
			pix[offset+B] = uint8(b >> 8)
			pix[offset+A] = uint8(a >> 8)
		}
	}
	return
}
func endEditPix(dst draw.Image, pix []uint8, stride int, startX, startY int) {
	if _, ok := dst.(*image.RGBA); ok {
		return
	}
	var w, h = sizeEidtPix(len(pix), stride)
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			offset := offsetEditPix(x, y, stride)
			dst.Set(startX+x, startY+y, color.RGBA{
				pix[offset+R],
				pix[offset+G],
				pix[offset+B],
				pix[offset+A],
			})
		}
	}
}

//
func createEditPix(w, h int) (pix []uint8, stride int) {
	stride = w * STRIDECOUNT
	pixlen := stride * h
	pix = make([]uint8, pixlen)
	return
}
func makeImageFromPix(pix []uint8, stride int) *image.RGBA {
	w, h := sizeEidtPix(len(pix), stride)
	return &image.RGBA{
		Pix:    pix,
		Stride: stride,
		Rect:   image.Rect(0, 0, w, h),
	}
}

//
func sizeEidtPix(length, stride int) (w, h int) {
	return stride / STRIDECOUNT, length / stride
}
func offsetEditPix(x, y, stride int) int {
	return x*STRIDECOUNT + y*stride
}

//
func paddingEmpty(src []uint8, srcStride int, padW, padH int) (dst []uint8, dstStride int) {
	var w, h = sizeEidtPix(len(src), srcStride)
	var dstW, dstH = w + padW*2, h + padH*2
	dst, dstStride = createEditPix(dstW, dstH)
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			dstOffset := offsetEditPix(padW+x, padH+y, dstStride)
			srcOffset := offsetEditPix(x, y, srcStride)
			dst[dstOffset+R] = src[srcOffset+R]
			dst[dstOffset+G] = src[srcOffset+G]
			dst[dstOffset+B] = src[srcOffset+B]
			dst[dstOffset+A] = src[srcOffset+A]
		}
	}
	return dst, dstStride
}
func paddingExtend(src []uint8, srcStride int, padW, padH int) (dst []uint8, dstStride int) {
	dst, dstStride = paddingEmpty(src, srcStride, padW, padH)
	dstW, dstH := sizeEidtPix(len(dst), dstStride)
	for y := 0; y < dstH; y++ {
		iy := y
		if iy < padH {
			iy = padH
		} else if iy >= dstH-padH {
			iy = dstH - padH - 1
		}
		for x := 0; x < dstW; x++ {
			ix := x
			if ix < padW {
				ix = padW
			} else if x >= dstW-padW {
				ix = dstW - padW - 1
			} else if iy == y {
				// This only enters if we are not in a y-padded area or
				// x-padded area, so nothing to extend here.
				// So simply jump to the next padded-x index.
				x = dstW - padW - 1
				continue
			}
			dstOffset := offsetEditPix(x, y, dstStride)
			edgeOffset := offsetEditPix(ix, iy, dstStride)
			dst[dstOffset+R] = dst[edgeOffset+R]
			dst[dstOffset+G] = dst[edgeOffset+G]
			dst[dstOffset+B] = dst[edgeOffset+B]
			dst[dstOffset+A] = dst[edgeOffset+A]
		}
	}
	return dst, dstStride
}
