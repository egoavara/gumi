package gumi

import (
	"image"
	"image/draw"
)

type nBackground struct {
	SingleStructure
	BoundStore
	StyleStore
}

func (s *nBackground) draw(frame *image.RGBA) {
	draw.Draw(frame, s.bound, s.style.Default.Face, s.style.Default.Face.Bounds().Min, draw.Over)
	s.child.draw(frame)
}

func (s nBackground) size() Size {
	return s.child.size()
}
func (s *nBackground) rect(rect image.Rectangle) {
	s.bound = rect
	s.child.rect(rect)
}

func (s *nBackground) update(info *Information, style *Style) {
	s.style = style
	s.child.update(info, style)
}

func (s *nBackground) Occur(event Event) {
	s.child.Occur(event)
}

func NBackground() *nBackground {
	return &nBackground{}
}
