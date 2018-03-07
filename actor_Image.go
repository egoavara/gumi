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
	rendererStore
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
	s.rnode.SetRect(rect)
}

// GUMIFunction / GUMIRender 		-> Define
func (s *AImage) GUMIRender(frame *image.RGBA) {
	s.drawer.Draw(frame)
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

// GUMITree / parent()						-> VoidNode::Default

// GUMITree / childrun()					-> VoidNode::Default

// GUMIRenderer / GUMIRenderSetup			-> Define
func (s *AImage) GUMIRenderSetup(tree *drawer.RenderTree, parentnode *drawer.RenderNode) {
	s.rtree = tree
	s.rnode = tree.New(parentnode)
}

// GUMIRenderer / GUMIDraw					-> Define
func (s *AImage) GUMIDraw() {
	s.GUMIRender(s.rnode.SubImage())
	s.rnode.Complete()
}

// GUMIRenderer / GUMIUpdate					-> Define
func (s *AImage) GUMIUpdate() {
	if s.rnode.Check(){
		s.GUMIDraw()
	}
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
