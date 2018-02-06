package gumi

import (
	"fmt"
	"github.com/fogleman/gg"
	"image"
)

const (
	mtProgressMinWidth        = 40
	mtProgressMinHeight       = 8
	mtProgressZeroToOneMillis = 600
)
const (
	mtProgressAnimationProgress = iota
	//
	mtProgressAnimationLength
)

type MTProgress struct {
	VoidStructure
	boundStore
	styleStore
	//
	mtColorFromTo
	studio *AnimationStudio
	//
	onChange MTProgressChange
	//
	cursorEnter, active bool
}
type MTProgressChange func(self *MTProgress, percent float64)

func (s *MTProgress) String() string {
	return fmt.Sprintf("%s(percent:%.2f%%)", "MTProgress", s.GetPercent()*100)
}
func (s *MTProgress) init() {
	s.studio = NewAnimationStudio(mtProgressAnimationLength)
	s.studio.Set(mtProgressAnimationProgress, &AnimationPercent{
		Delta:Animation.DeltaByMillis(mtProgressZeroToOneMillis),
		Fn: Material.DefaultAnimation.Progress,
	})

}
func (s *MTProgress) draw(frame *image.RGBA) {
	var baseColor0, mainColor0 = s.GetFromMaterialColor().Color()
	var baseColor1, mainColor1 = s.GetToMaterialColor().Color()
	var ctx = GGContextRGBASub(frame, s.bound)
	var w, h = float64(ctx.Width()), float64(ctx.Height())
	var percent = s.studio.Get(mtProgressAnimationProgress).(*AnimationPercent)
	radius := h / 2
	s.style.useContext(ctx)
	defer s.style.releaseContext(ctx)
	// background
	ctx.SetColor(Scale.Color(baseColor0, baseColor1, percent.Current))
	ctx.DrawArc(radius, radius, radius, gg.Radians(90), gg.Radians(270))
	ctx.DrawRectangle(radius, 0, w-radius*2, h)
	ctx.DrawArc(w-radius, radius, radius, gg.Radians(-90), gg.Radians(90))
	ctx.Fill()
	// progress bar
	percentLength := Scale.Length(w-radius*2, percent.Current)
	ctx.SetColor(Scale.Color(mainColor0, mainColor1, percent.Current))
	ctx.DrawArc(radius, radius, radius, gg.Radians(90), gg.Radians(270))
	ctx.DrawRectangle(radius, 0, percentLength, h)
	ctx.DrawArc(radius+percentLength, radius, radius, gg.Radians(-90), gg.Radians(90))
	ctx.Fill()
}
func (s *MTProgress) size() Size {
	return Size{
		MinLength(mtProgressMinHeight),
		MinLength(mtProgressMinWidth),
	}
}
func (s *MTProgress) rect(r image.Rectangle) {
	s.bound = r
}
func (s *MTProgress) update(info *Information, style *Style) {
	s.style = style
	s.studio.Animate(info)
}
func (s *MTProgress) Occur(event Event) {

}
//
func MTProgress0(from, to *MaterialColor) *MTProgress {
	temp := &MTProgress{}
	temp.SetFromMaterialColor(from)
	temp.SetToMaterialColor(to)
	return temp
}
func MTProgress1(mcl *MaterialColor) *MTProgress {
	temp := &MTProgress{}
	temp.SetFromMaterialColor(mcl)
	temp.SetToMaterialColor(mcl)
	return temp
}
//
func (s *MTProgress) Get() float64 {
	return s.GetPercent()
}
func (s *MTProgress) Set(percent float64) {
	s.SetPercent(percent)
}
func (s *MTProgress) GetPercent() float64 {
	return s.studio.Get(mtProgressAnimationProgress).(*AnimationPercent).To
}
func (s *MTProgress) SetPercent(percent float64) {
	s.studio.Get(mtProgressAnimationProgress).(*AnimationPercent).Request(percent)
	if s.onChange != nil {
		s.onChange(s, percent)
	}
}
func (s *MTProgress) OnChange(callback MTProgressChange) {
	s.onChange = callback
}
func (s *MTProgress) ReferChange() MTProgressChange {
	return s.onChange
}
