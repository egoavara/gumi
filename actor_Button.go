package gumi

import (
	"image"
	"github.com/fogleman/gg"
	"image/draw"
)

type aButton struct {
	VoidStructure
	BoundStore
	StyleStore
	cursorEnter bool
	text  string
}

func (s *aButton) draw(frame *image.RGBA) {

	var ctx = GGContextRGBASub(frame, s.bound)
	var w, h = float64(ctx.Width()), float64(ctx.Height())
	//
	s.style.useContext(ctx)
	defer s.style.releaseContext(ctx)
	expectw, expecth := ctx.MeasureString(s.text)
	radius := float64(s.bound.Dy() / 2)
	//
	ok, clr := IsColorImage(s.style.Face)
	if ok{
		ctx.SetColor(clr)
		ctx.DrawArc(radius, radius, radius, gg.Radians(90), gg.Radians(270))
		ctx.DrawRectangle(radius, 0 ,w - radius * 2, h)
		ctx.DrawArc(w - radius, radius, radius, gg.Radians(-90), gg.Radians(90))
		ctx.Fill()
	}else {
		draw.Draw(frame.SubImage(s.bound).(*image.RGBA), s.bound.Intersect(s.style.Face.Bounds()), s.style.Face, s.style.Face.Bounds().Min, draw.Over)
	}
	_, clr = IsColorImage(s.style.Line)
	ctx.SetColor(clr)
	ctx.DrawString(s.text, w/2 - expectw/2, h/2 + expecth/2 - 1)
	//
}
const aBUTTONPADDING = 5
func (s *aButton) size() Size {
	s.style.Font.Use()
	defer s.style.Font.Release()

	h, v := s.style.Font.CalculateSize(s.text)
	temp := Size{
		Horizontal: MinLength(uint16(h) + aBUTTONPADDING * 2),
		Vertical: MinLength(uint16(v) + aBUTTONPADDING * 2),
	}

	return temp
}
func (s *aButton) rect(r image.Rectangle) {
	s.bound = r
}
func (s *aButton) update(info *Information, style *Style) {
	s.style = style
}
func (s *aButton) Occur(event Event) {
	switch ev := event.(type) {
	case EventKeyPress:
		if ev.Key == KEY_MOUSE1{

		}
	case EventKeyRelease:
		if ev.Key == KEY_MOUSE1{

		}
	case EventCursor:
		x := int(ev.X)
		y := int(ev.Y)
		if (s.bound.Min.X <= x && x < s.bound.Max.X) && (s.bound.Min.Y <= y && y < s.bound.Max.Y){
			s.cursorEnter = true
		}else {
			s.cursorEnter = false
		}
	}
}

//
func AButton(str string) *aButton {
	return &aButton{
		text:  str,
	}
}
func (s *aButton) SetText(str string) {
	s.text = str
}
func (s *aButton) GetText() string {
	return s.text
}