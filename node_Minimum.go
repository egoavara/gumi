package gumi

import (
	"fmt"
	"image"
	"github.com/iamGreedy/gumi/gumre"
	"github.com/iamGreedy/gumi/drawer"
)
type NMinimum struct {
	SingleStructure
	axis gumre.Axis
}

func (s *NMinimum) GUMIInfomation(info Information) {
	s.child.GUMIInfomation(info)
}
func (s *NMinimum) GUMIStyle(style *Style) {
	s.child.GUMIStyle(style)
}
func (s *NMinimum) GUMIClip(r image.Rectangle) {
	s.child.GUMIClip(r)
}
func (s *NMinimum) GUMIRender(frame *image.RGBA) {
}
func (s *NMinimum) GUMIDraw(frame *image.RGBA) {
	s.child.GUMIDraw(frame)
}

func (s *NMinimum) GUMIRenderTree(tree *drawer.RenderTree, parentnode *drawer.RenderNode) {
	panic("implement me")
}
func (s *NMinimum) GUMIUpdate() {
	panic("implement me")
}

func (s *NMinimum) GUMIHappen(event Event) {
	s.child.GUMIHappen(event)
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
func (s *NMinimum) String() string {
	return fmt.Sprintf("%s", "NMinimum")
}

func NMinimum0(axis gumre.Axis, elem GUMI) *NMinimum {
	temp := &NMinimum{
		axis:axis,
	}
	elem.born(temp)
	temp.breed([]GUMI{elem})
	return temp
}

func (s *NMinimum) Get() gumre.Axis {
	return s.GetAxis()
}
func (s *NMinimum) Set(axis gumre.Axis) {
	s.Set(axis)
}
func (s *NMinimum) GetAxis() gumre.Axis {
	return s.axis
}
func (s *NMinimum) SetAxis(axis gumre.Axis) {
	s.axis = axis
}