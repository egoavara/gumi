package gumi

import (
	"fmt"
	"github.com/fogleman/gg"
	"github.com/iamGreedy/gumi/gutl"
	"strconv"
)

const RulerWidth = 10
const RulerDash1 = 2.
const RulerDash2 = 4.

type _Ruler struct {
	Graduation
	Grid
	Hint
	Dashgrid
}
type Graduation struct{}

func (Graduation) Vertical(pivot float64) Drawer {
	return FunctionDrawer{func(context *gg.Context, style *Style, di *DrawingInfo) {
		style.useContext(context)
		defer style.releaseContext(context)
		context.SetColor(style.Default.Line.At(0, 0))
		for f := 0.0; f <= float64(context.Height()); f += float64(pivot) {
			context.DrawLine(0, f, RulerWidth, f)
		}
		context.Stroke()
	}}
}
func (Graduation) Horizontal(pivot float64) Drawer {
	return FunctionDrawer{func(context *gg.Context, style *Style, di *DrawingInfo) {
		style.useContext(context)
		defer style.releaseContext(context)
		context.SetColor(style.Default.Line.At(0, 0))
		for f := 0.0; f <= float64(context.Width()); f += float64(pivot) {
			context.DrawLine(f, 0, f, RulerWidth)
		}
		context.Stroke()
	}}
}

type Grid struct{}

func (Grid) Vertical(pivot float64) Drawer {
	return FunctionDrawer{func(context *gg.Context, style *Style, di *DrawingInfo) {
		style.useContext(context)
		defer style.releaseContext(context)
		context.SetColor(style.Default.Line.At(0, 0))
		for f := 0.0; f <= float64(context.Width()); f += float64(pivot) {
			context.DrawLine(f, 0, f, float64(context.Height()))
		}
		context.Stroke()
	}}
}
func (Grid) Horizontal(pivot float64) Drawer {
	return FunctionDrawer{func(context *gg.Context, style *Style, di *DrawingInfo) {
		style.useContext(context)
		defer style.releaseContext(context)
		context.SetColor(style.Default.Line.At(0, 0))
		for f := 0.0; f <= float64(context.Height()); f += float64(pivot) {
			context.DrawLine(0, f, float64(context.Width()), f)
		}
		context.Stroke()
	}}
}

type Hint struct{}

func (Hint) Vertical(pivot float64) Drawer {
	return FunctionDrawer{func(context *gg.Context, style *Style, di *DrawingInfo) {
		style.useContext(context)
		defer style.releaseContext(context)
		context.SetColor(style.Default.Line.At(0, 0))
		for f := float64(pivot); f <= float64(context.Width()); f += float64(pivot) {
			txt := strconv.FormatInt(int64(f), 10)
			w, _ := context.MeasureString(txt)
			context.DrawString(txt, f-w, float64(style.Default.Font.FontHeight().Round()))
		}
		context.Stroke()
	}}
}
func (Hint) Horizontal(pivot float64) Drawer {

	return FunctionDrawer{func(context *gg.Context, style *Style, di *DrawingInfo) {
		style.useContext(context)
		defer style.releaseContext(context)
		context.SetColor(style.Default.Line.At(0, 0))
		for f := float64(pivot); f <= float64(context.Height()); f += float64(pivot) {
			context.DrawString(strconv.FormatInt(int64(f), 10), 0, f)
		}

		context.Stroke()
	}}
}

type Dashgrid struct{}

func (Dashgrid) Vertical(pivot float64) Drawer {
	return FunctionDrawer{func(context *gg.Context, style *Style, di *DrawingInfo) {
		style.useContext(context)
		defer style.releaseContext(context)
		context.SetColor(style.Default.Line.At(0, 0))

		context.SetDash(RulerDash1, RulerDash2)
		for f := 0.0; f <= float64(context.Width()); f += float64(pivot) {
			context.DrawLine(f, 0, f, float64(context.Height()))
		}
		context.Stroke()
	}}
}
func (Dashgrid) Horizontal(pivot float64) Drawer {
	return FunctionDrawer{func(context *gg.Context, style *Style, di *DrawingInfo) {
		style.useContext(context)
		defer style.releaseContext(context)
		context.SetColor(style.Default.Line.At(0, 0))
		context.SetDash(RulerDash1, RulerDash2)
		for f := 0.0; f <= float64(context.Height()); f += float64(pivot) {
			context.DrawLine(0, f, float64(context.Width()), f)
		}
		context.Stroke()
	}}
}

func (_Ruler) Size() Drawer {
	return FunctionDrawer{func(context *gg.Context, style *Style, di *DrawingInfo) {
		style.useContext(context)
		defer style.releaseContext(context)
		context.SetColor(style.Default.Line.At(0, 0))

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
	}}
}
func (_Ruler) Proportion() Drawer {
	return FunctionDrawer{func(context *gg.Context, style *Style, di *DrawingInfo) {
		style.useContext(context)
		defer style.releaseContext(context)
		context.SetColor(style.Default.Line.At(0, 0))

		w := context.Width()
		h := context.Height()
		//
		gcd := int(GCD(int64(w), int64(h)))
		txt := fmt.Sprintf("%d : %d", w/gcd, h/gcd)
		//
		context.DrawLine(0, 0, float64(w), float64(h))
		bdw, bdh := context.MeasureString(txt)
		context.DrawString(txt, float64(w)/2-bdw/2, float64(h)/2+bdh/4)
		context.Stroke()
	}}
}
func (_Ruler) Screen() Drawer {
	return FunctionDrawer{func(context *gg.Context, style *Style, di *DrawingInfo) {
		style.useContext(context)
		defer style.releaseContext(context)
		context.SetColor(style.Default.Line.At(0, 0))
		for _, v := range gutl.DefinedResolutions.Smaller(context.Width(), context.Height()) {
			context.DrawRectangle(0, 0, float64(v.Width), float64(v.Height))
			w, _ := context.MeasureString(v.Name)
			context.DrawString(v.Name, float64(v.Width)-w-5, float64(v.Height)-5)
		}
		context.Stroke()
	}}
}
func (_Ruler) Dots(pivot float64) Drawer {
	return FunctionDrawer{func(context *gg.Context, style *Style, di *DrawingInfo) {
		style.useContext(context)
		defer style.releaseContext(context)
		context.SetColor(style.Default.Line.At(0, 0))
		for x := 0.; x <= float64(context.Width()); x += float64(pivot) {
			for y := 0.; y <= float64(context.Height()); y += float64(pivot) {
				context.DrawPoint(x, y, style.Default.LineWidth)
			}
		}
		context.Stroke()
	}}
}
