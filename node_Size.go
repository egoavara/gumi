package gumi

import (
	"fmt"
	"image"
	"github.com/iamGreedy/gumi/gumre"
	"github.com/iamGreedy/gumi/drawer"
)

type NSize struct {
	SingleNode
	sz gumre.Size
}

func (s *NSize) GUMIInfomation(info Information) {
	s.child.GUMIInfomation(info)
}
func (s *NSize) GUMIStyle(style *Style) {
	s.child.GUMIStyle(style)
}
func (s *NSize) GUMIClip(r image.Rectangle) {
	s.child.GUMIClip(r)
}
func (s *NSize) GUMIRender(frame *image.RGBA) {

}
func (s *NSize) GUMIDraw(frame *image.RGBA) {
	s.child.GUMIDraw(frame)
}

func (s *NSize) GUMIRenderTree(tree *drawer.RenderTree, parentnode *drawer.RenderNode) {
	panic("implement me")
}
func (s *NSize) GUMIUpdate() {
	panic("implement me")
}

func (s *NSize) GUMIHappen(event Event) {
	s.child.GUMIHappen(event)
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
func (s *NSize) String() string {
	return fmt.Sprintf("%s(GUMISize:%v)", "NSize", s.sz)
}

func NSize0(sz gumre.Size) *NSize {
	return &NSize{
		sz: sz,
	}
}

func (s *NSize) Get() gumre.Size {
	return s.GetSize()
}
func (s *NSize) Set(sz gumre.Size) {
	s.SetSize(sz)
}
func (s *NSize) GetSize() gumre.Size {
	return s.sz
}
func (s *NSize) SetSize(sz gumre.Size) {
	s.sz = sz
}
