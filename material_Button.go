package gumi

import (
	"fmt"
	"github.com/fogleman/gg"
	"image"
	"github.com/iamGreedy/gumi/gumre"
	"github.com/iamGreedy/gumi/drawer"
)

// MTButton Default Values
const (
	mtButtonMinPadding = 5
	mtButtonAnimDeltaMillis = 250
)

// MTButton Animations
const (
	mtButtonAnimationHover = iota
	//
	mtButtonAnimationLength = iota
)

// Material::Button
//
// Material theme button
type MTButton struct {
	//
	VoidNode
	styleStore
	rendererStore
	//
	mtColorSingle
	studio *gumre.Studio
	hover *gumre.Percenting
	//
	text string
	//
	cursorEnter, active bool
	onClick             MTButtonClick
	onFocus             MTButtonFocus
}



// Material::Button<Callback> -> Focus
//
// When Cursor enter, leave this Elem call this function
type MTButtonFocus func(self *MTButton, focus bool)

// Material::Button<Callback> -> Click
//
// When Cursor click occur this
type MTButtonClick func(self *MTButton)


// GUMIFunction / GUMIInit 					-> Define
func (s *MTButton) GUMIInit() {
	s.studio = gumre.Animation.Studio(mtButtonAnimationLength)
	s.hover = s.studio.Set(mtButtonAnimationHover, &gumre.Percenting{
		Delta:gumre.Animation.PercentingByMillis(mtButtonAnimDeltaMillis),
		Fn: Material.DefaultAnimation.Button,
	}).(*gumre.Percenting)
}

// GUMIFunction / GUMIInfomation 			-> Define
func (s *MTButton) GUMIInfomation(info Information) {
	if s.studio.Animate(float64(info.Dt)){
		s.rnode.Require()
	}
}

// GUMIFunction / GUMIStyle 				-> Define
func (s *MTButton) GUMIStyle(style *Style) {
	s.style = style
}

// GUMIFunction / GUMIClip 					-> Define
func (s *MTButton) GUMIClip(r image.Rectangle) {
	s.rnode.SetRect(r)
}
// GUMIFunction / GUMIRender 				-> Define
func (s *MTButton) GUMIRender(frame *image.RGBA) {
	var ctx = createContext(frame)
	var w, h = float64(ctx.Width()), float64(ctx.Height())
	var baseColor, mainColor = s.GetMaterialColor().Color()
	radius := h / 2
	s.style.useContext(ctx)
	defer s.style.releaseContext(ctx)
	//

	if s.active {
		ctx.SetColor(baseColor)
		// background
		ctx.DrawArc(radius, radius, radius, gg.Radians(90), gg.Radians(270))
		ctx.DrawRectangle(radius, 0, w-radius*2, h)
		ctx.DrawArc(w-radius, radius, radius, gg.Radians(-90), gg.Radians(90))
		ctx.Fill()
		// text
		ctx.SetColor(mainColor)
		txtw, txth := ctx.MeasureString(s.text)
		ctx.DrawString(s.text, (w-txtw)/2, (h+txth)/2-1)
		ctx.Stroke()
	} else {

		ctx.SetColor(Scale.Color(baseColor, mainColor, s.hover.Value()))
		// background
		ctx.DrawArc(radius, radius, radius, gg.Radians(90), gg.Radians(270))
		ctx.DrawRectangle(radius, 0, w-radius*2, h)
		ctx.DrawArc(w-radius, radius, radius, gg.Radians(-90), gg.Radians(90))
		ctx.Fill()
		// text
		ctx.SetColor(Scale.Color(mainColor, baseColor, s.hover.Value()))
		txtw, txth := ctx.MeasureString(s.text)
		ctx.DrawString(s.text, (w-txtw)/2, (h+txth)/2-1)
		ctx.Stroke()
	}
}

