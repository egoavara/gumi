package drawer

import (
	"image"
	"image/color"
	"image/draw"
)

type Fixed struct {
	src  *image.RGBA
}

func NewFixed(img image.Image) *Fixed {
	var src *image.RGBA
	var ok bool
	if src, ok = img.(*image.RGBA); !ok {
		src = image.NewRGBA(img.Bounds())
		draw.Draw(src, src.Rect, img, image.ZP, draw.Src)
	}
	return &Fixed{
		src:  src,
	}
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
	draw.Draw(dst, dst.Bounds().Intersect(s.src.Rect), s.src, s.src.Rect.Min, draw.Src)
}
