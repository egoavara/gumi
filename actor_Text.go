package gumi

import (
	"golang.org/x/image/math/fixed"
	"image"
	"sync"
)

type aText struct {
	VoidStructure
	BoundStore
	StyleStore
	align Align
	text  string
	lk    *sync.RWMutex
}

func (s *aText) draw(frame *image.RGBA) {
	s.lk.Lock()
	defer s.lk.Unlock()
	s.style.Font.Use()
	defer s.style.Font.Release()
	s.style.Font.ChangeSource(s.style.Line)

	expectw, expecth := s.style.Font.CalculateSize(s.text)
	v, h := ParseAlign(s.align)
	var dot fixed.Point26_6
	switch v {
	case Align_BOTTOM:
		dot.Y = fixed.I(s.bound.Max.Y)
	case Align_VCENTER:
		dot.Y = fixed.I(s.bound.Max.Y/2 + (expecth)/2)
	case Align_TOP:
		dot.Y = fixed.I(0 + expecth)
	}
	switch h {
	case Align_RIGHT:
		dot.X = fixed.I(s.bound.Max.X - expectw)
	case Align_HCENTER:
		dot.X = fixed.I(s.bound.Max.X/2 - expectw/2)
	case Align_LEFT:
		dot.X = fixed.I(0)
	}
	s.style.Font.Draw(s.bound, frame, s.text, dot)
}
func (s *aText) size() Size {
	s.lk.RLock()
	defer s.lk.RUnlock()
	s.style.Font.Use()
	defer s.style.Font.Release()

	h, v := s.style.Font.CalculateSize(s.text)

	temp := Size{
		Horizontal: MinLength(uint16(h)),
		Vertical: MinLength(uint16(v)),
	}

	return temp
}
func (s *aText) rect(r image.Rectangle) {
	s.bound = r
}
func (s *aText) update(info *Information, style *Style) {
	s.style = style
}
func (s *aText) Occur(event Event) {
}

//
func AText(str string, align Align) *aText {
	return &aText{
		text:  str,
		align: align,
		lk:    new(sync.RWMutex),
	}
}
func (s *aText) Set(str string) {
	s.lk.Lock()
	defer s.lk.Unlock()
	s.text = str
}
func (s *aText) Get() string {
	s.lk.RLock()
	defer s.lk.RUnlock()
	return s.text
}

func (s *aText) Align(align Align) {
	s.lk.Lock()
	defer s.lk.Unlock()
	s.align = align
}
func (s *aText) Alignment() Align {
	s.lk.RLock()
	defer s.lk.RUnlock()
	return s.align
}
