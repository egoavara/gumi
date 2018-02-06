package gumi

import (
	"fmt"
	"image"
)

const (
	mtRadioMinWidth                  = 20
	mtRadioMinHeight                 = 20
	mtRadioAnimationOnOffDeltaMillis = 200
	mtRadioInnerRadiusDifference = 3
)

const (
	mtRadioAnimationOnOff  = iota
	mtRadioAnimationLength = iota
)

type MTRadio struct {
	//
	VoidStructure
	boundStore
	styleStore
	//
	mtColorFromTo
	studio *AnimationStudio
	//
	cursorEnter, active bool
	onActive            MTRadioActive
}
type MTRadioActive func(self *MTRadio, active bool)

func (s *MTRadio) String() string {
	return fmt.Sprintf("%s(active:%v)", "MTRadio", s.active)
}
func (s *MTRadio) init() {
	s.studio = NewAnimationStudio(mtRadioAnimationLength)
	s.studio.Set(mtRadioAnimationOnOff, &AnimationPercent{
		Delta: Animation.DeltaByMillis(mtRadioAnimationOnOffDeltaMillis),
		Fn:    Material.DefaultAnimation.Radio,
	})
}
func (s *MTRadio) draw(frame *image.RGBA) {
	var ctx = GGContextRGBASub(frame, s.bound)
	var w, h = float64(ctx.Width()), float64(ctx.Height())
	var baseColor0 = s.GetFromMaterialColor().BaseColor()
	var mainColor1 = s.GetToMaterialColor().MainColor()
	var onoff = s.studio.Get(mtRadioAnimationOnOff).(*AnimationPercent)
	var radius = h / 2
	var innerRadius = radius - mtRadioInnerRadiusDifference
	//
	ctx.SetColor(baseColor0)
	ctx.DrawCircle(w/2, h/2, radius)
	ctx.Fill()
	//
	ctx.SetColor(Scale.Color(baseColor0, mainColor1, onoff.Current))
	ctx.DrawCircle(w/2, h/2, innerRadius)
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
		s.studio.Get(mtRadioAnimationOnOff).(*AnimationPercent).Request(1)
	} else {
		s.studio.Get(mtRadioAnimationOnOff).(*AnimationPercent).Request(0)
	}
	s.studio.Animate(info)
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
func MTRadio0(active MTRadioActive) *MTRadio {
	temp := &MTRadio{
		onActive: active,
	}
	temp.SetFromMaterialColor(Material.Pallette.White)
	temp.SetToMaterialColor(Material.Pallette.White)
	return temp
}
func MTRadio1(from, to *MaterialColor, active MTRadioActive) *MTRadio {
	temp := &MTRadio{
		onActive: active,
	}
	temp.SetFromMaterialColor(from)
	temp.SetToMaterialColor(to)
	return temp
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
