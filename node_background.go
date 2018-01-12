package gumi

import (
	"image/draw"
)

type nBackground struct {
	GUMILINK_SINGLE
}

func NBackground() *nBackground {
	return &nBackground{}
}
func (s *nBackground) size(drawing *Drawing, style *Style) Size {

	return s.child.(GUMIElem).size(drawing, style)
}
func (s *nBackground) draw(drawing *Drawing, style *Style, frame Frame) {
	draw.Draw(frame, frame.Bounds(), style.Face, style.Face.Bounds().Min, draw.Over)
	if s.child != nil {
		s.child.(GUMIElem).draw(drawing, style, frame)
	}
}
