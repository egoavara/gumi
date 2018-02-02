package gumi

import (
	"fmt"
	"image"
)

const mtRadioMinWidth = 20
const mtRadioMinHeight = 20
const mtRadioAnimMillis = 200

type MTRadio struct {
	//
	VoidStructure
	BoundStore
	StyleStore
	//
	mtColorFromTo
	//
	handle float64
	anim   float64
	//
	cursorEnter, active bool
	onActive            MTRadioActive
}
type MTRadioActive func(self *MTRadio, active bool)

func (s *MTRadio) String() string {
	return fmt.Sprintf("%s(active:%v)", "MTRadio", s.active)
}
func (s *MTRadio) draw(frame *image.RGBA) {
	var ctx = GGContextRGBASub(frame, s.bound)
	var w, h = float64(ctx.Width()), float64(ctx.Height())

	radius := h / 2
	miniradius := radius - 3
	//
	mcl1 := s.style.Material.PalletteColor(s.mcl1)
	mcl2 := s.style.Material.PalletteColor(s.mcl2)
	//
	ctx.SetColor(mcl1[0])
	ctx.DrawCircle(w/2, h/2, radius)
	ctx.Fill()
	//
	ctx.SetColor(
		phaseColor(mcl1[0], mcl2[1], s.anim),
	)
	ctx.DrawCircle(w/2, h/2, miniradius)
	ctx.Fill()
}
func (s *MTRadio) size() Size {
	return Size{
		Vertical:   FixLength(mtRadioMinHeight),
		Horizontal: FixLength(mtRadioMinWidth),
	}
}
func (s *MTRadio) rect(r image.Rectangle) {
	s.bound = r
}
func (s *MTRadio) update(info *Information, style *Style) {
	s.style = style
	if s.active {
		if s.handle < mtRadioAnimMillis {
			s.handle = s.handle + float64(info.Dt)
			if s.handle > mtRadioAnimMillis {
				s.handle = mtRadioAnimMillis
			}
			s.anim = Animation.Material.Toggle(s.handle / mtRadioAnimMillis)
		}
	} else {
		if s.handle > 0 {
			s.handle = s.handle - float64(info.Dt)
			if s.handle < 0 {
				s.handle = 0
			}
			s.anim = 1 - Animation.Material.Toggle((mtRadioAnimMillis-s.handle)/mtRadioAnimMillis)
		}
	}
}
func (s *MTRadio) Occur(event Event) {
	switch ev := event.(type) {
	case EventKeyPress:
	case EventKeyRelease:
		if ev.Key == KEY_MOUSE1 {
			if s.cursorEnter {
				s.active = !s.active
				if s.onActive != nil {
					s.onActive(s, s.active)
				}
			}
		}
	case EventCursor:
		x := int(ev.X)
		y := int(ev.Y)
		if (s.bound.Min.X <= x && x < s.bound.Max.X) && (s.bound.Min.Y <= y && y < s.bound.Max.Y) {
			s.cursorEnter = true
		} else {
			s.cursorEnter = false
		}
	}
}

//
func MTRadio0(from, to MaterialColor, active MTRadioActive) *MTRadio {
	temp := &MTRadio{
		onActive: active,
	}
	temp.SetFromMaterialColor(from)
	temp.SetToMaterialColor(to)
	return temp
}
func MTRadio1(active MTRadioActive) *MTRadio {
	return &MTRadio{
		onActive: active,
	}
}

func (s *MTRadio) Get() bool {
	return s.GetActive()
}
func (s *MTRadio) Set(active bool) {
	s.SetActive(active)
}
func (s *MTRadio) GetActive() bool {
	return s.active
}
func (s *MTRadio) SetActive(active bool) {
	s.active = active
}
func (s *MTRadio) OnActive(callback MTRadioActive) {
	s.onActive = callback
}
func (s *MTRadio) ReferActive() MTRadioActive {
	return s.onActive
}
