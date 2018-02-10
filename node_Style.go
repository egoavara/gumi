package gumi

import (
	"image"
	"fmt"
	"github.com/iamGreedy/gumi/gumre"
)

type NStyle struct {
	SingleStructure
	s *Style
}

func (s *NStyle) String() string {
	return fmt.Sprintf("%s", "NStyle")
}

func (s *NStyle) GUMIRender(frame *image.RGBA) {
	s.child.GUMIRender(frame)
}
func (s *NStyle) GUMISize() gumre.Size {
	return s.child.GUMISize()
}
func (s *NStyle) GUMIClip(r image.Rectangle) {
	s.child.GUMIClip(r)
}
func (s *NStyle) GUMIUpdate(info *Information, style *Style) {
	s.child.GUMIUpdate(info, s.s)
}
func (s *NStyle) GUMIHappen(event Event) {
	s.child.GUMIHappen(event)
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
