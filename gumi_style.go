package gumi

import (
	"github.com/golang/freetype"
	"golang.org/x/image/font/gofont/goregular"
	"image"
	"image/color"
	"github.com/iamGreedy/gumi/gutl"
	"github.com/fogleman/gg"
)

type Style struct {
	Font *gutl.Font
	//
	LineWidth float64
	//
	Line image.Image
	Face image.Image
}



func DefaultStyle() *Style {
	f, _ := freetype.ParseFont(goregular.TTF)
	temp := &Style{
		Font:      gutl.NewFont(f, 12),
		LineWidth: 1,
		Line:      image.NewUniform(color.RGBA{113, 251, 254, 255}),
		Face:      image.NewUniform(color.RGBA{28, 30, 29, 255}),
	}
	return temp
}
func (s *Style) Create(style *Style) *Style {
	temp := &Style{
		Font:      s.Font,
		Face:      s.Face,
		Line:      s.Line,
		LineWidth: s.LineWidth,
	}
	if style.Font != nil {
		temp.Font = style.Font
	}
	if style.Line != nil {
		temp.Line = style.Line
	}
	if style.Face != nil {
		temp.Face = style.Face
	}
	if style.LineWidth != 0 {
		temp.LineWidth = style.LineWidth
	}
	return temp
}
func (s *Style) useContext(ctx *gg.Context) {
	s.Font.Use()
	ctx.SetLineWidth(s.LineWidth)
	ctx.SetFontFace(s.Font.Face())
}
func (s *Style) releaseContext(ctx *gg.Context) {
	s.Font.Release()
}

