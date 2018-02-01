package gumi

import (
	"golang.org/x/image/math/fixed"
	"image"
	"fmt"
)

type AText struct {
	VoidStructure
	BoundStore
	StyleStore
	align Align
	text  string
}

func (s *AText) String() string {
	return fmt.Sprintf("%s(text:%s)", "AText", s.text)
}

func (s *AText) draw(frame *image.RGBA) {
	s.style.Default.Font.Use()
	defer s.style.Default.Font.Release()
	s.style.Default.Font.ChangeSource(s.style.Default.Line)
	expectw, expecth := s.style.Default.Font.CalculateSize(s.text)
	v, h := ParseAlign(s.align)
	var dot fixed.Point26_6
	switch v {
	case Align_BOTTOM:
		dot.Y = fixed.I(s.bound.Max.Y)
	case Align_VCENTER:
		dot.Y = fixed.I(s.bound.Min.Y + (s.bound.Dy()/2) + (expecth)/2)
	case Align_TOP:
		dot.Y = fixed.I(0 + expecth)
	}
	switch h {
	case Align_RIGHT:
		dot.X = fixed.I(s.bound.Max.X - expectw)
	case Align_HCENTER:
		dot.X = fixed.I(s.bound.Min.X + (s.bound.Dx()/2) - expectw/2)
	case Align_LEFT:
		dot.X = fixed.I(0)
	}
	s.style.Default.Font.Draw(s.bound, frame, s.text, dot)
}
func (s *AText) size() Size {
	s.style.Default.Font.Use()
	defer s.style.Default.Font.Release()

	h, v := s.style.Default.Font.CalculateSize(s.text)

	temp := Size{
		Horizontal: MinLength(uint16(h)),
		Vertical: MinLength(uint16(v)),
	}

	return temp
}
func (s *AText) rect(r image.Rectangle) {
	s.bound = r
}
func (s *AText) update(info *Information, style *Style) {
	s.style = style

}
func (s *AText) Occur(event Event) {
}

//
func AText0(str string, align Align) *AText {
	return &AText{
		text:  str,
		align: align,
	}
}
func AText1(str string) *AText {
	return &AText{
		text:  str,
		align: Align_CENTER,
	}
}
func (s *AText) Set(str string) {
	s.text = str
}
func (s *AText) Get() string {
	return s.text
}

func (s *AText) SetAlign(align Align) {
	s.align = align
}
func (s *AText) GetAlign() Align {
	return s.align
}
