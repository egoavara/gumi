package gumi

import (
	"fmt"
	"image"
)
type NMinimum struct {
	SingleStructure
	axis Axis
}

func (s *NMinimum) String() string {
	return fmt.Sprintf("%s", "NMinimum")
}
func (s *NMinimum) draw(frame *image.RGBA) {
	s.child.draw(frame)
}
func (s *NMinimum) size() Size {
	sz := s.child.size()
	if AxisVertical == AxisVertical & s.axis{
		sz.Vertical.Max = sz.Vertical.Min
	}
	if AxisHorizontal == AxisHorizontal & s.axis{
		sz.Horizontal.Max = sz.Horizontal.Min
	}
	return sz
}
func (s *NMinimum) rect(r image.Rectangle) {
	s.child.rect(r)
}
func (s *NMinimum) update(info *Information, style *Style) {
	s.child.update(info, style)
}
func (s *NMinimum) Occur(event Event) {
	s.child.Occur(event)
}

func NMinimum0(axis Axis, elem GUMI) *NMinimum {
	temp := &NMinimum{
		axis:axis,
	}
	elem.Born(temp)
	temp.Breed([]GUMI{elem})
	return temp
}