package gumi

import (
	"image"
	"fmt"
)

type LCenter struct {
	SingleStructure
}

func (s *LCenter) String() string {
	return fmt.Sprintf("%s", "LCenter")
}

func (s *LCenter) draw(frame *image.RGBA) {
	s.child.draw(frame)
}
func (s *LCenter) size() Size {
	return s.child.size()
}
func (s *LCenter) rect(r image.Rectangle) {
	sz := s.child.size()
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
	s.child.rect(image.Rect(left, top, left + hori, top + vert))
}
func (s *LCenter) update(info *Information, style *Style) {
	s.child.update(info, style)
}
func (s *LCenter) Occur(event Event) {
	s.child.Occur(event)
}

func LCenter0(elem GUMI) *LCenter {
	temp := &LCenter{

	}
	elem.Born(temp)
	temp.Breed([]GUMI{elem})
	return temp
}
