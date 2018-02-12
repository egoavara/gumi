package gumi

import (
	"github.com/fogleman/gg"
	"fmt"
	"github.com/iamGreedy/gumi/gumre"
)


type Drawer interface {
	Draw(context *gg.Context, style *Style, di Information)
}
type FunctionDrawer struct {
	fn func(context *gg.Context, style *Style, di Information)
}
func (s FunctionDrawer)Draw(context *gg.Context, style *Style, di Information)  {
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
func (s *fpsDrawer ) Draw(context *gg.Context, style *Style, di Information)  {
	style.useContext(context)
	defer style.releaseContext(context)
	//
	s.dts[s.i] = float64(di.Dt)
	s.i = (s.i + 1) % fpsDrawerHistory
	//
	context.SetColor(rulerColor)
	avg := ifZeroBelowToOne(gumre.Average(s.dts[:]))
	txt := fmt.Sprintf("FPS : %.2f - AVG : %2.5f", 1000 / float64(avg), avg)
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