package gumi

import (
	"github.com/golang/freetype"
	"github.com/iamGreedy/gumi/gutl"
	"golang.org/x/image/font/gofont/goregular"
	"image"
	"image/color"
)

const (
	THEME_COLOR      = 5
	THEME_NORMAL     = 3
	THEME_BACKGROUND = 2
)
const (
	INTENSE1 = 0
	INTENSE2 = 1
	INTENSE3 = 2
)
const (
	BACKGROUND_INACTIVE = 0
	BACKGROUND_ACTIVE   = 1
)

var DefaultDarkTheme Theme

func init() {
	f, _ := freetype.ParseFont(goregular.TTF)
	DefaultDarkTheme = Theme{
		Font:      gutl.NewFont(f, 12),
		LineWidth: 2,
		NormalColors: [THEME_NORMAL]image.Image{
			image.NewUniform(color.RGBA{72, 79, 87, 255}),
			image.NewUniform(color.RGBA{152, 166, 173, 255}),
			image.NewUniform(color.RGBA{252, 252, 252, 255}),
		},
		Colors: [THEME_COLOR]image.Image{
			image.NewUniform(color.RGBA{66, 176, 216, 255}),
			image.NewUniform(color.RGBA{245, 101, 151, 255}),
			image.NewUniform(color.RGBA{99, 189, 170, 255}),
			image.NewUniform(color.RGBA{51, 51, 51, 255}),
			image.NewUniform(color.RGBA{253, 169, 42, 255}),
		},
		Background: [THEME_BACKGROUND]image.Image{
			image.NewUniform(color.RGBA{51, 59, 68, 255}),
			image.NewUniform(color.RGBA{54 + 10, 64 + 10, 73 + 10, 255 - 15}),
		},
	}
}

type Theme struct {
	Font         *gutl.Font
	LineWidth    float64
	NormalColors [THEME_NORMAL]image.Image
	Colors       [THEME_COLOR]image.Image
	Background   [THEME_BACKGROUND]image.Image
}

func (s *Theme) From(theme Theme) Theme {
	temp := *s
	if theme.Font != nil {
		temp.Font = theme.Font
	}
	if theme.LineWidth != 0 {
		temp.LineWidth = theme.LineWidth
	}
	for i, v := range theme.NormalColors {
		if v != nil {
			temp.NormalColors[i] = v
		}
	}
	for i, v := range theme.Colors {
		if v != nil {
			temp.Colors[i] = v
		}
	}
	for i, v := range theme.Background {
		if v != nil {
			temp.Background[i] = v
		}
	}
	return temp
}

func (s *Theme) BackgroundStyle() *Style {
	return &Style{
		Default: StyleDefault{
			Font:      s.Font,
			LineWidth: s.LineWidth,
			Face:      s.Background[BACKGROUND_INACTIVE],
			Line:      s.NormalColors[INTENSE3],
		},
	}
}
func (s *Theme) Style(intense int) *Style {
	return &Style{
		Default: StyleDefault{
			Font:      s.Font,
			LineWidth: s.LineWidth,
			Face:      s.Background[BACKGROUND_ACTIVE],
			Line:      s.NormalColors[intense],
		},
	}
}
func (s *Theme) ColorLine(color int) *Style {
	temp := StyleDefault{
		Font:      s.Font,
		LineWidth: s.LineWidth,
	}
	temp.Face = s.Background[BACKGROUND_ACTIVE]
	temp.Line = s.Colors[color]
	return &Style{
		Default: temp,
	}
}
func (s *Theme) ColorFace(color int, intense int) *Style {
	temp := StyleDefault{
		Font:      s.Font,
		LineWidth: s.LineWidth,
	}
	temp.Face = s.Colors[color]
	temp.Line = s.NormalColors[intense]
	return &Style{
		Default: temp,
	}
}
