package gumi

import (
	"github.com/fogleman/gg"
)

type DrawFunc func(context *gg.Context, style *Style)
type nDrawing struct {
	GUMILINK_SINGLE
	afterdraw bool
	drawfuncs []DrawFunc
}

func (s *nDrawing) size(drawing *Drawing, style *Style) Size {
	return s.child.(GUMIElem).size(drawing, style)
}
func (s *nDrawing) draw(drawing *Drawing, style *Style, frame Frame) {
	ctx := gg.NewContextForRGBA(frame.rgbaCompatible())
	ctx.SetFontFace(style.Font.face)
	ctx.SetLineWidth(style.LineWidth)
	ctx.SetColor(style.Line.At(0, 0))
	if !s.afterdraw {
		style.Font.Use()
		for _, f := range s.drawfuncs {
			ctx.Push()
			f(ctx, style)
			ctx.Pop()
		}
		style.Font.Release()
	}
	s.child.(GUMIElem).draw(drawing, style, frame)
	if s.afterdraw {
		style.Font.Use()
		for _, f := range s.drawfuncs {
			ctx.Push()
			f(ctx, style)
			ctx.Pop()
		}
		style.Font.Release()
	}
}

//
func (s *nDrawing) AfterDraw(on bool) *nDrawing {
	s.afterdraw = on
	return s
}
func (s *nDrawing) IsAfterDraw() bool {
	return s.afterdraw
}

//
func NDrawing(beforedraw bool, drawFuncs ...DrawFunc) *nDrawing {
	return &nDrawing{
		afterdraw: !beforedraw,
		drawfuncs: drawFuncs,
	}
}
func NDrawing1(drawFuncs ...DrawFunc) *nDrawing {
	return &nDrawing{
		afterdraw: true,
		drawfuncs: drawFuncs,
	}
}
