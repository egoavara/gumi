package gumi

import (
	"fmt"
	"image"
	"github.com/iamGreedy/gumi/gumre"
)
type NMinimum struct {
	SingleStructure
	axis gumre.Axis
}

func (s *NMinimum) String() string {
	return fmt.Sprintf("%s", "NMinimum")
}
func (s *NMinimum) GUMIRender(frame *image.RGBA) {
	s.child.GUMIRender(frame)
}
func (s *NMinimum) GUMISize() gumre.Size {
	sz := s.child.GUMISize()
	if gumre.AxisVertical == gumre.AxisVertical & s.axis{
		sz.Vertical.Max = sz.Vertical.Min
	}
	if gumre.AxisHorizontal == gumre.AxisHorizontal & s.axis{
		sz.Horizontal.Max = sz.Horizontal.Min
	}
	return sz
}
func (s *NMinimum) GUMIClip(r image.Rectangle) {
	s.child.GUMIClip(r)
}
func (s *NMinimum) GUMIUpdate(info *Information, style *Style) {
	s.child.GUMIUpdate(info, style)
}
func (s *NMinimum) GUMIHappen(event Event) {
	s.child.GUMIHappen(event)
}

func NMinimum0(axis gumre.Axis, elem GUMI) *NMinimum {
	temp := &NMinimum{
		axis:axis,
	}
	elem.born(temp)
	temp.breed([]GUMI{elem})
	return temp
}