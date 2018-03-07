package temp

import (
	"fmt"
	"github.com/fogleman/gg"
	"image"
	"github.com/iamGreedy/gumi/gumre"
	"github.com/iamGreedy/gumi/drawer"
)

// MTProgress Default Values
const (
	mtProgressMin                        = 8
	mtProgressAnimationProgressPixelPerSecond = 512
)

// MTProgress Animations
const (
	mtProgressAnimationProgress = iota
	//
	mtProgressAnimationLength
)

// Material::Progress
//
// Material theme progress bar
type MTProgress struct {
	VoidNode
	boundStore
	styleStore
	rendererStore
	//
	mtColorFromTo
	studio   *gumre.Studio
	progress *gumre.Percenting
	//
	axis gumre.Axis
	//
	onChange MTProgressChange
	//
	cursorEnter, active bool
}

// Material::Progress<Callback> -> Change
//
// If percentage changed, it happen
type MTProgressChange func(self *MTProgress, percent float64)

// GUMIFunction / GUMIInit 					-> Define
func (s *MTProgress) GUMIInit() {
	s.studio = gumre.Animation.Studio(mtProgressAnimationLength)
	s.progress = s.studio.Set(mtProgressAnimationProgress, &gumre.Percenting{
		Fn: Material.DefaultAnimation.Progress,
	}).(*gumre.Percenting)

}

// GUMIFunction / GUMIInfomation 			-> Define
func (s *MTProgress) GUMIInfomation(info Information) {
	s.studio.Animate(float64(info.Dt))
}

// GUMIFunction / GUMIStyle 				-> Define
func (s *MTProgress) GUMIStyle(style *Style) {
	s.style = style
}

// GUMIFunction / GUMIClip 					-> Define
func (s *MTProgress) GUMIClip(r image.Rectangle) {
	s.bound = r
	switch s.axis {
	default:
		fallthrough
	case gumre.AxisHorizontal:
		s.progress.Delta = gumre.Animation.ReachingBySpeed(float64(s.bound.Dx()), mtProgressAnimationProgressPixelPerSecond)
	case gumre.AxisVertical:
		s.progress.Delta = gumre.Animation.ReachingBySpeed(float64(s.bound.Dy()), mtProgressAnimationProgressPixelPerSecond)
	}
}

// GUMIFunction / GUMIRender 				-> Define
func (s *MTProgress) GUMIRender(frame *image.RGBA) {
	var baseColor0, mainColor0 = s.GetFromMaterialColor().Color()
	var baseColor1, mainColor1 = s.GetToMaterialColor().Color()
	var ctx = createContextRGBASub(frame, s.bound)
	var w, h = float64(ctx.Width()), float64(ctx.Height())
	var percentpr = s.progress.Value()

	s.style.useContext(ctx)
	defer s.style.releaseContext(ctx)
	//
	switch s.axis {
	default:
		fallthrough
	case gumre.AxisHorizontal:
		var radius = h / 2
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
	case gumre.AxisVertical:
		var radius = w / 2
		// background
		ctx.SetColor(Scale.Color(baseColor0, baseColor1, percentpr))
		ctx.DrawArc(radius, radius, radius, gg.Radians(180), gg.Radians(360))
		ctx.DrawRectangle(0, radius, w, h-radius*2)
		ctx.DrawArc(radius, h-radius, radius, gg.Radians(0), gg.Radians(180))
		ctx.Fill()
		// progress bar
		percentLength := Scale.Length(h-radius*2, percentpr)
		ctx.SetColor(Scale.Color(mainColor0, mainColor1, percentpr))
		ctx.DrawArc(radius, h - radius - percentLength, radius, gg.Radians(180), gg.Radians(360))
		ctx.DrawRectangle(0, h - radius - percentLength, w, percentLength)
		ctx.DrawArc(radius, h - radius, radius, gg.Radians(0), gg.Radians(180))
		ctx.Fill()
	}
}

// GUMIFunction / GUMISize 					-> Define
func (s *MTProgress) GUMISize() gumre.Size {
	return gumre.Size{
		gumre.MinLength(mtProgressMin),
		gumre.MinLength(mtProgressMin),
	}
}

// GUMITree / born 							-> VoidNode::Default

// GUMITree / breed 						-> VoidNode::Default

// GUMITree / parent()						-> VoidNode::Default

// GUMITree / childrun()					-> VoidNode::Default

// GUMIRenderer / GUMIRenderSetup 			-> Define
func (s *MTProgress) GUMIRenderSetup(frame *image.RGBA, tree *drawer.RenderTree, parentnode *drawer.RenderNode) {
	s.frame = frame
}

// GUMIRenderer / GUMIUpdate 				-> Define
func (s *MTProgress) GUMIUpdate() {
	panic("implement me")
}

// GUMIRenderer / GUMIDraw 				-> Define
func (s *MTProgress) GUMIDraw() {
	s.GUMIRender(s.frame)
}

// GUMIEventer / GUMIHappen					-> Define
func (s *MTProgress) GUMIHappen(event Event) {

}

// fmt.Stringer / String					-> Define
func (s *MTProgress) String() string {
	return fmt.Sprintf("%s(axis: %v, percent:%.2f%%)", "MTProgress", s.axis, s.GetPercent()*100)
}

// Constructor 0
func MTProgress0(mcl *MaterialColor) *MTProgress {
	temp := &MTProgress{
		axis:gumre.AxisHorizontal,
	}
	temp.SetFromMaterialColor(mcl)
	temp.SetToMaterialColor(mcl)
	return temp
}

// Constructor 1
func MTProgress1(from, to *MaterialColor) *MTProgress {
	temp := &MTProgress{
		axis:gumre.AxisHorizontal,
	}
	temp.SetFromMaterialColor(from)
	temp.SetToMaterialColor(to)
	return temp
}

// Constructor 2
func MTProgress2(from, to *MaterialColor, axis gumre.Axis) *MTProgress {
	temp := &MTProgress{
		axis:axis,
	}
	temp.SetFromMaterialColor(from)
	temp.SetToMaterialColor(to)
	return temp
}


// Method / Get -> GetPercent
func (s *MTProgress) Get() float64 {
	return s.GetPercent()
}

// Method / Set -> SetPercent
func (s *MTProgress) Set(percent float64) {
	s.SetPercent(percent)
}

// Method / Get
func (s *MTProgress) GetPercent() float64 {
	return s.progress.To
}

// Method / Set
func (s *MTProgress) SetPercent(percent float64) {
	s.progress.Request(percent)
	if s.onChange != nil {
		s.onChange(s, percent)
	}
}

// Method / Get
func (s *MTProgress) GetAxis() gumre.Axis {
	return s.axis
}

// Method / Set
func (s *MTProgress) SetAxis(axis gumre.Axis) {
	s.axis = axis
}

// Method / Set Callback
func (s *MTProgress) OnChange(callback MTProgressChange) {
	s.onChange = callback
}

// Method / Get Callback
func (s *MTProgress) ReferChange() MTProgressChange {
	return s.onChange
}
