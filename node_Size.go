package gumi

import (
	"fmt"
	"image"
)

type NSize struct {
	SingleStructure
	sz Size
}

func (s *NSize) String() string {
	return fmt.Sprintf("%s(size:%v)", "NSize", s.sz)
}
func (s *NSize) draw(frame *image.RGBA) {
	s.child.draw(frame)
}
func (s *NSize) size() Size {
	temp := s.sz
	c := s.child.size()
	if temp.Vertical == AUTOLENGTH {
		temp.Vertical = c.Vertical
	} else if temp.Vertical == MINLENGTH {
		temp.Vertical = c.Vertical
		temp.Vertical.Max = c.Vertical.Min
	} else if temp.Vertical == MAXLENGTH {
		temp.Vertical = c.Vertical
		temp.Vertical.Min = c.Vertical.Max
	}
	if temp.Horizontal == AUTOLENGTH {
		temp.Horizontal = c.Horizontal
	} else if temp.Horizontal == MINLENGTH {
		temp.Horizontal = c.Horizontal
		temp.Horizontal.Max = c.Horizontal.Min
	} else if temp.Horizontal == MINLENGTH {
		temp.Horizontal = c.Horizontal
		temp.Horizontal.Min = c.Horizontal.Max
	}
	return temp
}
func (s *NSize) rect(r image.Rectangle) {
	s.child.rect(r)
}
func (s *NSize) update(info *Information, style *Style) {
	s.child.update(info, style)
}
func (s *NSize) Occur(event Event) {
	s.child.Occur(event)
}

func NSize0(sz Size) *NSize {
	return &NSize{
		sz: sz,
	}
}

//
func (s *NSize) Set(sz Size) {
	s.sz = sz
}
func (s *NSize) Get() Size {
	return s.sz
}
