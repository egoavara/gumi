package gumi

import (
	"github.com/fogleman/gg"
	"github.com/golang/freetype"
	"github.com/iamGreedy/gumi/gutl"
	"golang.org/x/image/font/gofont/goregular"
	"image"
	"image/color"
)

type Style struct {
	Default  StyleDefault
	Material StyleMaterial
}
type StyleDefault struct {
	Font *gutl.Font
	//
	LineWidth float64
	//
	Line image.Image
	Face image.Image
}
type StyleMaterial struct {
	White  [2]image.Image
	Red    [2]image.Image
	Green  [2]image.Image
	Blue   [2]image.Image
	Yellow [2]image.Image
}
type MaterialColor uint8

const (
	White  MaterialColor = iota
	Red    MaterialColor = iota
	Green  MaterialColor = iota
	Blue   MaterialColor = iota
	Yellow MaterialColor = iota
)

func (s *StyleMaterial) PalletteImage(c MaterialColor) [2]image.Image{
	switch c {
	default:
		fallthrough
	case White:
		return s.White
	case Red:
		return s.Red
	case Green:
		return s.Green
	case Blue:
		return s.Blue
	case Yellow:
		return s.Yellow

	}
}
func (s *StyleMaterial) PalletteColor(c MaterialColor) [2]color.Color{
	switch c {
	default:
		fallthrough
	case White:
		return [2]color.Color{
			s.White[0].At(0,0),
			s.White[1].At(0,0),
		}
	case Red:
		return [2]color.Color{
			s.Red[0].At(0,0),
			s.Red[1].At(0,0),
		}
	case Green:
		return [2]color.Color{
			s.Green[0].At(0,0),
			s.Green[1].At(0,0),
		}
	case Blue:
		return [2]color.Color{
			s.Blue[0].At(0,0),
			s.Blue[1].At(0,0),
		}
	case Yellow:
		return [2]color.Color{
			s.Yellow[0].At(0,0),
			s.Yellow[1].At(0,0),
		}

	}
}

func DefaultStyle() *Style {
	f, _ := freetype.ParseFont(goregular.TTF)
	temp := &Style{
		Default: StyleDefault{
			Font:      gutl.NewFont(f, 12),
			LineWidth: 1,
			Line:      image.NewUniform(color.RGBA{113, 251, 254, 255}),
			Face:      image.NewUniform(color.RGBA{28, 30, 29, 255}),
		},
		Material: StyleMaterial{
			White: [2]image.Image{
				image.NewUniform(color.RGBA{28, 30, 29, 255}),
				image.NewUniform(color.RGBA{113, 251, 254, 255}),
			},
			Red: [2]image.Image{
				image.NewUniform(color.RGBA{28, 30, 29, 255}),
				image.NewUniform(color.RGBA{113, 251, 254, 255}),
			},
			Green: [2]image.Image{
				image.NewUniform(color.RGBA{28, 80, 29, 255}),
				image.NewUniform(color.RGBA{93, 255, 224, 255}),
			},
			Blue: [2]image.Image{
				image.NewUniform(color.RGBA{28, 30, 29, 255}),
				image.NewUniform(color.RGBA{113, 251, 254, 255}),
			},
			Yellow: [2]image.Image{
				image.NewUniform(color.RGBA{28, 30, 29, 255}),
				image.NewUniform(color.RGBA{113, 251, 254, 255}),
			},
		},
	}
	return temp
}

func (s *Style) useContext(ctx *gg.Context) {
	s.Default.Font.Use()
	ctx.SetLineWidth(s.Default.LineWidth)
	ctx.SetFontFace(s.Default.Font.Face())
}
func (s *Style) releaseContext(ctx *gg.Context) {
	s.Default.Font.Release()
}
