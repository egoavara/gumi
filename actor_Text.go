package gumi

import (
	"golang.org/x/image/math/fixed"
	"sync"
)

type aText struct {
	GUMILINK_EMPTY
	align Align
	text  string
	lk *sync.RWMutex
}

func (s *aText) size(drawing *Drawing, style *Style) Size {
	s.lk.RLock()
	defer s.lk.RUnlock()
	style.Font.Use()
	defer style.Font.Release()

	h, v := style.Font.CalculateSize(s.text)
	temp := Size{
		Horizontal: Length{
			Min: uint16(h),
			Max: AUTOLENGTH.Max,
		},
		Vertical: Length{
			Min: uint16(v),
			Max: AUTOLENGTH.Max,
		},
	}

	return temp
}
func (s *aText) draw(drawing *Drawing, style *Style, frame Frame) {
	s.lk.Lock()
	defer s.lk.Unlock()
	style.Font.Use()
	defer style.Font.Release()
	style.Font.ChangeSource(style.Line)
	framebound := frame.Bounds()
	expectw, expecth := style.Font.CalculateSize(s.text)
	v, h := ParseAlign(s.align)
	var dot fixed.Point26_6
	switch v {
	case Align_BOTTOM:
		dot.Y = fixed.I(framebound.Max.Y)
	case Align_VCENTER:
		dot.Y = fixed.I(framebound.Max.Y/2 + (expecth)/2)
	case Align_TOP:
		dot.Y = fixed.I(0 + expecth)
	}
	switch h {
	case Align_RIGHT:
		dot.X = fixed.I(framebound.Max.X - expectw)
	case Align_HCENTER:
		dot.X = fixed.I(framebound.Max.X/2 - expectw/2)
	case Align_LEFT:
		dot.X = fixed.I(0)
	}
	//
	//dot.X += fixed.I(1)
	//dot.Y += fixed.I(1)
	style.Font.Draw(framebound, frame, s.text, dot)
}

//
func AText(str string, align Align) *aText {
	return &aText{
		text:  str,
		align: align,
		lk: new(sync.RWMutex),
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
