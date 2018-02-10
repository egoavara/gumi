package gumi

import (
	"fmt"
	"image"
	"github.com/iamGreedy/gumi/gumre"
)

type NSize struct {
	SingleStructure
	sz gumre.Size
}

func (s *NSize) String() string {
	return fmt.Sprintf("%s(GUMISize:%v)", "NSize", s.sz)
}
func (s *NSize) GUMIRender(frame *image.RGBA) {
	s.child.GUMIRender(frame)
}
func (s *NSize) GUMISize() gumre.Size {
	temp := s.sz
	c := s.child.GUMISize()
	if temp.Vertical == gumre.AUTOLENGTH {
		temp.Vertical = c.Vertical
	} else if temp.Vertical == gumre.MINLENGTH {
		temp.Vertical = c.Vertical
		temp.Vertical.Max = c.Vertical.Min
	} else if temp.Vertical == gumre.MAXLENGTH {
		temp.Vertical = c.Vertical
		temp.Vertical.Min = c.Vertical.Max
	}
	if temp.Horizontal == gumre.AUTOLENGTH {
		temp.Horizontal = c.Horizontal
	} else if temp.Horizontal == gumre.MINLENGTH {
		temp.Horizontal = c.Horizontal
		temp.Horizontal.Max = c.Horizontal.Min
	} else if temp.Horizontal == gumre.MINLENGTH {
		temp.Horizontal = c.Horizontal
		temp.Horizontal.Min = c.Horizontal.Max
	}
	return temp
}
func (s *NSize) GUMIClip(r image.Rectangle) {
	s.child.GUMIClip(r)
}
func (s *NSize) GUMIUpdate(info *Information, style *Style) {
	s.child.GUMIUpdate(info, style)
}
func (s *NSize) GUMIHappen(event Event) {
	s.child.GUMIHappen(event)
}

func NSize0(sz gumre.Size) *NSize {
	return &NSize{
		sz: sz,
	}
}

//
func (s *NSize) Set(sz gumre.Size) {
	s.sz = sz
}
func (s *NSize) Get() gumre.Size {
	return s.sz
}
