package drawer

import (
	"image"
	"image/color"
	"image/draw"
)

type Uniform struct{
	uni *image.Uniform
}

func NewUniform(c color.Color) *Uniform {
	return &Uniform{uni:image.NewUniform(c),}
}

func (Uniform) Bound() image.Rectangle {
	return image.Rect(0,0,0,0)
}
func (s Uniform) Draw(dst draw.Image) {
	draw.Draw(dst, dst.Bounds(), s.uni, image.ZP, draw.Over)
}