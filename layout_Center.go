package gumi

import (
	"image"
	"fmt"
	"github.com/iamGreedy/gumi/gumre"
)

type LCenter struct {
	SingleStructure
}

func (s *LCenter) String() string {
	return fmt.Sprintf("%s", "LCenter")
}

func (s *LCenter) GUMIRender(frame *image.RGBA) {
	s.child.GUMIRender(frame)
}
func (s *LCenter) GUMISize() gumre.Size {
	return s.child.GUMISize()
}
func (s *LCenter) GUMIClip(r image.Rectangle) {
	sz := s.child.GUMISize()
	var vert, hori int
	if int(sz.Vertical.Max) < r.Dy(){
		vert = int(sz.Vertical.Max)
	}else {
		if int(sz.Vertical.Min) < r.Dy(){
			vert = int(sz.Vertical.Min)
		}else {
			vert = r.Dy()
		}
	}
	if int(sz.Horizontal.Max) < r.Dx(){
		hori = int(sz.Horizontal.Max)
	}else {
		if int(sz.Horizontal.Min) < r.Dx(){
			hori = int(sz.Horizontal.Min)
		}else {
			hori = r.Dx()
		}
	}
	left := (r.Dx() - hori) / 2 + r.Min.X
	top := (r.Dy() - vert) / 2+ r.Min.Y
	s.child.GUMIClip(image.Rect(left, top, left + hori, top + vert))
}
func (s *LCenter) GUMIUpdate(info *Information, style *Style) {
	s.child.GUMIUpdate(info, style)
}
func (s *LCenter) GUMIHappen(event Event) {
	s.child.GUMIHappen(event)
}

func LCenter0(elem GUMI) *LCenter {
	temp := &LCenter{

	}
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