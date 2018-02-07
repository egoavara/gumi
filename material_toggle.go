package gumi

import (
	"fmt"
	"github.com/fogleman/gg"
	"image"
)

const (
	mtToggleMinWidth                  = 40
	mtToggleMinHeight                 = 20
	mtToggleAnimationOnOffDeltaMillis = 200
	mtToggleInnerRadiusDifference = 3
)

const (
	mtToggleAnimationOnOff  = iota
	mtToggleAnimationLength = iota
)
type MTToggle struct {
	//
	VoidStructure
	boundStore
	styleStore
	//
	mtColorFromTo
	studio *AnimationStudio
	onoff *AnimationPercent
	//
	cursorEnter, active bool
	onActive            MTToggleActive
}

// Event Callbacks
type MTToggleActive func(self *MTToggle, active bool)

// GUMI Structure
func (s *MTToggle) String() string {
	return fmt.Sprintf("%s(active:%v)", "MTToggle", s.active)
}
func (s *MTToggle) init() {
	s.studio = NewAnimationStudio(mtToggleAnimationLength)
	s.onoff = s.studio.Set(mtToggleAnimationOnOff, &AnimationPercent{
		Delta: Animation.DeltaByMillis(mtToggleAnimationOnOffDeltaMillis),
		Fn:    Material.DefaultAnimation.Toggle,
	}).(*AnimationPercent)
}
func (s *MTToggle) draw(frame *image.RGBA) {
	var ctx = GGContextRGBASub(frame, s.bound)
	var w, h = float64(ctx.Width()), float64(ctx.Height())
	var baseColor0, mainColor0 = s.GetFromMaterialColor().Color()
	var mainColor1 = s.GetToMaterialColor().MainColor()
	var radius = h / 2
	var innerRadius = radius - mtToggleInnerRadiusDifference
	//
	ctx.SetColor(Scale.Color(baseColor0, mainColor1, s.onoff.Value()), )
	//ctx.SetColor(color.RGBA{94, 97, 97, 255})
	ctx.DrawArc(radius, radius, radius, gg.Radians(90), gg.Radians(270))
	ctx.DrawRectangle(radius, 0, w-radius*2, h)
	ctx.DrawArc(w-radius, radius, radius, gg.Radians(-90), gg.Radians(90))
	ctx.Fill()
	ctx.SetColor(mainColor0)
	//ctx.SetColor(color.RGBA{213, 217, 255, 255})
	ctx.DrawCircle(radius+Scale.Length(w-2*radius, s.onoff.Value()), radius, innerRadius)
	ctx.Fill()
}
func (s *MTToggle) size() Size {
	return Size{
		Vertical:   MinLength(mtToggleMinHeight),
		Horizontal: MinLength(mtToggleMinWidth),
	}
}
func (s *MTToggle) rect(r image.Rectangle) {
	s.bound = r
}
func (s *MTToggle) update(info *Information, style *Style) {
	s.style = style
	if s.active {
		s.onoff.Request(1)
	} else {
		s.onoff.Request(0)
	}
	s.studio.Animate(info)
}
func (s *MTToggle) Occur(event Event) {
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

// Constructors
func MTToggle0(active MTToggleActive) *MTToggle {
	temp := &MTToggle{
		onActive: active,
	}
	temp.SetFromMaterialColor(Material.Pallette.White)
	temp.SetToMaterialColor(Material.Pallette.White)
	return temp
}
func MTToggle1(from, to *MaterialColor, active MTToggleActive) *MTToggle {
	temp := &MTToggle{
		onActive: active,
	}
	temp.SetFromMaterialColor(from)
	temp.SetToMaterialColor(to)
	return temp
}


// Element Property
func (s *MTToggle) Get() bool {
	return s.GetActive()
}
func (s *MTToggle) Set(active bool) {
	s.SetActive(active)
}
func (s *MTToggle) GetActive() bool {
	return s.active
}
func (s *MTToggle) SetActive(active bool) {
	s.active = active
}
func (s *MTToggle) OnActive(callback MTToggleActive) {
	s.onActive = callback
}
func (s *MTToggle) ReferActive() MTToggleActive {
	return s.onActive
}
