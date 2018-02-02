package gumi

import (
	"github.com/fogleman/gg"
	"fmt"
)

type DrawingInfo struct {
	Dt int64
}
type Drawer interface {
	Draw(context *gg.Context, style *Style, di *DrawingInfo)
}
type FunctionDrawer struct {
	fn func(context *gg.Context, style *Style, di *DrawingInfo)
}
func (s FunctionDrawer)Draw(context *gg.Context, style *Style, di *DrawingInfo)  {
	s.fn(context, style, di)
}

var Drawing _Drawing
type _Drawing struct {
	Ruler _Ruler
}

const FPSPos = 10

func (_Drawing) FPS() Drawer {
	return &fpsDrawer{}
}

const fpsDrawerHistory = 16

type fpsDrawer struct {
	dts [fpsDrawerHistory]float64
	i int
}
func (s *fpsDrawer ) Draw(context *gg.Context, style *Style, di *DrawingInfo)  {
	style.useContext(context)
	defer style.releaseContext(context)
	//
	s.dts[s.i] = float64(di.Dt)
	s.i = (s.i + 1) % fpsDrawerHistory
	//
	context.SetColor(style.Default.Line.At(0, 0))
	txt := fmt.Sprintf("%.2f", 1000 / float64(ifZeroBelowToOne(Average(s.dts[:]))))
	w := float64(context.Width())
	mw, mh := context.MeasureString(txt)
	context.DrawString(txt, w - FPSPos - mw, FPSPos + mh)
	context.Stroke()
}
func ifZeroBelowToOne(i float64) float64 {
	if i < 0{
		return 1
	}
	return i
}