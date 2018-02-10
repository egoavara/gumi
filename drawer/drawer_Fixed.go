package drawer

import (
	"image"
	"image/color"
	"image/draw"
)

type Fixed struct {
	src  *image.RGBA
	Mode FillupMode
	//

}

func (s Fixed) ColorModel() color.Model {
	return s.src.ColorModel()
}

func (s Fixed) Bounds() image.Rectangle {
	return s.src.Bounds()
}

func (s Fixed) At(x, y int) color.Color {
	return s.src.At(x, y)
}

func (s Fixed) Draw(dst draw.Image){
	draw.Draw(dst, dst.Bounds(), s.src, s.src.Rect.Min, draw.Over)
}
