package gumi

import "image"

type aCanvas struct {
	VoidStructure
	BoundStore
	StyleStore
	//
	w, h uint16
	fn DrawFunc
}

func (s *aCanvas) draw(frame *image.RGBA) {
	ctx := GGContextRGBASub(frame, s.bound)
	s.fn(ctx, s.style)
}

func (s aCanvas) size() Size {
	return Size{
		Horizontal: FixLength(uint16(s.w)),
		Vertical:   FixLength(uint16(s.h)),
	}
}
func (s *aCanvas) rect(rect image.Rectangle) {
	s.bound = rect
}

func (s *aCanvas) update(info *Information, style *Style) {
	s.style = style
}

func (s *aCanvas) Occur(event Event) {
}

func ACanvas(w, h uint16, fn DrawFunc) *aCanvas {
	return &aCanvas{
		w:w,
		h:h,
		fn:fn,
	}
}