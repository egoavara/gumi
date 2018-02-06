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
	Map      map[string]interface{}
}

func (s *Style) useContext(ctx *gg.Context) {
	s.Default.Font.Use()
	ctx.SetLineWidth(s.Default.LineWidth)
	ctx.SetFontFace(s.Default.Font.Face())
}
func (s *Style) releaseContext(ctx *gg.Context) {
	s.Default.Font.Release()
}

type StyleDefault struct {
	Font *gutl.Font
	//
	LineWidth float64
	//
	Line image.Image
	Face image.Image
}

var (
	defaultStyleSingleton      *Style = nil
	defaultStyleSingletonMutex        = new(sync.RWMutex)
)

func DefaultStyle() *Style {
	if defaultStyleSingleton == nil {
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
			Map: map[string]interface{}{
				"dummy": nil,
			},
		}
		defaultStyleSingleton = temp
	} else {
		defaultStyleSingletonMutex.RLock()
		defer defaultStyleSingletonMutex.RUnlock()
	}
	return defaultStyleSingleton
}
