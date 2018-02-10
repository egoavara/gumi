package gumi

import (
	"github.com/fogleman/gg"
	"github.com/golang/freetype"
	"github.com/iamGreedy/gumi/gumre"
	"golang.org/x/image/font/gofont/goregular"
	"sync"
	"github.com/golang/freetype/truetype"
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
	Font *gumre.Font
	//
	LineWidth float64
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
				Font:      gumre.NewFont(f, 12),
				LineWidth: 1,
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
func ModifyDefaultStyle(font *truetype.Font, size float64) {
	defaultStyleSingletonMutex.Lock()
	defer defaultStyleSingletonMutex.Unlock()
	if defaultStyleSingleton == nil {
		temp := &Style{
			Default: StyleDefault{
				Font:      gumre.NewFont(font, size),
				LineWidth: 1,
			},
			Map: map[string]interface{}{
				"dummy": nil,
			},
		}
		defaultStyleSingleton = temp
	} else {
		defaultStyleSingleton.Default.Font = gumre.NewFont(font, size)
	}
}
