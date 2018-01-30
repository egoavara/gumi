package gumi

import (
	"image"
)

type nMargin struct {
	SingleStructure
	b Blank
}

func (s *nMargin) draw(frame *image.RGBA) {
	s.child.draw(frame)
}
func (s *nMargin) size() Size {
	sz := s.child.size()

	hmin := sz.Horizontal.Min + s.b.L.Min + s.b.R.Min
	var hmax uint16
	if uint(sz.Horizontal.Max) + uint(s.b.L.Max) + uint(s.b.R.Max) > uint(AUTOLENGTH.Max){
		hmax = AUTOLENGTH.Max
	}else {
		hmax = sz.Horizontal.Max + s.b.L.Max + s.b.R.Max
	}


	vmin := sz.Vertical.Min + s.b.B.Min + s.b.T.Min
	var vmax uint16
	if uint(sz.Vertical.Max) + uint(s.b.B.Max) + uint(s.b.T.Max) > uint(AUTOLENGTH.Max){
		vmax = AUTOLENGTH.Max
	}else {
		vmax = sz.Vertical.Max + s.b.L.Max + s.b.R.Max
	}
	return Size{
		Length{vmin, vmax},
		Length{hmin, hmax},
	}
}
func helper(have int, l, a, b Length) (resl, resa, resb int) {
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
		// 최저길이만 만족가능\
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
func (s *nMargin) rect(rect image.Rectangle) {
	sz := s.child.size()
	//

	var w, l, _ = helper(rect.Dx(), sz.Horizontal, s.b.L, s.b.R)
	var h, _, t = helper(rect.Dy(), sz.Vertical, s.b.B, s.b.T)
	s.child.rect(image.Rect(
		rect.Min.X + l,
		rect.Min.Y + t,
		rect.Min.X + l + w,
		rect.Min.Y + t + h,
	))
}
func (s *nMargin) update(info *Information, style *Style) {
	s.child.update(info, style)
}
func (s *nMargin) Occur(event Event) {
	s.child.Occur(event)
}
func NMargin(sz Blank) *nMargin {
	return &nMargin{
		b: sz,
	}
}
func (s *nMargin) Set(sz Blank) {
	s.b = sz
}
func (s *nMargin) Get() Blank {
	return s.b
}


type nSize struct {
	SingleStructure
	sz Size
}

func (s *nSize) draw(frame *image.RGBA) {
	s.child.draw(frame)
}

func (s *nSize) size() Size {
	return s.sz
}

func (s *nSize) rect(r image.Rectangle) {
	s.child.rect(r)
}

func (s *nSize) update(info *Information, style *Style) {
	s.child.update(info, style)
}

func (s *nSize) Occur(event Event) {
	s.child.Occur(event)
}

func NSize(sz Size) *nSize {
	return &nSize{
		sz: sz,
	}
}


//
func (s *nSize) Set(sz Size) {
	s.sz = sz
}
func (s *nSize) Get() Size {
	return s.sz
}