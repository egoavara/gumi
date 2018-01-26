package gumi

import "image"

type nStyle struct {
	SingleStructure
	s *Style
}

func (s *nStyle) draw(frame *image.RGBA) {
	s.child.draw(frame)
}

func (s *nStyle) size() Size {
	return s.child.size()
}

func (s *nStyle) rect(r image.Rectangle) {
	s.child.rect(r)
}

func (s *nStyle) update(info *Information, style *Style) {
	s.child.update(info, s.s)
}

func (s *nStyle) Occur(event Event) {
	s.child.Occur(event)
}

func NStyle(s *Style) *nStyle {
	if s == nil {
		s = DefaultStyle()
	}
	return &nStyle{
		s: s,
	}
}
func (s *nStyle) Set(st *Style) {
	s.s = st
}
func (s *nStyle) Get() *Style {
	return s.s
}
