package gumi

import (
	"fmt"
	"github.com/iamGreedy/gumi/drawer"
	"github.com/iamGreedy/gumi/gumre"
	"image"
)

// Actor::Image
//
// AImage is an element for outputting images.
// The image uses iamGreedy / drawer.Drawer rather than image.Image
type AImage struct {
	VoidNode
	boundStore
	frameStore
	//
	drawer drawer.Drawer
}

// GUMIFunction / GUMIInit 					-> VoidNode::Default

// GUMIFunction / GUMIInfomation 			-> Define::Empty
func (s *AImage) GUMIInfomation(info Information) {
}

// GUMIFunction / GUMIStyle 			-> Define::Empty
func (s *AImage) GUMIStyle(style *Style) {
}

// GUMIFunction / GUMIClip 			-> Define
func (s *AImage) GUMIClip(rect image.Rectangle) {
	if s.bound != rect {
		s.bound = rect
	}
}

// GUMIFunction / GUMIRender 		-> Define
func (s *AImage) GUMIRender(frame *image.RGBA) {
	s.drawer.Draw(frame.SubImage(s.bound).(*image.RGBA))
}

// GUMIFunction / GUMISize 		-> Define
func (s AImage) GUMISize() gumre.Size {
	bd := s.drawer.Bound()
	return gumre.Size{
		Horizontal: gumre.MinLength(uint16(bd.Dx())),
		Vertical:   gumre.MinLength(uint16(bd.Dy())),
	}
}

// GUMITree / born 							-> VoidNode::Default

// GUMITree / breed 						-> VoidNode::Default

// GUMITree / Parent()						-> VoidNode::Default

// GUMITree / Childrun()					-> VoidNode::Default

// GUMIRenderer / GUMIRenderSetup			-> Define
func (s *AImage) GUMIRenderSetup(frame *image.RGBA, tree *drawer.RenderTree, parentnode *drawer.RenderNode) {
	s.frame = frame
	// TODO : Cache
}

// GUMIRenderer / GUMIDraw					-> Define
func (s *AImage) GUMIDraw() {
	s.GUMIRender(s.frame)
}

// GUMIRenderer / GUMIUpdate					-> Define
func (s *AImage) GUMIUpdate() {
	// TODO
	panic("implement me")
}

// GUMIEventer / GUMIHappen					-> Define
func (s *AImage) GUMIHappen(event Event) {
}

// fmt.Stringer / String					-> Define
func (s *AImage) String() string {
	return fmt.Sprintf("%s", "AImage")
}

// Constructor
func AImage0(drawer drawer.Drawer) *AImage {
	return &AImage{
		drawer: drawer,
	}
}
