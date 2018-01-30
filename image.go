package gumi

import (
	"image"
	"image/color"
	"github.com/fogleman/gg"
)

func IsColorImage(img image.Image) (ok bool, c color.Color) {

	temp, ok := img.(*image.Uniform)
	if ok {
		return ok, temp.C
	}
	return false, img.At(0,0)
}

func GGContextRGBASub(rgba *image.RGBA, rect image.Rectangle) (*gg.Context) {
	sub := rgba.SubImage(rect).(*image.RGBA)
	sub.Rect = image.Rect(0,0, sub.Rect.Dx(), sub.Rect.Dy())
	ctx := gg.NewContextForRGBA(sub)
	return ctx
}