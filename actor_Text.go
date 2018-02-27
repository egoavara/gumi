package gumi

import (
	"fmt"
	"image"
	"image/color"
	"github.com/iamGreedy/gumi/gumre"
	"github.com/iamGreedy/gumi/drawer"
)

// Actor::Text
//
// AText use for render text
type AText struct {
	VoidNode
	boundStore
	styleStore
	frameStore
	//
	align     gumre.Align
	text      string
	textColor color.Color
	//
}

// GUMIFunction / GUMIInfomation 			-> Define::Empty
func (s *AText) GUMIInfomation(info Information) {
}

// GUMIFunction / GUMIStyle 				-> Define
func (s *AText) GUMIStyle(style *Style) {
	s.style = style
}

// GUMIFunction / GUMIClip 					-> Define
func (s *AText) GUMIClip(r image.Rectangle) {
	s.bound = r
}

// GUMIFunction / GUMIRender 				-> Define
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

// GUMIFunction / GUMISize 					-> Define
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

// GUMITree / born 							-> VoidNode::Default

// GUMITree / breed 						-> VoidNode::Default

// GUMITree / Parent()						-> VoidNode::Default

// GUMITree / Childrun()					-> VoidNode::Default

// GUMIRenderer / GUMIRenderSetup			-> Define
func (s *AText) GUMIRenderSetup(frame *image.RGBA, tree *drawer.RenderTree, parentnode *drawer.RenderNode) {
	s.frame = frame
	// TODO : Cache
}

// GUMIRenderer / GUMIDraw					-> Define
func (s *AText) GUMIDraw() {
	s.GUMIRender(s.frame)
}

// GUMIRenderer / GUMIUpdate				-> Define
func (s *AText) GUMIUpdate() {
	// TODO
	panic("implement me")
}

// GUMIEventer / GUMIHappen					-> Define
func (s *AText) GUMIHappen(event Event) {
}

// fmt.Stringer / String					-> Define
func (s *AText) String() string {
	return fmt.Sprintf("%s(text:%s)", "AText", s.text)
}

// Constructor 0
func AText0(str string) *AText {
	temp := &AText{}
	temp.SetText(str)
	temp.SetAlign(gumre.AlignCenter)
	temp.SetColor(color.White)
	return temp
}

// Constructor 1
func AText1(str string, align gumre.Align) *AText {
	temp := &AText{}
	temp.SetText(str)
	temp.SetAlign(align)
	temp.SetColor(color.White)
	return temp
}

// Constructor 2
func AText2(str string, align gumre.Align, textColor color.Color) *AText {
	temp := &AText{}
	temp.SetText(str)
	temp.SetAlign(align)
	temp.SetColor(textColor)
	return temp
}

// Method / Set -> SetText(...)
func (s *AText) Set(text string) {
	s.SetText(text)
}

// Method / Get -> GetText()
func (s *AText) Get() string {
	return s.GetText()
}

// Method / SetText
func (s *AText) SetText(text string) {
	s.text = text
}

// Method / GetText
func (s *AText) GetText() string {
	return s.text
}

// Method / SetAlign
func (s *AText) SetAlign(align gumre.Align) {
	s.align = align
}

// Method / GetAlign
func (s *AText) GetAlign() gumre.Align {
	return s.align
}

// Method / SetColor
func (s *AText) SetColor(textColor color.Color) {
	s.textColor = textColor
}

// Method / GetColor
func (s *AText) GetColor() color.Color {
	return s.textColor
}
