package gumi

import (
	"fmt"
	"image"
	"image/draw"
	"github.com/iamGreedy/gumi/gumre"
)

type NBackground struct {
	SingleStructure
	boundStore
	styleStore
	//
	img image.Image
}

func (s *NBackground) String() string {
	return fmt.Sprintf("%s", "NBackground")
}
func (s *NBackground) GUMIRender(frame *image.RGBA) {
	rect := s.img.Bounds()
	rect.Add(s.bound.Min)

	draw.Draw(frame, s.bound.Intersect(rect), s.img, rect.Min, draw.Over)
	s.child.GUMIRender(frame)
}
func (s NBackground) GUMISize() gumre.Size {
	return s.child.GUMISize()
}
func (s *NBackground) GUMIClip(rect image.Rectangle) {
	s.bound = rect
	s.child.GUMIClip(rect)
}
func (s *NBackground) GUMIUpdate(info *Information, style *Style) {
	s.style = style
	s.child.GUMIUpdate(info, style)
}
func (s *NBackground) GUMIHappen(event Event) {
	s.child.GUMIHappen(event)
}

func NBackground0(img image.Image) *NBackground {
	return &NBackground{
		img: img,
	}
}
