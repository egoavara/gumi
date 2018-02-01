package gumi

import (
	"fmt"
	"github.com/fogleman/gg"
	"strconv"
	"github.com/iamGreedy/gumi/gutl"
)

type RulerType uint16

const (
	RULER_GRADUATION_VERTICAL   RulerType = 1 << iota
	RULER_GRADUATION_HORIZONTAL RulerType = 1 << iota
	RULER_GRID_VERTICAL         RulerType = 1 << iota
	RULER_GRID_HORIZONTAL       RulerType = 1 << iota
	RULER_HINT_VERTICAL         RulerType = 1 << iota
	RULER_HINT_HORIZONTAL       RulerType = 1 << iota
	RULER_GRIDDASH_VERTICAL     RulerType = 1 << iota
	RULER_GRIDDASH_HORIZONTAL   RulerType = 1 << iota
	//
	RULER_STANDARDSCREEN RulerType = 1 << iota
	RULER_SCREEN         RulerType = 1 << iota
	RULER_PROPORTION     RulerType = 1 << iota
	RULER_DOTS           RulerType = 1 << iota
)

const RulerWidth = 10
const RulerDash1 = 2.
const RulerDash2 = 4.

type DrawFunc func(context *gg.Context, style *Style)
func BuildRuler(ruler RulerType, pivot int) []DrawFunc {
	var temp []DrawFunc
	if ruler&RULER_GRADUATION_VERTICAL == RULER_GRADUATION_VERTICAL {
		temp = append(temp, func(context *gg.Context, style *Style) {
			for f := 0.0; f <= float64(context.Height()); f += float64(pivot) {
				context.DrawLine(0, f, RulerWidth, f)
			}
			context.Stroke()
		})
	}
	if ruler&RULER_GRADUATION_HORIZONTAL == RULER_GRADUATION_HORIZONTAL {
		temp = append(temp, func(context *gg.Context, style *Style) {
			for f := 0.0; f <= float64(context.Width()); f += float64(pivot) {
				context.DrawLine(f, 0, f, RulerWidth)
			}
			context.Stroke()
		})
	}
	if ruler&RULER_GRID_VERTICAL == RULER_GRID_VERTICAL {
		temp = append(temp, func(context *gg.Context, style *Style) {
			for f := 0.0; f <= float64(context.Width()); f += float64(pivot) {
				context.DrawLine(f, 0, f, float64(context.Height()))
			}
			context.Stroke()
		})
	}
	if ruler&RULER_GRID_HORIZONTAL == RULER_GRID_HORIZONTAL {
		temp = append(temp, func(context *gg.Context, style *Style) {
			for f := 0.0; f <= float64(context.Height()); f += float64(pivot) {
				context.DrawLine(0, f, float64(context.Width()), f)
			}
			context.Stroke()
		})
	}
	if ruler&RULER_HINT_VERTICAL == RULER_HINT_VERTICAL {
		temp = append(temp, func(context *gg.Context, style *Style) {
			for f := float64(pivot); f <= float64(context.Width()); f += float64(pivot) {
				txt := strconv.FormatInt(int64(f), 10)
				w, _ := context.MeasureString(txt)
				context.DrawString(txt, f-w, float64(style.Default.Font.FontHeight().Round()))
			}
			context.Stroke()
		})
	}
	if ruler&RULER_HINT_HORIZONTAL == RULER_HINT_HORIZONTAL {
		temp = append(temp, func(context *gg.Context, style *Style) {
			for f := float64(pivot); f <= float64(context.Height()); f += float64(pivot) {
				context.DrawString(strconv.FormatInt(int64(f), 10), 0, f)
			}

			context.Stroke()
		})
	}
	if ruler&RULER_GRIDDASH_VERTICAL == RULER_GRIDDASH_VERTICAL {
		temp = append(temp, func(context *gg.Context, style *Style) {
			context.SetDash(RulerDash1, RulerDash2)
			for f := 0.0; f <= float64(context.Width()); f += float64(pivot) {
				context.DrawLine(f, 0, f, float64(context.Height()))
			}
			context.Stroke()
		})
	}
	if ruler&RULER_GRIDDASH_HORIZONTAL == RULER_GRIDDASH_HORIZONTAL {
		temp = append(temp, func(context *gg.Context, style *Style) {
			context.SetDash(RulerDash1, RulerDash2)
			for f := 0.0; f <= float64(context.Height()); f += float64(pivot) {
				context.DrawLine(0, f, float64(context.Width()), f)
			}
			context.Stroke()
		})
	}
	//==========================================================================================
	if ruler&RULER_SCREEN == RULER_SCREEN {
		temp = append(temp, func(context *gg.Context, style *Style) {
			context.DrawLine(0, RulerWidth, float64(context.Width()), RulerWidth)
			context.DrawLine(0, RulerWidth, RulerWidth/2, RulerWidth/2)
			context.DrawLine(0, RulerWidth, RulerWidth/2, RulerWidth/2*3)
			context.DrawLine(float64(context.Width()), RulerWidth, float64(context.Width())-RulerWidth/2, RulerWidth/2)
			context.DrawLine(float64(context.Width()), RulerWidth, float64(context.Width())-RulerWidth/2, RulerWidth/2*3)
			//
			context.DrawLine(RulerWidth, 0, RulerWidth, float64(context.Height()))
			context.DrawLine(RulerWidth, 0, RulerWidth/2, RulerWidth/2)
			context.DrawLine(RulerWidth, 0, RulerWidth/2*3, RulerWidth/2)
			context.DrawLine(RulerWidth, float64(context.Height()), RulerWidth/2, float64(context.Height())-RulerWidth/2)
			context.DrawLine(RulerWidth, float64(context.Height()), RulerWidth/2*3, float64(context.Height())-RulerWidth/2)
			//
			w, h := context.MeasureString(strconv.FormatInt(int64(context.Width()), 10))
			context.DrawString(strconv.FormatInt(int64(context.Width()), 10), float64(context.Width()/2)-w/2, RulerWidth+h/4)
			w, h = context.MeasureString(strconv.FormatInt(int64(context.Height()), 10))
			context.DrawString(strconv.FormatInt(int64(context.Height()), 10), RulerWidth, float64(context.Height()/2)+h/4)
			//
			context.Stroke()
		})
	}
	if ruler&RULER_PROPORTION == RULER_PROPORTION {
		temp = append(temp, func(context *gg.Context, style *Style) {
			w := context.Width()
			h := context.Height()
			//
			gcd := GCD(w, h)
			txt := fmt.Sprintf("%d : %d", w/gcd, h/gcd)
			//
			context.DrawLine(0, 0, float64(w), float64(h))
			bdw, bdh := context.MeasureString(txt)
			context.DrawString(txt, float64(w)/2-bdw/2, float64(h)/2+bdh/4)
			context.Stroke()
		})
	}
	if ruler&RULER_STANDARDSCREEN == RULER_STANDARDSCREEN {
		temp = append(temp, func(context *gg.Context, style *Style) {
			for _, v := range gutl.DefinedResolutions.Smaller(context.Width(), context.Height()){
				context.DrawRectangle(0,0, float64(v.Width), float64(v.Height))
				w, _ := context.MeasureString(v.Name)
				context.DrawString(v.Name, float64(v.Width) - w - 5, float64(v.Height) - 5)
			}
			context.Stroke()
		})
	}
	if ruler&RULER_DOTS == RULER_DOTS {
		temp = append(temp, func(context *gg.Context, style *Style) {
			for x := 0.; x <= float64(context.Width()); x += float64(pivot) {
				for y := 0.; y <= float64(context.Height()); y += float64(pivot) {
					context.DrawPoint(x, y, style.Default.LineWidth)
				}
			}
			context.Stroke()
		})
	}
	return temp
}
