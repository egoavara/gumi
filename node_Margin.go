package gumi

import (
	"image"
	"fmt"
	"github.com/iamGreedy/gumi/gumre"
	"github.com/iamGreedy/gumi/drawer"
)

type NMargin struct {
	SingleStructure
	b gumre.Blank
}

func (s *NMargin) GUMIInfomation(info Information) {
	s.child.GUMIInfomation(info)
}
func (s *NMargin) GUMIStyle(style *Style) {
	s.child.GUMIStyle(style)
}
func (s *NMargin) GUMIClip(rect image.Rectangle) {
	sz := s.child.GUMISize()
	//

	var w, l, _ = calcMargin(rect.Dx(), sz.Horizontal, s.b.L, s.b.R)
	var h, _, t = calcMargin(rect.Dy(), sz.Vertical, s.b.B, s.b.T)
	s.child.GUMIClip(image.Rect(
		rect.Min.X + l,
		rect.Min.Y + t,
		rect.Min.X + l + w,
		rect.Min.Y + t + h,
	))
}
func (s *NMargin) GUMIRender(frame *image.RGBA) {

}
func (s *NMargin) GUMIDraw(frame *image.RGBA) {
	s.child.GUMIDraw(frame)
}

func (s *NMargin) GUMIRenderTree(tree *drawer.RenderTree, parentnode *drawer.RenderNode) {
	panic("implement me")
}
func (s *NMargin) GUMIUpdate() {
	panic("implement me")
}

func (s *NMargin) GUMIHappen(event Event) {
	s.child.GUMIHappen(event)
}
func (s *NMargin) GUMISize() gumre.Size {
	sz := s.child.GUMISize()

	hmin := sz.Horizontal.Min + s.b.L.Min + s.b.R.Min
	var hmax uint16
	if uint(sz.Horizontal.Max) + uint(s.b.L.Max) + uint(s.b.R.Max) > uint(gumre.AUTOLENGTH.Max){
		hmax = gumre.AUTOLENGTH.Max
	}else {
		hmax = sz.Horizontal.Max + s.b.L.Max + s.b.R.Max
	}


	vmin := sz.Vertical.Min + s.b.B.Min + s.b.T.Min
	var vmax uint16
	if uint(sz.Vertical.Max) + uint(s.b.B.Max) + uint(s.b.T.Max) > uint(gumre.AUTOLENGTH.Max){
		vmax = gumre.AUTOLENGTH.Max
	}else {
		vmax = sz.Vertical.Max + s.b.L.Max + s.b.R.Max
	}
	return gumre.Size{
		gumre.Length{vmin, vmax},
		gumre.Length{hmin, hmax},
	}
}
func (s *NMargin) String() string {
	return fmt.Sprintf("%s(margin:%v)", "NMargin", s.b)
}
//
func calcMargin(have int, l, a, b gumre.Length) (resl, resa, resb int) {
	if int(l.Max) + int(a.Max) + int(b.Max) <= have{
		// 최대값도 만족 가능
		resl = int(l.Max)
		resa = int(a.Max)
		resb = int(b.Max)
	}else if int(l.Max) + int(a.Min) + int(b.Min) <= have{
		// 최대길이 만족, 최대여백 불가
		resl = int(l.Max)
		temp := have - resl
		resa = (temp)/(int(a.Min) + int(b.Min)) * int(a.Min)
		resb = temp - resa
	}else if int(l.Min) + int(a.Min) + int(b.Min) <= have{
		// 최저길이만 만족가능
		resa = int(a.Min)
		resb = int(b.Min)
		resl = have - resa - resb
	}else if int(l.Min) <= have{
		resl = int(l.Min)
		temp := have - resl
		resa = (temp)/(int(a.Min) + int(b.Min)) * int(a.Min)
		resb = temp - resa
	}else {
		resl = have
	}
	return
}
//
func NMargin0(sz gumre.Blank) *NMargin {
	return &NMargin{
		b: sz,
	}
}

func (s *NMargin) Set(sz gumre.Blank) {
	s.SetMargin(sz)
}
func (s *NMargin) Get() gumre.Blank {
	return s.GetMargin()
}
func (s *NMargin) SetMargin(sz gumre.Blank) {
	s.b = sz
}
func (s *NMargin) GetMargin() gumre.Blank {
	return s.b
}


