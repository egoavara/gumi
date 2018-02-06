package gumi

import (
	"image"
	"fmt"
)

type NDrawing struct {
	SingleStructure
	boundStore
	styleStore
	//
	afterdraw bool
	drawfuncs []Drawer
//
	di *DrawingInfo
}

func (s *NDrawing) draw(frame *image.RGBA) {
	var fn = func(){
		var ctx = GGContextRGBASub(frame, s.bound)
		for _, f := range s.drawfuncs {
			ctx.Push()
			f.Draw(ctx, s.style, s.di)
			ctx.Pop()
		}
	}
	if !s.afterdraw {
		fn()
	}else {
		defer fn()
	}
	s.child.draw(frame)
}
func (s *NDrawing) size() Size {
	return s.child.size()
}
func (s *NDrawing) rect(r image.Rectangle) {
	s.bound = r
	s.child.rect(r)
}
func (s *NDrawing) update(info *Information, style *Style) {
	s.style = style
	s.di = &DrawingInfo{
		info.Dt,
	}
	s.child.update(info, style)
}
func (s *NDrawing) Occur(event Event) {
	s.child.Occur(event)
}
func (s *NDrawing) String() string {
	return fmt.Sprintf("%s(drawing:%d draw)", "NDrawing", len(s.drawfuncs))
}

//
func (s *NDrawing) AfterDraw(on bool) *NDrawing {
	s.afterdraw = on
	return s
}
func (s *NDrawing) IsAfterDraw() bool {
	return s.afterdraw
}

//
func NDrawing0(beforedraw bool, drawFuncs ...Drawer) *NDrawing {
	return &NDrawing{
		afterdraw: !beforedraw,
		drawfuncs: drawFuncs,
	}
}
func NDrawing1(drawFuncs ...Drawer) *NDrawing {
	return &NDrawing{
		afterdraw: true,
		drawfuncs: drawFuncs,
	}
}
