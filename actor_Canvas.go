package gumi

import (
	"fmt"
	"github.com/iamGreedy/gumi/drawer"
	"github.com/iamGreedy/gumi/gumre"
	"image"
)

// Actor::Canvas
//
// ACanvas using for render Vector image
type ACanvas struct {
	VoidNode
	boundStore
	styleStore
	frameStore
	//
	w, h uint16
	fn   Drawer
	//
	di Information
}

// GUMIFunction / GUMIInit 					-> VoidNode::Default

// GUMIFunction / GUMIInfomation
func (s *ACanvas) GUMIInfomation(info Information) {
	s.di = info
}

// GUMIFunction / GUMIStyle					-> Define
func (s *ACanvas) GUMIStyle(style *Style) {
	s.style = style

}

// GUMIFunction / GUMIClip					-> Define
func (s *ACanvas) GUMIClip(rect image.Rectangle) {
	s.bound = rect
}

// GUMIFunction / GUMIRender				-> Define
func (s *ACanvas) GUMIRender(frame *image.RGBA) {
	ctx := createContextRGBASub(frame, s.bound)
	s.fn.Draw(ctx, s.style, s.di)
}

// GUMIFunction / GUMISize					-> Define
func (s ACanvas) GUMISize() gumre.Size {
	return gumre.Size{
		Horizontal: gumre.FixLength(uint16(s.w)),
		Vertical:   gumre.FixLength(uint16(s.h)),
	}
}

// GUMITree / born 							-> VoidNode::Default

// GUMITree / breed 						-> VoidNode::Default

// GUMITree / Parent()						-> VoidNode::Default

// GUMITree / Childrun()					-> VoidNode::Default

// GUMIRenderer / GUMIRenderSetup			-> Define
func (s *ACanvas) GUMIRenderSetup(frame *image.RGBA, tree *drawer.RenderTree, parentnode *drawer.RenderNode) {
	s.frame = frame
	// TODO : Cache
}
// GUMIRenderer / GUMIRenderSetup			-> Define
func (s *ACanvas) GUMIDraw() {
	s.GUMIRender(s.frame)
}
// GUMIRenderer / GUMIRenderSetup			-> Define
func (s *ACanvas ) GUMIUpdate()  {
	// TODO : Cache
}

// GUMIEventer / GUMIHappen					-> Define
func (s *ACanvas) GUMIHappen(event Event) {
}

// fmt.Stringer / String					-> Define
func (s *ACanvas) String() string {
	return fmt.Sprintf("%s", "ACanvas")
}


// Constructor
func ACanvas0(w, h uint16, fn Drawer) *ACanvas {
	return &ACanvas{
		w:  w,
		h:  h,
		fn: fn,
	}
}
