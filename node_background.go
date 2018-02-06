package gumi

import (
	"image"
	"image/draw"
	"fmt"
)

type NBackground struct {
	SingleStructure
	boundStore
	styleStore
}

func (s *NBackground) String() string {
	return fmt.Sprintf("%s", "NBackground")
}
func (s *NBackground) draw(frame *image.RGBA) {
	draw.Draw(frame, s.bound, s.style.Default.Face, s.style.Default.Face.Bounds().Min, draw.Over)
	s.child.draw(frame)
}
func (s NBackground) size() Size {
	return s.child.size()
}
func (s *NBackground) rect(rect image.Rectangle) {
	s.bound = rect
	s.child.rect(rect)
}
func (s *NBackground) update(info *Information, style *Style) {
	s.style = style
	s.child.update(info, style)
}
func (s *NBackground) Occur(event Event) {
	s.child.Occur(event)
}

func NBackground0() *NBackground {
	return &NBackground{}
}
