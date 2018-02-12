package gumi

import (
	"fmt"
	"github.com/iamGreedy/gumi/drawer"
	"github.com/iamGreedy/gumi/gumre"
	"image"
)

type LCenter struct {
	SingleStructure
}

func (s *LCenter) GUMIInfomation(info Information) {
	s.child.GUMIInfomation(info)
}
func (s *LCenter) GUMIStyle(style *Style) {
	s.child.GUMIStyle(style)
}
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
func (s *LCenter) GUMIRender(frame *image.RGBA) {
}
func (s *LCenter) GUMIDraw(frame *image.RGBA) {
	s.child.GUMIDraw(frame)
}

func (s *LCenter) GUMIRenderTree(tree *drawer.RenderTree, parentnode *drawer.RenderNode) {
	panic("implement me")
}
func (s *LCenter) GUMIUpdate() {
	panic("implement me")
}

func (s *LCenter) GUMIHappen(event Event) {
	s.child.GUMIHappen(event)
}
func (s *LCenter) GUMISize() gumre.Size {
	return s.child.GUMISize()
}
func (s *LCenter) String() string {
	return fmt.Sprintf("%s", "LCenter")
}

func LCenter0(elem GUMI) *LCenter {
	temp := &LCenter{}
	elem.born(temp)
	temp.breed([]GUMI{elem})
	return temp
}

func (s *LCenter) LoadElement() GUMI {
	return s.child
}
func (s *LCenter) SaveElement(elem GUMI) {
	s.child = elem
}
