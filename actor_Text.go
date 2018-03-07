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
	styleStore
	rendererStore
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
	s.rnode.SetRect(r)
}

// GUMIFunction / GUMIRender 				-> Define
func (s *AText) GUMIRender(frame *image.RGBA) {
	ctx := createContext(frame)
	s.style.useContext(ctx)
	defer s.style.releaseContext(ctx)
	ctx.SetColor(s.textColor)
	expectw, expecth := ctx.MeasureString(s.text)
	v, h := gumre.ParseAlign(s.align)
	var drawX, drawY float64
	switch v {
	case gumre.AlignBottom:
		drawY = float64(ctx.Height())
	case gumre.AlignVertical:
		drawY = float64(ctx.Height())/2 + expecth/2
	case gumre.AlignTop:
		drawY = expecth
	}
	switch h {
	case gumre.AlignRight:
		drawX = float64(ctx.Width()) - expectw
	case gumre.AlignHorizontal:
		drawX = float64(ctx.Width())/2 - expectw/2
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

// GUMITree / parent()						-> VoidNode::Default

// GUMITree / childrun()					-> VoidNode::Default

// GUMIRenderer / GUMIRenderSetup			-> Define
func (s *AText) GUMIRenderSetup(tree *drawer.RenderTree, parentnode *drawer.RenderNode) {
	s.rtree = tree
	s.rnode = tree.New(parentnode)
}


// GUMIRenderer / GUMIUpdate				-> Define
func (s *AText) GUMIUpdate() {
	if s.rnode.Check(){
		s.GUMIRender(s.rnode.SubImage())
		s.rnode.Complete()
	}
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
	temp := &AText{
		text:str,
		align:gumre.AlignCenter,
		textColor:color.White,
	}
	return temp
}

// Constructor 1
func AText1(str string, align gumre.Align) *AText {
	temp := &AText{
		text:str,
		align:align,
		textColor:color.White,
	}
	return temp
}

// Constructor 2
func AText2(str string, align gumre.Align, textColor color.Color) *AText {
	temp := &AText{
		text:str,
		align:align,
		textColor:textColor,
	}
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
	if s.text != text{
		s.text = text
		s.rnode.Require()
	}
}

// Method / GetText
func (s *AText) GetText() string {
	return s.text
}

// Method / SetAlign
func (s *AText) SetAlign(align gumre.Align) {
	if s.align != align{
		s.align = align
		s.rnode.Require()
	}
}

// Method / GetAlign
func (s *AText) GetAlign() gumre.Align {
	return s.align
}

// Method / SetColor
func (s *AText) SetColor(textColor color.Color) {
	r1,g1,b1,a1 := s.textColor.RGBA()
	r2,g2,b2,a2 := textColor.RGBA()
	if r1 != r2 || g1 != g2 || b1 != b2 || a1 != a2{
		s.textColor = textColor
		s.rnode.Require()
	}
}

// Method / GetColor
func (s *AText) GetColor() color.Color {
	return s.textColor
}
