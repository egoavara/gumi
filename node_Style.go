package gumi

import (
	"image"
	"fmt"
)

type NStyle struct {
	SingleStructure
	s *Style
}

func (s *NStyle) String() string {
	return fmt.Sprintf("%s", "NStyle")
}

func (s *NStyle) draw(frame *image.RGBA) {
	s.child.draw(frame)
}
func (s *NStyle) size() Size {
	return s.child.size()
}
func (s *NStyle) rect(r image.Rectangle) {
	s.child.rect(r)
}
func (s *NStyle) update(info *Information, style *Style) {
	s.child.update(info, s.s)
}
func (s *NStyle) Occur(event Event) {
	s.child.Occur(event)
}

func NStyle0(s *Style) *NStyle {
	if s == nil {
		s = DefaultStyle()
	}
	return &NStyle{
		s: s,
	}
}
func (s *NStyle) Set(st *Style) {
	s.s = st
}
func (s *NStyle) Get() *Style {
	return s.s
}
