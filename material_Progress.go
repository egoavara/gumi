package gumi

import (
	"fmt"
	"github.com/fogleman/gg"
	"image"
)

const (
	mtProgressMin                        = 8
	mtProgressAnimationProgressPixelPerSecond = 512
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
	studio   *AnimationStudio
	progress *AnimationPercent
	//
	axis Axis
	//
	onChange MTProgressChange
	//
	cursorEnter, active bool
}
type MTProgressChange func(self *MTProgress, percent float64)

func (s *MTProgress) String() string {
	return fmt.Sprintf("%s(axis: %v, percent:%.2f%%)", "MTProgress", s.axis, s.GetPercent()*100)
}
func (s *MTProgress) init() {
	s.studio = NewAnimationStudio(mtProgressAnimationLength)
	s.progress = s.studio.Set(mtProgressAnimationProgress, &AnimationPercent{
		Fn: Material.DefaultAnimation.Progress,
	}).(*AnimationPercent)

}
func (s *MTProgress) draw(frame *image.RGBA) {
	var baseColor0, mainColor0 = s.GetFromMaterialColor().Color()
	var baseColor1, mainColor1 = s.GetToMaterialColor().Color()
	var ctx = GGContextRGBASub(frame, s.bound)
	var w, h = float64(ctx.Width()), float64(ctx.Height())
	var percentpr = s.progress.Value()
	var radius = h / 2
	s.style.useContext(ctx)
	defer s.style.releaseContext(ctx)
	//
	switch s.axis {
	default:
		fallthrough
	case AxisHorizontal:
		// background
		ctx.SetColor(Scale.Color(baseColor0, baseColor1, percentpr))
		ctx.DrawArc(radius, radius, radius, gg.Radians(90), gg.Radians(270))
		ctx.DrawRectangle(radius, 0, w-radius*2, h)
		ctx.DrawArc(w-radius, radius, radius, gg.Radians(-90), gg.Radians(90))
		ctx.Fill()
		// progress bar
		percentLength := Scale.Length(w-radius*2, percentpr)
		ctx.SetColor(Scale.Color(mainColor0, mainColor1, percentpr))
		ctx.DrawArc(radius, radius, radius, gg.Radians(90), gg.Radians(270))
		ctx.DrawRectangle(radius, 0, percentLength, h)
		ctx.DrawArc(radius+percentLength, radius, radius, gg.Radians(-90), gg.Radians(90))
		ctx.Fill()
	case AxisVertical:
		// background
		ctx.SetColor(Scale.Color(baseColor0, baseColor1, percentpr))
		ctx.DrawArc(radius, radius, radius, gg.Radians(180), gg.Radians(360))
		ctx.DrawRectangle(0, radius, w, h-radius*2)
		ctx.DrawArc(radius, h-radius, radius, gg.Radians(0), gg.Radians(180))
		ctx.Fill()
		// progress bar
		percentLength := Scale.Length(h-radius*2, percentpr)
		ctx.SetColor(Scale.Color(mainColor0, mainColor1, percentpr))
		ctx.DrawArc(radius, radius, radius, gg.Radians(180), gg.Radians(360))
		ctx.DrawRectangle(0, radius, w, percentLength)
		ctx.DrawArc(radius, radius + percentLength, radius, gg.Radians(0), gg.Radians(180))
		ctx.Fill()
	}
}
func (s *MTProgress) size() Size {
	return Size{
		MinLength(mtProgressMin),
		MinLength(mtProgressMin),
	}
}
func (s *MTProgress) rect(r image.Rectangle) {
	s.bound = r
	switch s.axis {
	default:
		fallthrough
	case AxisHorizontal:
		s.progress.Delta = Animation.DeltaBySpeed(float64(s.bound.Dx()), mtProgressAnimationProgressPixelPerSecond)
	case AxisVertical:
		s.progress.Delta = Animation.DeltaBySpeed(float64(s.bound.Dy()), mtProgressAnimationProgressPixelPerSecond)
	}
}
func (s *MTProgress) update(info *Information, style *Style) {
	s.style = style
	s.studio.Animate(info)
}
func (s *MTProgress) Occur(event Event) {

}

//
func MTProgress0(mcl *MaterialColor) *MTProgress {
	temp := &MTProgress{
		axis:AxisHorizontal,
	}
	temp.SetFromMaterialColor(mcl)
	temp.SetToMaterialColor(mcl)
	return temp
}
func MTProgress1(from, to *MaterialColor) *MTProgress {
	temp := &MTProgress{
		axis:AxisHorizontal,
	}
	temp.SetFromMaterialColor(from)
	temp.SetToMaterialColor(to)
	return temp
}
func MTProgress2(from, to *MaterialColor, axis Axis) *MTProgress {
	temp := &MTProgress{
		axis:axis,
	}
	temp.SetFromMaterialColor(from)
	temp.SetToMaterialColor(to)
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
	return s.progress.To
}
func (s *MTProgress) SetPercent(percent float64) {
	s.progress.Request(percent)
	if s.onChange != nil {
		s.onChange(s, percent)
	}
}
func (s *MTProgress) GetAxis() Axis {
	return s.axis
}
func (s *MTProgress) SetAxis(axis Axis) {
	s.axis = axis
}
func (s *MTProgress) OnChange(callback MTProgressChange) {
	s.onChange = callback
}
func (s *MTProgress) ReferChange() MTProgressChange {
	return s.onChange
}