// GUMIFunction / GUMISize 					-> Define
func (s *MTButton) GUMISize() gumre.Size {

	s.style.Default.Font.Use()
	defer s.style.Default.Font.Release()
	var hori, vert = s.style.Default.Font.CalculateSize(s.text)
	//
	return gumre.Size{
		Vertical:   gumre.MinLength(uint16(vert + mtButtonMinPadding*2)),
		Horizontal: gumre.MinLength(uint16(hori + mtButtonMinPadding*2)),
	}
}

// GUMITree / born 							-> VoidNode::Default

// GUMITree / breed 						-> VoidNode::Default

// GUMITree / parent()						-> VoidNode::Default

// GUMITree / childrun()					-> VoidNode::Default

// GUMIRenderer / GUMIRenderSetup 			-> Define
func (s *MTButton) GUMIRenderSetup(tree *drawer.RenderTree, parentnode *drawer.RenderNode) {
	s.rtree = tree
	s.rnode = tree.New(parentnode)
}

// GUMIRenderer / GUMIUpdate 				-> Define
func (s *MTButton) GUMIUpdate() {
	if s.rnode.Check(){
		s.GUMIRender(s.rnode.SubImage())
		s.rnode.Complete()
	}

}

// GUMIEventer / GUMIHappen					-> Define
func (s *MTButton) GUMIHappen(event Event) {
	switch ev := event.(type) {
	case EventKeyPress:
		if ev.Key == KEY_MOUSE1 {
			if s.cursorEnter {
				s.active = true
				s.rnode.Require()
			}
		}
	case EventKeyRelease:
		if ev.Key == KEY_MOUSE1 {
			if s.active {
				if s.onClick != nil {
					s.onClick(s)
				}
				s.active = false
				s.rnode.Require()
			}
		}
	case EventCursor:
		x := int(ev.X)
		y := int(ev.Y)
		bd := s.rnode.GetRect()
		if (bd.Min.X <= x && x < bd.Max.X) && (bd.Min.Y <= y && y < bd.Max.Y) {
			if s.onFocus != nil {
				s.onFocus(s, true)
			}
			//
			s.hover.Delta = gumre.Animation.PercentingByMillis(mtButtonAnimDeltaMillis)
			s.hover.Request(1)
			s.cursorEnter = true
		} else {
			if s.onFocus != nil {
				s.onFocus(s, false)
			}
			s.hover.Delta = gumre.Animation.PercentingByMillis(mtButtonAnimDeltaMillis)
			s.hover.Request(0)
			s.cursorEnter = false
		}
	}
}

// fmt.Stringer / String					-> Define
func (s *MTButton) String() string {
	return fmt.Sprintf("%s", "MTButton")
}

// Constructor 0
func MTButton0(text string, onclick MTButtonClick) *MTButton {
	temp := &MTButton{
		text:    text,
		onClick: onclick,
	}
	temp.SetMaterialColor(Material.Pallette.White)
	return temp
}

// Constructor 1
func MTButton1(mcl *MaterialColor, text string, onclick MTButtonClick) *MTButton {
	temp := &MTButton{
		text:    text,
		onClick: onclick,
	}
	temp.SetMaterialColor(mcl)
	return temp
}

// Method / Set -> SetText
func (s *MTButton) Set(txt string) {
	s.SetText(txt)
}

// Method / Get -> GetText
func (s *MTButton) Get() string {
	return s.GetText()
}

// Method / Set
func (s *MTButton) SetText(txt string) {
	s.text = txt
}

// Method / Get
func (s *MTButton) GetText() string {
	return s.text
}

// Method / Set Callback
func (s *MTButton) OnClick(callback MTButtonClick) {
	s.onClick = callback
}

// Method / Get Callback
func (s *MTButton) ReferClick() MTButtonClick {
	return s.onClick
}

// Method / Set Callback
func (s *MTButton) OnFocus(callback MTButtonClick) {
	s.onClick = callback
}

// Method / Get Callback
func (s *MTButton) ReferFocus() MTButtonClick {
	return s.onClick
}
