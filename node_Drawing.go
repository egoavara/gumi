package gumi

import (
	"github.com/fogleman/gg"
	"image"
)

type nDrawing struct {
	SingleStructure
	BoundStore
	StyleStore
	//
	afterdraw bool
	drawfuncs []DrawFunc
}

func (s *nDrawing) draw(frame *image.RGBA) {
	var ctx = GGContextRGBASub(frame, s.bound)
	s.style.useContext(ctx)
	defer  s.style.releaseContext(ctx)
	ctx.SetColor(s.style.Default.Line.At(0, 0))
	if !s.afterdraw {
		for _, f := range s.drawfuncs {
			ctx.Push()
			f(ctx, s.style)
			ctx.Pop()
		}
	}
	s.child.draw(frame)
	if s.afterdraw {
		for _, f := range s.drawfuncs {
			ctx.Push()
			f(ctx, s.style)
			ctx.Pop()
		}
	}
}
func (s *nDrawing) size() Size {
	return s.child.size()
}
func (s *nDrawing) rect(r image.Rectangle) {
	s.bound = r
	s.child.rect(r)
}
func (s *nDrawing) update(info *Information, style *Style) {
	s.style = style
	s.child.update(info, style)
}
func (s *nDrawing) Occur(event Event) {
	s.child.Occur(event)
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
