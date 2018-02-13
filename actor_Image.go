package gumi

import (
	"image"
	"fmt"
	"github.com/iamGreedy/gumi/gumre"
	"github.com/iamGreedy/gumi/drawer"
)

type AImage struct {
	VoidStructure
	boundStore
	//
	drawer drawer.Drawer
}



func (s *AImage) GUMIInfomation(info Information) {
}
func (s *AImage) GUMIStyle(style *Style) {
}
func (s *AImage) GUMIClip(rect image.Rectangle) {
	if s.bound != rect{
		s.bound = rect
	}
}
func (s *AImage) GUMIRender(frame *image.RGBA) {

	s.drawer.Draw(frame.SubImage(s.bound).(*image.RGBA))
}
func (s *AImage) GUMIDraw(frame *image.RGBA) {
	s.GUMIRender(frame)
}
func (s AImage) GUMISize() gumre.Size {
	bd := s.drawer.Bound()
	return gumre.Size{
		Horizontal: gumre.MinLength(uint16(bd.Dx())),
		Vertical:   gumre.MinLength(uint16(bd.Dy())),
	}
}
func (s *AImage) GUMIHappen(event Event) {
}
func (s *AImage) GUMIRenderTree(tree *drawer.RenderTree, parentnode *drawer.RenderNode) {
	// TODO
	panic("implement me")
}
func (s *AImage) GUMIUpdate() {
	// TODO
	panic("implement me")
}
func (s *AImage) String() string {
	return fmt.Sprintf("%s", "AImage")
}

func AImage0(drawer drawer.Drawer) *AImage {
	return &AImage{
		drawer:drawer,
	}
}
