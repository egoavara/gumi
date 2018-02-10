package gumi

import (
	"image"
	"fmt"
	"github.com/iamGreedy/gumi/gumre"
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

func (s *NDrawing) GUMIRender(frame *image.RGBA) {
	var fn = func(){
		var ctx = createContextRGBASub(frame, s.bound)
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
	s.child.GUMIRender(frame)
}
func (s *NDrawing) GUMISize() gumre.Size {
	return s.child.GUMISize()
}
func (s *NDrawing) GUMIClip(r image.Rectangle) {
	s.bound = r
	s.child.GUMIClip(r)
}
func (s *NDrawing) GUMIUpdate(info *Information, style *Style) {
	s.style = style
	s.di = &DrawingInfo{
		info.Dt,
	}
	s.child.GUMIUpdate(info, style)
}
func (s *NDrawing) GUMIHappen(event Event) {
	s.child.GUMIHappen(event)
}
func (s *NDrawing) String() string {
	return fmt.Sprintf("%s(drawing:%d GUMIRender)", "NDrawing", len(s.drawfuncs))
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
