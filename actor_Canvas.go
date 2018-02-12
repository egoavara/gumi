package gumi

import (
	"fmt"
	"github.com/iamGreedy/gumi/gumre"
	"image"
	"github.com/iamGreedy/gumi/drawer"
)

type ACanvas struct {
	VoidStructure
	boundStore
	styleStore
	//
	w, h uint16
	fn   Drawer
	//
	di Information
}

func (s *ACanvas) GUMIInfomation(info Information) {
	s.di = info
}
func (s *ACanvas) GUMIStyle(style *Style) {
	s.style = style

}
func (s *ACanvas) GUMIClip(rect image.Rectangle) {
	s.bound = rect
}
func (s *ACanvas) GUMIRender(frame *image.RGBA) {
	ctx := createContextRGBASub(frame, s.bound)
	s.fn.Draw(ctx, s.style, s.di)
}
func (s *ACanvas) GUMIDraw(frame *image.RGBA) {
	s.GUMIRender(frame)
}
func (s ACanvas) GUMISize() gumre.Size {
	return gumre.Size{
		Horizontal: gumre.FixLength(uint16(s.w)),
		Vertical:   gumre.FixLength(uint16(s.h)),
	}
}
func (s *ACanvas) GUMIHappen(event Event) {
}
func (s *ACanvas) GUMIRenderTree(tree *drawer.RenderTree, parentnode *drawer.RenderNode) {
	// TODO : Cache
	panic("implement me")
}
func (s *ACanvas) GUMIUpdate() {
	// TODO : Cache
	panic("implement me")
}

func (s *ACanvas) String() string {
	return fmt.Sprintf("%s", "ACanvas")
}
func ACanvas0(w, h uint16, fn Drawer) *ACanvas {
	return &ACanvas{
		w:  w,
		h:  h,
		fn: fn,
	}
}
