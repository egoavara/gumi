package gumi

import (
	"image"
	"github.com/fogleman/gg"
)

func createContextRGBASub(rgba *image.RGBA, rect image.Rectangle) (*gg.Context) {
	sub := rgba.SubImage(rect).(*image.RGBA)
	sub.Rect = image.Rect(0,0, sub.Rect.Dx(), sub.Rect.Dy())
	ctx := gg.NewContextForRGBA(sub)
	return ctx
}