package gumi

import (
	"fmt"
	"image"
	"github.com/iamGreedy/gumi/gumre"
	"github.com/iamGreedy/gumi/drawer"
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
	VoidNode
	boundStore
	styleStore
	//
	mtColorFromTo
	studio *gumre.Studio
	onoff *gumre.Percenting
	//
	cursorEnter, active bool
	onActive            MTRadioActive
}
type MTRadioActive func(self *MTRadio, active bool)

func (s *MTRadio) GUMIInit() {
	s.studio = gumre.Animation.Studio(mtRadioAnimationLength)
	s.onoff = s.studio.Set(mtRadioAnimationOnOff, &gumre.Percenting{
		Delta: gumre.Animation.PercentingByMillis(mtRadioAnimationOnOffDeltaMillis),
		Fn:    Material.DefaultAnimation.Radio,
	}).(*gumre.Percenting)
}
func (s *MTRadio) GUMIInfomation(info Information) {
	if s.active {
		s.onoff.Request(1)
	} else {
		s.onoff.Request(0)
	}
	s.studio.Animate(float64(info.Dt))
}
func (s *MTRadio) GUMIStyle(style *Style) {
	s.style = style
}
func (s *MTRadio) GUMIClip(r image.Rectangle) {
	s.bound = r
}
func (s *MTRadio) GUMIRender(frame *image.RGBA) {
	var ctx = createContextRGBASub(frame, s.bound)
	var w, h = float64(ctx.Width()), float64(ctx.Height())
	var baseColor0 = s.GetFromMaterialColor().BaseColor()
	var mainColor1 = s.GetToMaterialColor().MainColor()
	var radius = h / 2
	var innerRadius = radius - mtRadioInnerRadiusDifference
	//
	ctx.SetColor(baseColor0)
	ctx.DrawCircle(w/2, h/2, radius)
	ctx.Fill()
	//
	ctx.SetColor(Scale.Color(baseColor0, mainColor1, s.onoff.Value()))
	ctx.DrawCircle(w/2, h/2, innerRadius)
	ctx.Fill()
}
func (s *MTRadio) GUMIDraw(frame *image.RGBA) {
	s.GUMIRender(frame)
}

func (s *MTRadio) GUMIRenderTree(tree *drawer.RenderTree, parentnode *drawer.RenderNode) {
	panic("implement me")
}
func (s *MTRadio) GUMIUpdate() {
	panic("implement me")
}

func (s *MTRadio) GUMIHappen(event Event) {
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
func (s *MTRadio) GUMISize() gumre.Size {
	return gumre.Size{
		Vertical:   gumre.FixLength(mtRadioMinHeight),
		Horizontal: gumre.FixLength(mtRadioMinWidth),
	}
}
func (s *MTRadio) String() string {
	return fmt.Sprintf("%s(active:%v)", "MTRadio", s.active)
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
