package gumi

import (
	"fmt"
	"github.com/iamGreedy/gumi/drawer"
	"github.com/iamGreedy/gumi/gumre"
	"image"
)

// Actor::Spacer
//
// Aspacer is responsible for creating the margins.
// It can also be created with a combination of NMargin and AEmpty,
// But it exists separately for convenience.
type ASpacer struct {
	VoidNode
	//
	verical    gumre.Length
	horizontal gumre.Length
}

// GUMIFunction / GUMIInit 					-> VoidNode::Default

// GUMIFunction / GUMIInfomation 			-> Define::Empty
func (s ASpacer) GUMIInfomation(info Information) {
}

// GUMIFunction / GUMIStyle 			-> Define::Empty
func (s ASpacer) GUMIStyle(style *Style) {
}

// GUMIFunction / GUMIClip 				-> Define::Empty
func (ASpacer) GUMIClip(r image.Rectangle) {
}

// GUMIFunction / GUMIRender 			-> Define::Empty
func (ASpacer) GUMIRender(frame *image.RGBA) {

}

// GUMIFunction / GUMISize 				-> Define
func (s *ASpacer) GUMISize() gumre.Size {
	return gumre.Size{
		Vertical:   s.verical,
		Horizontal: s.horizontal,
	}
}

// GUMITree / born 							-> VoidNode::Default

// GUMITree / breed 						-> VoidNode::Default

// GUMITree / Parent						-> VoidNode::Default

// GUMITree / Childrun						-> VoidNode::Default

// GUMIRenderer / GUMIRenderSetup			-> Define
func (s ASpacer) GUMIRenderSetup(frame *image.RGBA, tree *drawer.RenderTree, parentnode *drawer.RenderNode) {
}

// GUMIRenderer / GUMIDraw					-> Define
func (s ASpacer) GUMIDraw() {

}

// GUMIRenderer / GUMIUpdate				-> Define
func (s ASpacer) GUMIUpdate() {
}

// GUMIEventer / GUMIHappen					-> Define
func (ASpacer) GUMIHappen(event Event) {
}

// fmt.Stringer / String					-> Define
func (ASpacer) String() string {
	return fmt.Sprintf("%s", "ASpacer")
}

// Constructor 0
func ASpacer0(horizontal, vertical gumre.Length) *ASpacer {
	temp := &ASpacer{}
	temp.Set(horizontal, vertical)
	return temp
}

// Constructor 1
func ASpacer1(horizontal gumre.Length) *ASpacer {
	temp := &ASpacer{}
	temp.Set(horizontal, gumre.AUTOLENGTH)
	return temp
}

// Constructor 2
func ASpacer2(vertical gumre.Length) *ASpacer {
	temp := &ASpacer{}
	temp.Set(gumre.AUTOLENGTH, vertical)
	return temp
}

// Method / Get -> GetHorizontal(), GetVertical()
func (s *ASpacer) Get() (horizontal, vertical gumre.Length) {
	return s.horizontal, s.verical
}

// Method / Set -> SetHorizontal(...), SetVertical(...)
func (s *ASpacer) Set(horizontal, vertical gumre.Length) {
	s.horizontal, s.verical = horizontal, vertical
}

// Method / GetHorizontal
func (s *ASpacer) GetHorizontal() gumre.Length {
	return s.horizontal
}

// Method / SetHorizontal
func (s *ASpacer) SetHorizontal(horizontal gumre.Length) {
	s.horizontal = horizontal
}

// Method / GetVertical
func (s *ASpacer) GetVertical() gumre.Length {
	return s.verical
}

// Method / SetVertical
func (s *ASpacer) SetVertical(vertical gumre.Length) {
	s.verical = vertical
}
