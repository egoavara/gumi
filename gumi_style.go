package gumi

import (
	"github.com/fogleman/gg"
	"github.com/golang/freetype"
	"github.com/iamGreedy/gumi/gutl"
	"golang.org/x/image/font/gofont/goregular"
	"image"
	"image/color"
	"sync"
)

type Style struct {
	Default  StyleDefault
	Material StyleMaterial
	Map      map[string]interface{}
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

func (s MaterialColor) String() string {
	switch s {
	case White:
		return "White"
	case Red:
		return "Red"
	case Green:
		return "Green"
	case Blue:
		return "Blue"
	case Yellow:
		return "Yellow"
	default:
		return "Unknown"
	}
}

const (
	White  MaterialColor = iota
	Red    MaterialColor = iota
	Green  MaterialColor = iota
	Blue   MaterialColor = iota
	Yellow MaterialColor = iota
)

func (s *StyleMaterial) PalletteImage(c MaterialColor) [2]image.Image {
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
func (s *StyleMaterial) PalletteColor(c MaterialColor) [2]color.Color {
	switch c {
	default:
		fallthrough
	case White:
		return [2]color.Color{
			s.White[0].At(0, 0),
			s.White[1].At(0, 0),
		}
	case Red:
		return [2]color.Color{
			s.Red[0].At(0, 0),
			s.Red[1].At(0, 0),
		}
	case Green:
		return [2]color.Color{
			s.Green[0].At(0, 0),
			s.Green[1].At(0, 0),
		}
	case Blue:
		return [2]color.Color{
			s.Blue[0].At(0, 0),
			s.Blue[1].At(0, 0),
		}
	case Yellow:
		return [2]color.Color{
			s.Yellow[0].At(0, 0),
			s.Yellow[1].At(0, 0),
		}

	}
}

var (
	defaultStyleSingleton *Style = nil
	defaultStyleSingletonMutex = new(sync.RWMutex)
)

func DefaultStyle() *Style {
	if defaultStyleSingleton == nil{
		defaultStyleSingletonMutex.Lock()
		defer defaultStyleSingletonMutex.Unlock()
		f, _ := freetype.ParseFont(goregular.TTF)
		temp := &Style{
			Default: StyleDefault{
				Font:      gutl.NewFont(f, 12),
				LineWidth: 1,
				Line:      image.NewUniform(color.RGBA{252, 252, 252, 255}),
				Face:      image.NewUniform(color.RGBA{64, 74, 83, 255}),
			},
			Material: StyleMaterial{
				White: [2]image.Image{
					image.NewUniform(color.RGBA{94, 97, 97, 255}),
					image.NewUniform(color.RGBA{213, 217, 218, 255}),
				},
				Red: [2]image.Image{
					image.NewUniform(color.RGBA{127, 51, 49, 255}),
					image.NewUniform(color.RGBA{255, 84, 74, 255}),
				},
				Green: [2]image.Image{
					image.NewUniform(color.RGBA{61, 127, 59, 255}),
					image.NewUniform(color.RGBA{110, 204, 102, 255}),
				},
				Blue: [2]image.Image{
					image.NewUniform(color.RGBA{59, 70, 127, 255}),
					image.NewUniform(color.RGBA{118, 139, 255, 255}),
				},
				Yellow: [2]image.Image{
					image.NewUniform(color.RGBA{122, 127, 67, 255}),
					image.NewUniform(color.RGBA{231, 235, 118, 255}),
				},
			},
			Map: map[string]interface{}{
				"dummy": nil,
			},
		}
		defaultStyleSingleton = temp
	}else {
		defaultStyleSingletonMutex.RLock()
		defer defaultStyleSingletonMutex.RUnlock()
	}
	return defaultStyleSingleton
}

func (s *Style) useContext(ctx *gg.Context) {
	s.Default.Font.Use()
	ctx.SetLineWidth(s.Default.LineWidth)
	ctx.SetFontFace(s.Default.Font.Face())
}
func (s *Style) releaseContext(ctx *gg.Context) {
	s.Default.Font.Release()
}