package gumi

import (
	"image"
	"fmt"
	"github.com/iamGreedy/gumi/gumre"
)

type ACanvas struct {
	VoidStructure
	boundStore
	styleStore
	//
	w, h uint16
	fn   Drawer
//
	di *DrawingInfo
}

func (s *ACanvas) GUMIRender(frame *image.RGBA) {
	ctx := createContextRGBASub(frame, s.bound)
	s.fn.Draw(ctx, s.style, s.di)
}
func (s ACanvas) GUMISize() gumre.Size {
	return gumre.Size{
		Horizontal: gumre.FixLength(uint16(s.w)),
		Vertical:   gumre.FixLength(uint16(s.h)),
	}
}
func (s *ACanvas) GUMIClip(rect image.Rectangle) {
	s.bound = rect
}
func (s *ACanvas) GUMIUpdate(info *Information, style *Style) {
	s.style = style
	s.di = &DrawingInfo{
		Dt:info.Dt,
	}
}
func (s *ACanvas) GUMIHappen(event Event) {
}
func (s *ACanvas) String() string {
	return fmt.Sprintf("%s", "ACanvas")
}
func ACanvas0(w, h uint16, fn Drawer) *ACanvas {
	return &ACanvas{
		w:w,
		h:h,
		fn:fn,
	}
}