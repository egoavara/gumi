package gumi

import (
	"github.com/fogleman/gg"
	"image"
	"fmt"
)

const (
	mtButtonMinWidth   = 40
	mtButtonMinHeight  = 20
	mtButtonAnimMillis = 100
)

type MTButton struct {
	//
	VoidStructure
	BoundStore
	StyleStore
	//
	mtColorSingle
	//
	handle float64
	anim   float64
	//
	text string
	//
	cursorEnter, active bool
	onClick             MTButtonClick
	onFocus             MTButtonFocus
}

func (s *MTButton) String() string {
	return fmt.Sprintf("%s", "MTButton")
}
type MTButtonFocus func(self *MTButton, focus bool)
type MTButtonClick func(self *MTButton)

func (s *MTButton) draw(frame *image.RGBA) {
	var ctx = GGContextRGBASub(frame, s.bound)
	var w, h = float64(ctx.Width()), float64(ctx.Height())
	radius := h / 2
	s.style.useContext(ctx)
	defer s.style.releaseContext(ctx)
	//
	if s.active{
		ctx.SetColor(s.style.Material.PalletteColor(s.mtColorSingle.mcl1)[0])
		ctx.DrawArc(radius, radius, radius, gg.Radians(90), gg.Radians(270))
		ctx.DrawRectangle(radius, 0, w-radius*2, h)
		ctx.DrawArc(w-radius, radius, radius, gg.Radians(-90), gg.Radians(90))
		ctx.Fill()
		//
		ctx.SetColor(s.style.Material.PalletteColor(s.mtColorSingle.mcl1)[1])
		txtw, txth := ctx.MeasureString(s.text)
		ctx.DrawString(s.text, (w-txtw)/2, (h+txth)/2)
		ctx.Stroke()
	}else {
		ctx.SetColor(
			phaseColor(s.style.Material.PalletteColor(s.mtColorSingle.mcl1)[0], s.style.Material.PalletteColor(s.mtColorSingle.mcl1)[1], s.anim),
		)
		ctx.DrawArc(radius, radius, radius, gg.Radians(90), gg.Radians(270))
		ctx.DrawRectangle(radius, 0, w-radius*2, h)
		ctx.DrawArc(w-radius, radius, radius, gg.Radians(-90), gg.Radians(90))
		ctx.Fill()
		//
		ctx.SetColor(
			phaseColor(s.style.Material.PalletteColor(s.mtColorSingle.mcl1)[1], s.style.Material.PalletteColor(s.mtColorSingle.mcl1)[0], s.anim),
		)
		txtw, txth := ctx.MeasureString(s.text)
		ctx.DrawString(s.text, (w-txtw)/2, (h+txth)/2)
		ctx.Stroke()
	}
}
func (s *MTButton) size() Size {
	return Size{
		Vertical:   MinLength(mtButtonMinHeight),
		Horizontal: MinLength(mtButtonMinWidth),
	}
}
func (s *MTButton) rect(r image.Rectangle) {
	s.bound = r
}
func (s *MTButton) update(info *Information, style *Style) {
	s.style = style
	if s.cursorEnter {
		if s.handle < mtButtonAnimMillis {
			s.handle = s.handle + float64(info.Dt)
			if s.handle > mtButtonAnimMillis {
				s.handle = mtButtonAnimMillis
			}
			s.anim = Animation.Material.Button(s.handle / mtButtonAnimMillis)
		}
	} else {
		if s.handle > 0 {
			s.handle = s.handle - float64(info.Dt)
			if s.handle < 0 {
				s.handle = 0
			}
			s.anim = 1 - Animation.Material.Button((mtButtonAnimMillis-s.handle)/mtButtonAnimMillis)
		}
	}
}
func (s *MTButton) Occur(event Event) {
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
				s.onFocus(s,true)
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

func MTButton0(mcl MaterialColor, text string, onclick MTButtonClick) *MTButton {
	temp := &MTButton{
		text:    text,
		onClick: onclick,
	}
	temp.SetMaterialColor(mcl)
	return temp
}
func MTButton1(text string, onclick MTButtonClick) *MTButton {
	return &MTButton{
		text:    text,
		onClick: onclick,
	}
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
