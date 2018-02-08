package gumi

import (
	"fmt"
	"image"
	"image/color"
)

type AText struct {
	VoidStructure
	boundStore
	styleStore
	//
	align     Align
	text      string
	textColor color.Color
	//
}

func (s *AText) String() string {
	return fmt.Sprintf("%s(text:%s)", "AText", s.text)
}

func (s *AText) draw(frame *image.RGBA) {
	ctx := GGContextRGBASub(frame, s.bound)
	s.style.useContext(ctx)
	defer s.style.releaseContext(ctx)
	ctx.SetColor(s.textColor)

	expectw, expecth := ctx.MeasureString(s.text)
	v, h := ParseAlign(s.align)
	var drawX, drawY float64
	switch v {
	case Align_BOTTOM:
		drawY = float64(s.bound.Dy())
	case Align_VCENTER:
		drawY = float64(s.bound.Dy())/2 + expecth/2
	case Align_TOP:
		drawY = expecth
	}
	switch h {
	case Align_RIGHT:
		drawX = float64(s.bound.Dx()) - expectw
	case Align_HCENTER:
		drawX = float64(s.bound.Dx())/2 - expectw/2
	case Align_LEFT:
		drawX = 0
	}
	ctx.DrawString(s.text, drawX, drawY - 3)
}
func (s *AText) size() Size {
	s.style.Default.Font.Use()
	defer s.style.Default.Font.Release()

	h, v := s.style.Default.Font.CalculateSize(s.text)

	temp := Size{
		Horizontal: MinLength(uint16(h)),
		Vertical:   MinLength(uint16(v)),
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
func AText0(str string) *AText {
	return &AText{
		text:      str,
		align:     Align_CENTER,
		textColor: color.White,
	}
}
func AText1(str string, align Align) *AText {
	return &AText{
		text:      str,
		align:     align,
		textColor: color.White,
	}
}
func AText2(str string, align Align, textColor color.Color) *AText {
	return &AText{
		text:      str,
		align:     align,
		textColor: textColor,
	}
}

func (s *AText) Set(text string) {
	s.SetText(text)
}
func (s *AText) Get() string {
	return s.GetText()
}
func (s *AText) SetText(text string) {
	s.text = text
}
func (s *AText) GetText() string {
	return s.text
}
func (s *AText) SetAlign(align Align) {
	s.align = align
}
func (s *AText) GetAlign() Align {
	return s.align
}
func (s *AText) SetColor(textColor color.Color) {
	s.textColor = textColor
}
func (s *AText) GetColor() color.Color {
	return s.textColor
}
