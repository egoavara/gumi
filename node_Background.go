package gumi

import (
	"fmt"
	"image"
	"github.com/iamGreedy/gumi/gumre"
	"github.com/iamGreedy/gumi/drawer"
)

type NBackground struct {
	SingleNode
	boundStore
	//
	drawer drawer.Drawer
}

func (s *NBackground) GUMIInfomation(info Information) {
	s.child.GUMIInfomation(info)
}
func (s *NBackground) GUMIStyle(style *Style) {
	s.child.GUMIStyle(style)
}
func (s *NBackground) GUMIClip(rect image.Rectangle) {
	s.bound = rect
	s.child.GUMIClip(rect)
}
func (s *NBackground) GUMIRender(frame *image.RGBA) {
	s.drawer.Draw(frame.SubImage(s.bound).(*image.RGBA))

}
func (s *NBackground) GUMIDraw(frame *image.RGBA) {
	s.GUMIRender(frame)
	s.child.GUMIDraw(frame)
}

func (s *NBackground) GUMIRenderTree(tree *drawer.RenderTree, parentnode *drawer.RenderNode) {
	panic("implement me")
}
func (s *NBackground) GUMIUpdate() {
	panic("implement me")
}

func (s *NBackground) GUMIHappen(event Event) {
	s.child.GUMIHappen(event)
}
func (s NBackground) GUMISize() gumre.Size {
	return s.child.GUMISize()
}
func (s *NBackground) String() string {
	return fmt.Sprintf("%s", "NBackground")
}

func NBackground0(draw drawer.Drawer) *NBackground {
	return &NBackground{
		drawer:draw,
	}
}
