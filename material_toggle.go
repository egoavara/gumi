package gumi

import (
	"github.com/fogleman/gg"
	"image"
	"image/color"
)

const mtToggleMinWidth = 30
const mtToggleMinHeight = 20
const mtToggleAnimMillis  = 700
type mtToggle struct {
	GUMI
	//
	VoidStructure
	BoundStore
	StyleStore
	//
	handle float64
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
	ctx.SetColor(
		phaseColor(s.style.Material.PalletteColor(White)[0], s.style.Material.PalletteColor(Green)[0], s.handle),
	)
	ctx.DrawArc(radius, radius, radius, gg.Radians(90), gg.Radians(270))
	ctx.DrawRectangle(radius, 0, w-radius*2, h-1)
	ctx.DrawArc(w-radius, radius, radius, gg.Radians(-90), gg.Radians(90))
	ctx.Fill()
	ctx.SetColor(s.style.Material.PalletteColor(White)[1])
	ctx.DrawCircle(phasePos(w, radius, s.handle), radius, miniradius)
	ctx.Fill()
}
func phaseColor(c1, c2 color.Color, handle float64) color.Color {
	r1, g1, b1, a1 := c1.RGBA()
	r2, g2, b2, a2 := c2.RGBA()
	var temp = float64(handle) / 256
	var res color.RGBA
	res.R = uint8(float64(r2 - r1) * temp)
	res.G = uint8(float64(g2 - g1) * temp)
	res.B = uint8(float64(b2 - b1) * temp)
	res.A = uint8(float64(a2 - a1) * temp)
	return res


}
func phasePos(width, radius float64, handle float64) float64 {
	return radius + (width - radius * 2) * float64(handle) / 256
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
	if s.active{
		if s.handle == 1 {
			s.handle = 1
		}
	}else {

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
func MTToggle(active MTToggleActive) *mtToggle {
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
