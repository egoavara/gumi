package gumi

import (
	"fmt"
	"github.com/iamGreedy/gumi/drawer"
	"github.com/iamGreedy/gumi/gumre"
	"image"
)

// Actor::Empty
//
// AEmpty exists only for the GUMI Tree as an element that does nothing
type AEmpty struct {
	VoidNode
}

// GUMIFunction / GUMIInit 					-> VoidNode::Default

// GUMIFunction / GUMIInfomation 			-> Define::Empty
func (s AEmpty) GUMIInfomation(info Information) {
}

// GUMIFunction / GUMIStyle 			-> Define::Empty
func (s AEmpty) GUMIStyle(style *Style) {
}

// GUMIFunction / GUMIClip 			-> Define::Empty
func (AEmpty) GUMIClip(image.Rectangle) {
}

// GUMIFunction / GUMISize 			-> Define
func (AEmpty) GUMISize() gumre.Size {
	return gumre.Size{
		gumre.AUTOLENGTH,
		gumre.AUTOLENGTH,
	}
}

// GUMITree / born 							-> VoidNode::Default

// GUMITree / breed 						-> VoidNode::Default

// GUMITree / Parent()						-> VoidNode::Default

// GUMITree / Childrun()					-> VoidNode::Default

// GUMIRenderer / GUMIRenderSetup			-> Define::Empty
func (s AEmpty) GUMIRenderSetup(frame *image.RGBA, tree *drawer.RenderTree, parentnode *drawer.RenderNode) {
}

// GUMIRenderer / GUMIRender				-> Define::Empty
func (s AEmpty) GUMIUpdate() {
}

// GUMIRenderer / GUMIDraw					-> Define::Empty
func (s AEmpty) GUMIDraw() {
}

// GUMIEventer / GUMIHappen					-> Define
func (AEmpty) GUMIHappen(event Event) {
}

// fmt.Stringer / String					-> Define
func (s AEmpty) String() string {
	return fmt.Sprintf("%s", "AEmpty")
}

// Constructor
func AEmpty0() *AEmpty {
	return &AEmpty{}
}
