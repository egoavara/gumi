package gumi

import (
	"fmt"
	"github.com/fogleman/gg"
	"image"
	"github.com/iamGreedy/gumi/gumre"
)

const (
	mtButtonMinPadding = 5
)
const (
	mtButtonAnimationHover = iota
	//
	mtButtonAnimationLength = iota
)

type MTButton struct {
	//
	VoidStructure
	boundStore
	styleStore
	//
	mtColorSingle
	studio *gumre.Studio
	hover *gumre.Percenting
	//
	text string
	//
	cursorEnter, active bool
	onClick             MTButtonClick
	onFocus             MTButtonFocus
}
type MTButtonFocus func(self *MTButton, focus bool)
type MTButtonClick func(self *MTButton)

func (s *MTButton) String() string {
	return fmt.Sprintf("%s", "MTButton")
}
func (s *MTButton) GUMIInit() {
	s.studio = gumre.Animation.Studio(mtButtonAnimationLength)
	s.hover = s.studio.Set(mtButtonAnimationHover, &gumre.Percenting{
		Delta:gumre.Animation.PercentingByMillis(250),
		Fn: Material.DefaultAnimation.Button,
	}).(*gumre.Percenting)
}
func (s *MTButton) GUMIRender(frame *image.RGBA) {
	var ctx = createContextRGBASub(frame, s.bound)
	var w, h = float64(ctx.Width()), float64(ctx.Height())
	var baseColor, mainColor = s.GetMaterialColor().Color()
	radius := h / 2
	s.style.useContext(ctx)
	defer s.style.releaseContext(ctx)
	//

	if s.active {
		ctx.SetColor(baseColor)
		// background
		ctx.DrawArc(radius, radius, radius, gg.Radians(90), gg.Radians(270))
		ctx.DrawRectangle(radius, 0, w-radius*2, h)
		ctx.DrawArc(w-radius, radius, radius, gg.Radians(-90), gg.Radians(90))
		ctx.Fill()
		// text
		ctx.SetColor(mainColor)
		txtw, txth := ctx.MeasureString(s.text)
		ctx.DrawString(s.text, (w-txtw)/2, (h+txth)/2-1)
		ctx.Stroke()
	} else {

		ctx.SetColor(Scale.Color(baseColor, mainColor, s.hover.Value()))
		// background
		ctx.DrawArc(radius, radius, radius, gg.Radians(90), gg.Radians(270))
		ctx.DrawRectangle(radius, 0, w-radius*2, h)
		ctx.DrawArc(w-radius, radius, radius, gg.Radians(-90), gg.Radians(90))
		ctx.Fill()
		// text
		ctx.SetColor(Scale.Color(mainColor, baseColor, s.hover.Value()))
		txtw, txth := ctx.MeasureString(s.text)
		ctx.DrawString(s.text, (w-txtw)/2, (h+txth)/2-1)
		ctx.Stroke()
	}
}
func (s *MTButton) GUMISize() gumre.Size {

	s.style.Default.Font.Use()
	defer s.style.Default.Font.Release()
	var hori, vert = s.style.Default.Font.CalculateSize(s.text)
	//
	return gumre.Size{
		Vertical:   gumre.MinLength(uint16(vert + mtButtonMinPadding*2)),
		Horizontal: gumre.MinLength(uint16(hori + mtButtonMinPadding*2)),
	}
}
func (s *MTButton) GUMIClip(r image.Rectangle) {
	s.bound = r
}
func (s *MTButton) GUMIUpdate(info *Information, style *Style) {
	s.style = style
	if s.cursorEnter {
		s.hover.Request(1)
	} else {
		s.hover.Request(0)
	}
	s.studio.Animate(float64(info.Dt))
}
func (s *MTButton) GUMIHappen(event Event) {
	switch ev := event.(type) {
	case EventKeyPress:
		if ev.Key == KEY_MOUSE1 {
			if s.cursorEnter {
				s.active = true
			}
		}
	case EventKeyRelease:
		if ev.Key == KEY_MOUSE1 {
			if s.active {
				if s.onClick != nil {
					s.onClick(s)
				}
				s.active = false
			}
		}
	case EventCursor:

		x := int(ev.X)
		y := int(ev.Y)
		if (s.bound.Min.X <= x && x < s.bound.Max.X) && (s.bound.Min.Y <= y && y < s.bound.Max.Y) {
			if s.onFocus != nil {
				s.onFocus(s, true)
			}
			s.cursorEnter = true
		} else {
			if s.onFocus != nil {
				s.onFocus(s, false)
			}
			s.cursorEnter = false
		}
	}
}

func MTButton0(text string, onclick MTButtonClick) *MTButton {
	temp := &MTButton{
		text:    text,
		onClick: onclick,
	}
	temp.SetMaterialColor(Material.Pallette.White)
	return temp
}
func MTButton1(mcl *MaterialColor, text string, onclick MTButtonClick) *MTButton {
	temp := &MTButton{
		text:    text,
		onClick: onclick,
	}
	temp.SetMaterialColor(mcl)
	return temp
}

func (s *MTButton) SetText(txt string) {
	s.text = txt
}
func (s *MTButton) GetText() string {
	return s.text
}
func (s *MTButton) OnClick(callback MTButtonClick) {
	s.onClick = callback
}
func (s *MTButton) ReferClick() MTButtonClick {
	return s.onClick
}
func (s *MTButton) OnFocus(callback MTButtonClick) {
	s.onClick = callback
}
func (s *MTButton) ReferFocus() MTButtonClick {
	return s.onClick
}
