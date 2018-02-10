package drawer

import (
	"image"
	"image/color"
	"image/draw"
)

const (
	RED         = 0
	GREEN       = 1
	BLUE        = 2
	ALPHA       = 3
	STRIDECOUNT = 4
)

func startEditPix(src image.Image) (pix []uint8, stride int) {
	if v, ok := src.(*image.RGBA); ok {
		return v.Pix, v.Stride
	}
	var imageSize = src.Bounds().Size()
	stride = imageSize.X * STRIDECOUNT
	var pixlen = stride * imageSize.Y
	//

	pix = make([]uint8, pixlen)
	// src Copy to pix
	for x := 0; x < imageSize.X; x++ {
		for y := 0; y < imageSize.Y; y++ {
			r, g, b, a := src.At(x, y).RGBA()
			offset := offsetEditPix(x, y, stride)
			pix[offset+RED] = uint8(r >> 8)
			pix[offset+GREEN] = uint8(g >> 8)
			pix[offset+BLUE] = uint8(b >> 8)
			pix[offset+ALPHA] = uint8(a >> 8)
		}
	}
	return
}
func endEditPix(dst draw.Image, pix []uint8, stride int) {
	if _, ok := dst.(*image.RGBA); ok {
		return
	}
	var w, h = sizeEidtPix(len(pix), stride)
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			offset := offsetEditPix(x, y, stride)
			dst.Set(x, y, color.RGBA{
				pix[offset+RED],
				pix[offset+GREEN],
				pix[offset+BLUE],
				pix[offset+ALPHA],
			})
		}
	}
}
func sizeEidtPix(length, stride int) (w, h int) {
	return stride / STRIDECOUNT, length / stride
}
func offsetEditPix(x, y, stride int) int {
	return x*STRIDECOUNT + y*stride
}
