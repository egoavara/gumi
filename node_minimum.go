package gumi

import (
	"fmt"
	"image"
)
type NMinimum struct {
	SingleStructure
	sz Size
}

func (s *NMinimum) String() string {
	return fmt.Sprintf("%s(size:%v)", "NMinimum", s.sz)
}
func (s *NMinimum) draw(frame *image.RGBA) {
	s.child.draw(frame)
}
func (s *NMinimum) size() Size {
	temp := s.sz
	c := s.child.size()
	if temp.Vertical == AUTOLENGTH{
		temp.Vertical = c.Vertical
	}
	if temp.Horizontal == AUTOLENGTH{
		temp.Horizontal = c.Horizontal
	}
	return temp
}
func (s *NMinimum) rect(r image.Rectangle) {
	s.child.rect(r)
}
func (s *NMinimum) update(info *Information, style *Style) {
	s.child.update(info, style)
}
func (s *NMinimum) Occur(event Event) {
	s.child.Occur(event)
}

func NMinimum0(sz Size) *NMinimum {
	return &NMinimum{
		sz: sz,
	}
}


//
func (s *NMinimum) Set(sz Size) {
	s.sz = sz
}
func (s *NMinimum) Get() Size {
	return s.sz
}