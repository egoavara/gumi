package gumi

import (
	"fmt"
	"image"
	"image/color"
	"github.com/iamGreedy/gumi/gumre"
	"github.com/iamGreedy/gumi/drawer"
)

type AText struct {
	VoidStructure
	boundStore
	styleStore
	//
	align     gumre.Align
	text      string
	textColor color.Color
	//
}

func (s *AText) GUMIInfomation(info Information) {
}

func (s *AText) GUMIStyle(style *Style) {
	s.style = style
}

func (s *AText) GUMIClip(r image.Rectangle) {
	s.bound = r
}

func (s *AText) GUMIRender(frame *image.RGBA) {
	ctx := createContextRGBASub(frame, s.bound)
	s.style.useContext(ctx)
	defer s.style.releaseContext(ctx)
	ctx.SetColor(s.textColor)
	expectw, expecth := ctx.MeasureString(s.text)
	v, h := gumre.ParseAlign(s.align)
	var drawX, drawY float64
	switch v {
	case gumre.AlignBottom:
		drawY = float64(s.bound.Dy())
	case gumre.AlignVertical:
		drawY = float64(s.bound.Dy())/2 + expecth/2
	case gumre.AlignTop:
		drawY = expecth
	}
	switch h {
	case gumre.AlignRight:
		drawX = float64(s.bound.Dx()) - expectw
	case gumre.AlignHorizontal:
		drawX = float64(s.bound.Dx())/2 - expectw/2
	case gumre.AlignLeft:
		drawX = 0
	}
	ctx.DrawString(s.text, drawX, drawY - 1)
}

func (s *AText) GUMIDraw(frame *image.RGBA) {
	s.GUMIRender(frame)
}
func (s *AText) GUMIRenderTree(tree *drawer.RenderTree, parentnode *drawer.RenderNode) {
	panic("implement me")
}

func (s *AText) GUMIUpdate() {
	panic("implement me")
}


func (s *AText) GUMIHappen(event Event) {
}
func (s *AText) GUMISize() gumre.Size {
	s.style.Default.Font.Use()
	defer s.style.Default.Font.Release()

	h, v := s.style.Default.Font.CalculateSize(s.text)

	temp := gumre.Size{
		Horizontal: gumre.MinLength(uint16(h)),
		Vertical:   gumre.MinLength(uint16(v)),
	}

	return temp
}
func (s *AText) String() string {
	return fmt.Sprintf("%s(text:%s)", "AText", s.text)
}

//
func AText0(str string) *AText {
	return &AText{
		text:      str,
		align:     gumre.AlignCenter,
		textColor: color.White,
	}
}
func AText1(str string, align gumre.Align) *AText {
	return &AText{
		text:      str,
		align:     align,
		textColor: color.White,
	}
}
func AText2(str string, align gumre.Align, textColor color.Color) *AText {
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
func (s *AText) SetAlign(align gumre.Align) {
	s.align = align
}
func (s *AText) GetAlign() gumre.Align {
	return s.align
}
func (s *AText) SetColor(textColor color.Color) {
	s.textColor = textColor
}
func (s *AText) GetColor() color.Color {
	return s.textColor
}
