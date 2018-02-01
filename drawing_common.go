package gumi

import "github.com/fogleman/gg"

type DrawingFn func(context *gg.Context, style *Style)

var Drawing _Drawing

type _Drawing struct {
	Ruler _Ruler
}

