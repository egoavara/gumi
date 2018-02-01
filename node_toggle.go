package gumi

import (
	"github.com/fogleman/gg"
	"image"
	"image/color"
	"image/draw"
	"fmt"
)

type NToggle struct {
	SingleStructure
	BoundStore
	StyleStore
	//
	cursorEnter, active bool
	//
	onActive NToggleActive
}

func (s *NToggle) String() string {
	return fmt.Sprintf("%s(active:%v)", "NToggle", s.active)
}

type NToggleActive func(active bool)

func (s *NToggle) draw(frame *image.RGBA) {
	var ctx = GGContextRGBASub(frame, s.bound)
	var w, h = float64(ctx.Width()), float64(ctx.Height())
	//
	radius := float64(s.bound.Dy() / 2)
	//
	var ok bool
	var clr color.Color
	if s.active {
		_, clr = IsColorImage(s.style.Default.Line)
		ctx.SetColor(clr)
		ctx.DrawArc(radius, radius, radius, gg.Radians(90), gg.Radians(270))
		ctx.DrawRectangle(radius, 0, w-radius*2, h-1)
		ctx.DrawArc(w-radius, radius, radius, gg.Radians(-90), gg.Radians(90))
		ctx.Fill()
	} else {
		ok, clr = IsColorImage(s.style.Default.Face)
		if ok {
			ctx.SetColor(clr)
			ctx.DrawArc(radius, radius, radius, gg.Radians(90), gg.Radians(270))
			ctx.DrawRectangle(radius, 0, w-radius*2, h-1)
			ctx.DrawArc(w-radius, radius, radius, gg.Radians(-90), gg.Radians(90))
			ctx.Fill()
		} else {
			draw.Draw(frame.SubImage(s.bound).(*image.RGBA), s.bound.Intersect(s.style.Default.Face.Bounds()), s.style.Default.Face, s.style.Default.Face.Bounds().Min, draw.Over)
		}
		if s.cursorEnter {
			_, clr = IsColorImage(s.style.Default.Line)
			ctx.SetColor(clr)
			ctx.DrawLine(radius, 0, w-radius, 0)
			ctx.Stroke()
			ctx.DrawArc(radius, radius, radius, gg.Radians(90), gg.Radians(270))
			ctx.Stroke()
			ctx.DrawLine(radius, h-1, w-radius, h-1)
			ctx.Stroke()
			ctx.DrawArc(w-radius, radius, radius, gg.Radians(-90), gg.Radians(90))
			ctx.Stroke()
		}
	}
	//
	s.child.draw(frame)
}

func (s *NToggle) size() Size {
	sz := s.child.size()
	sz.Vertical.Min += aBUTTONPADDING * 2
	sz.Horizontal.Min += aBUTTONPADDING * 2
	if !(sz.Horizontal.Min < sz.Horizontal.Max) {
		sz.Horizontal.Max = sz.Horizontal.Min
	}
	if !(sz.Vertical.Min < sz.Vertical.Max) {
		sz.Vertical.Max = sz.Vertical.Min
	}
	return sz
}
func (s *NToggle) rect(r image.Rectangle) {
	s.bound = r
	s.child.rect(image.Rect(
		r.Min.X+aBUTTONPADDING-1,
		r.Min.Y+aBUTTONPADDING-1,
		r.Max.X-aBUTTONPADDING+1,
		r.Max.Y-aBUTTONPADDING+1,
	))
}
func (s *NToggle) update(info *Information, style *Style) {
	s.style = style
	s.child.update(info, style)
}
func (s *NToggle) Occur(event Event) {
	switch ev := event.(type) {
	case EventKeyPress:
	case EventKeyRelease:
		if ev.Key == KEY_MOUSE1 {
			if s.cursorEnter {
				s.active = !s.active
				if s.onActive != nil{
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
	s.child.Occur(event)
}

//
func NToggle0(active NToggleActive) *NToggle {
	return &NToggle{
		onActive: active,
	}
}

func (s *NToggle) OnActive(callback NToggleActive) {
	s.onActive = callback
}
func (s *NToggle) ReferActive() NToggleActive {
	return s.onActive
}
