package gumi

import (
	"github.com/fogleman/gg"
	"image"
)

const (
	mtButtonMinWidth   = 40
	mtButtonMinHeight  = 20
	mtButtonAnimMillis = 100
)

type mtButton struct {
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
type MTButtonFocus func(focus bool)
type MTButtonClick func()

func (s *mtButton) draw(frame *image.RGBA) {
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

func (s *mtButton) size() Size {
	return Size{
		Vertical:   MinLength(mtButtonMinHeight),
		Horizontal: MinLength(mtButtonMinWidth),
	}
}

func (s *mtButton) rect(r image.Rectangle) {
	s.bound = r
}

func (s *mtButton) update(info *Information, style *Style) {
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

func (s *mtButton) Occur(event Event) {
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
					s.onClick()
				}
				s.active = false
			}
		}
	case EventCursor:
		x := int(ev.X)
		y := int(ev.Y)
		if (s.bound.Min.X <= x && x < s.bound.Max.X) && (s.bound.Min.Y <= y && y < s.bound.Max.Y) {
			if s.onFocus != nil {
				s.onFocus(true)
			}
			s.cursorEnter = true
		} else {
			if s.onFocus != nil {
				s.onFocus(false)
			}
			s.cursorEnter = false
		}
	}
}

func MTButton(mcl MaterialColor, text string, onclick MTButtonClick) *mtButton {
	temp := &mtButton{
		text:    text,
		onClick: onclick,
	}
	temp.SetMaterialColor(mcl)
	return temp
}
func MTButton1(text string, onclick MTButtonClick) *mtButton {
	return &mtButton{
		text:    text,
		onClick: onclick,
	}
}

func (s *mtButton) SetText(txt string) {
	s.text = txt
}
func (s *mtButton) GetText() string {
	return s.text
}
func (s *mtButton) OnClick(callback MTButtonClick) {
	s.onClick = callback
}
func (s *mtButton) ReferClick() MTButtonClick {
	return s.onClick
}
func (s *mtButton) OnFocus(callback MTButtonClick) {
	s.onClick = callback
}
func (s *mtButton) ReferFocus() MTButtonClick {
	return s.onClick
}
