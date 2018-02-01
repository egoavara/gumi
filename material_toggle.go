package gumi

import (
	"github.com/fogleman/gg"
	"image"
)

const mtToggleMinWidth = 40
const mtToggleMinHeight = 20
const mtToggleAnimMillis = 200

type mtToggle struct {
	//
	VoidStructure
	BoundStore
	StyleStore
	//
	mtColorFromTo
	//
	handle float64
	anim float64
	//
	cursorEnter, active bool
	onActive            MTToggleActive
}
type MTToggleActive func(active bool)

func (s *mtToggle) draw(frame *image.RGBA) {
	var ctx = GGContextRGBASub(frame, s.bound)
	var w, h = float64(ctx.Width()), float64(ctx.Height())

	radius := h / 2
	miniradius := radius - 3
	//
	mcl1 := s.style.Material.PalletteColor(s.mcl1)
	mcl2 := s.style.Material.PalletteColor(s.mcl2)
	//
	ctx.SetColor(
		phaseColor(mcl1[0], mcl2[1], s.anim),
	)
	//ctx.SetColor(color.RGBA{94, 97, 97, 255})
	ctx.DrawArc(radius, radius, radius, gg.Radians(90), gg.Radians(270))
	ctx.DrawRectangle(radius, 0, w-radius*2, h)
	ctx.DrawArc(w-radius, radius, radius, gg.Radians(-90), gg.Radians(90))
	ctx.Fill()
	ctx.SetColor(mcl1[1])
	//ctx.SetColor(color.RGBA{213, 217, 255, 255})
	ctx.DrawCircle(radius + phasePos(w - 2 * radius , s.anim), radius, miniradius)
	ctx.Fill()
}

func (s *mtToggle) size() Size {
	return Size{
		Vertical:   MinLength(mtToggleMinHeight),
		Horizontal: MinLength(mtToggleMinWidth),
	}
}
func (s *mtToggle) rect(r image.Rectangle) {
	s.bound = r
}
func (s *mtToggle) update(info *Information, style *Style) {
	s.style = style
	if s.active {
		if s.handle < mtToggleAnimMillis {
			s.handle = s.handle + float64(info.Dt)
			if s.handle > mtToggleAnimMillis{
				s.handle = mtToggleAnimMillis
			}
			s.anim = Animation.Material.Toggle(s.handle / mtToggleAnimMillis)
		}
	} else {
		if s.handle > 0 {
			s.handle = s.handle - float64(info.Dt)
			if s.handle < 0{
				s.handle = 0
			}
			s.anim = 1 - Animation.Material.Toggle((mtToggleAnimMillis - s.handle) / mtToggleAnimMillis)
		}
	}
}
func (s *mtToggle) Occur(event Event) {
	switch ev := event.(type) {
	case EventKeyPress:
	case EventKeyRelease:
		if ev.Key == KEY_MOUSE1 {
			if s.cursorEnter {
				s.active = !s.active
				if s.onActive != nil {
					s.onActive(s.active)
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
func MTToggle(from, to MaterialColor, active MTToggleActive) *mtToggle {
	temp := &mtToggle{
		onActive: active,
	}
	temp.SetFromMaterialColor(from)
	temp.SetToMaterialColor(to)
	return temp
}
func MTToggle1(active MTToggleActive) *mtToggle {
	return &mtToggle{
		onActive: active,
	}
}
func (s *mtToggle) OnActive(callback MTToggleActive) {
	s.onActive = callback
}
func (s *mtToggle) ReferActive() MTToggleActive {
	return s.onActive
}
