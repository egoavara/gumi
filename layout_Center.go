package gumi

import (
	"fmt"
	"github.com/iamGreedy/gumi/drawer"
	"github.com/iamGreedy/gumi/gumre"
	"image"
)

// Layout::Center
//
// Make all child center
type LCenter struct {
	SingleNode
	//

}

// GUMIFunction / GUMIInit 					-> SingleNode::Default

// GUMIFunction / GUMIInfomation 			-> Define
func (s *LCenter) GUMIInfomation(info Information) {
	s.child.GUMIInfomation(info)
}

// GUMIFunction / GUMIStyle 				-> Define
func (s *LCenter) GUMIStyle(style *Style) {
	s.child.GUMIStyle(style)
}

// GUMIFunction / GUMIClip 					-> Define
func (s *LCenter) GUMIClip(r image.Rectangle) {
	sz := s.child.GUMISize()
	var vert, hori int
	if int(sz.Vertical.Max) < r.Dy() {
		vert = int(sz.Vertical.Max)
	} else {
		if int(sz.Vertical.Min) < r.Dy() {
			vert = int(sz.Vertical.Min)
		} else {
			vert = r.Dy()
		}
	}
	if int(sz.Horizontal.Max) < r.Dx() {
		hori = int(sz.Horizontal.Max)
	} else {
		if int(sz.Horizontal.Min) < r.Dx() {
			hori = int(sz.Horizontal.Min)
		} else {
			hori = r.Dx()
		}
	}
	left := (r.Dx()-hori)/2 + r.Min.X
	top := (r.Dy()-vert)/2 + r.Min.Y
	s.child.GUMIClip(image.Rect(left, top, left+hori, top+vert))
}

// GUMIFunction / GUMIRender 				-> Define::Empty
func (s *LCenter) GUMIRender(frame *image.RGBA) {
}

// GUMIFunction / GUMISize 					-> Define
func (s *LCenter) GUMISize() gumre.Size {
	return s.child.GUMISize()
}

// GUMITree / born 							-> SingleNode::Default

// GUMITree / breed 						-> SingleNode::Default

// GUMITree / parent()						-> SingleNode::Default

// GUMITree / childrun()					-> SingleNode::Default

// GUMIRenderer / GUMIRenderSetup 			-> Define::Empty
func (s *LCenter) GUMIRenderSetup(tree *drawer.RenderTree, parentnode *drawer.RenderNode) {
	s.child.GUMIRenderSetup(tree, parentnode)
}

// GUMIRenderer / GUMIUpdate 				-> Define
func (s *LCenter) GUMIUpdate() {
	s.child.GUMIUpdate()
}


// GUMIEventer / GUMIHappen					-> Define
func (s *LCenter) GUMIHappen(event Event) {
	s.child.GUMIHappen(event)
}

// fmt.Stringer / String					-> Define
func (s *LCenter) String() string {
	return fmt.Sprintf("%s", "LCenter")
}

// Constructor 0
func LCenter0(elem GUMI) *LCenter {
	temp := &LCenter{}
	elem.born(temp)
	temp.breed([]GUMI{elem})
	return temp
}

// Get Elements
func (s *LCenter) LoadElement() GUMI {
	return s.child
}

// Set Elements
func (s *LCenter) SaveElement(elem GUMI) {
	s.child = elem
}
