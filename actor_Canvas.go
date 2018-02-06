package gumi

import (
	"image"
	"fmt"
)

type ACanvas struct {
	VoidStructure
	BoundStore
	StyleStore
	//
	w, h uint16
	fn   Drawer
//
	di *DrawingInfo
}

func (s *ACanvas) draw(frame *image.RGBA) {
	ctx := GGContextRGBASub(frame, s.bound)
	s.fn.Draw(ctx, s.style, s.di)
}
func (s ACanvas) size() Size {
	return Size{
		Horizontal: FixLength(uint16(s.w)),
		Vertical:   FixLength(uint16(s.h)),
	}
}
func (s *ACanvas) rect(rect image.Rectangle) {
	s.bound = rect
}
func (s *ACanvas) update(info *Information, style *Style) {
	s.style = style
	s.di = &DrawingInfo{
		Dt:info.Dt,
	}
}
func (s *ACanvas) Occur(event Event) {
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