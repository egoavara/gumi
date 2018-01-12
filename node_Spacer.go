package gumi

import (
	"image"
)

type nMargin struct {
	GUMILINK_SINGLE
	b Blank
	s Size
}

func (s *nMargin) size(drawing *Drawing, style *Style) Size {
	temp := s.child.(GUMIElem).size(drawing, style)
	if temp.Vertical.Min < s.s.Vertical.Min {
		temp.Vertical.Min = s.s.Vertical.Min
	}
	if temp.Horizontal.Min < s.s.Horizontal.Min {
		temp.Horizontal.Min = s.s.Horizontal.Min
	}
	if temp.Vertical.Max > s.s.Vertical.Max {
		temp.Vertical.Max = s.s.Vertical.Max
	}
	if temp.Horizontal.Max > s.s.Horizontal.Max {
		temp.Horizontal.Max = s.s.Horizontal.Max
	}
	temp.Vertical.Min += s.b.VMin()
	temp.Horizontal.Min += s.b.HMin()
	return temp
}
func (s *nMargin) draw(drawing *Drawing, style *Style, frame Frame) {
	sz := s.child.(GUMIElem).size(drawing, style)
	rectangle := frame.Bounds()
	//
	var v, h uint16 = 0, 0
	var l, t uint16 = 0, 0
	//var r, b = 0, 0
	l, _, h = size_help(
		uint16(rectangle.Dx()),
		sz.Horizontal,
		s.b.L,
		s.b.R,
	)
	t, _, v = size_help(
		uint16(rectangle.Dy()),
		sz.Vertical,
		s.b.T,
		s.b.B,
	)
	//

	s.child.(GUMIElem).draw(drawing, style, frame.SubFrame(
		image.Rect(int(l), int(t), int(l+h), int(t+v)),
	))
}

//
func NMargin(s Size, sz Blank) *nMargin {
	return &nMargin{
		s: s,
		b: sz,
	}
}
func (s *nMargin) Set(sz Size) {
	s.s = sz
}
func (s *nMargin) Get() Size {
	return s.s
}
func (s *nMargin) SetMargin(sz Blank) {
	s.b = sz
}
func (s *nMargin) GetMargin() Blank {
	return s.b
}

//
func size_help(length uint16, elemLength Length, a, b Length) (resa, resb, resl uint16) {
	var ulength = uint(length)
	var uelemlMin = uint(elemLength.Min)
	var uelemlMax = uint(elemLength.Max)
	var uaMin = uint(a.Min)
	var uaMax = uint(a.Max)
	var ubMin = uint(b.Min)
	var ubMax = uint(b.Max)

	//
	if ulength >= uelemlMax+uaMax+ubMax {
		// 모두 만족할 정도로 공간이 존재
		resa = a.Max
		resb = b.Max
		resl = elemLength.Max
	} else if ulength >= uelemlMax+uaMin+ubMin {
		// 엘레먼트 크기는 최대만족하나 패딩은 최소만족만 하는 경우
		resa, resb = Proportion(length-elemLength.Max, a.Max-a.Min, b.Max-b.Min)
		resl = elemLength.Max
	} else if ulength >= uelemlMin+uaMin+ubMin {
		// 엘레먼트 크기도 최소만족, 패딩도 최소만족하는 경우
		resa, resb = a.Min, b.Min
		resl = length - (a.Min + b.Min)
	} else {
		// 그 무엇도 최소만족 못함
		resa, resb = 0, 0
		resl = length
	}
	return
}

// TODO: nPadding
