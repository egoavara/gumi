package gumi

import (
	"fmt"
	"image"
	"github.com/iamGreedy/gumi/gumre"
	"github.com/iamGreedy/gumi/drawer"
)

type gumiRoot struct {
	SingleStructure
	scr *Screen
}

func (s *gumiRoot) GUMIInfomation(info Information) {
	s.child.GUMIInfomation(info)
}
func (s *gumiRoot) GUMIStyle(style *Style) {
	s.child.GUMIStyle(style)
}
func (s *gumiRoot) GUMIClip(r image.Rectangle) {
	s.child.GUMIClip(r)
}
func (s *gumiRoot) GUMIRender(frame *image.RGBA) {

}
func (s *gumiRoot) GUMIDraw(frame *image.RGBA) {
	s.child.GUMIDraw(frame)
}
func (s *gumiRoot) GUMIRenderTree(tree *drawer.RenderTree, parentnode *drawer.RenderNode) {
	panic("implement me")
}
func (s *gumiRoot) GUMIUpdate() {
	panic("implement me")
}
func (s *gumiRoot) GUMIHappen(event Event) {
	s.child.GUMIHappen(event)
}
func (s *gumiRoot) GUMISize() gumre.Size {
	return s.child.GUMISize()
}
func (s *gumiRoot) String() string {
	return fmt.Sprintf("%s", "GUMI Root")
}
// GUMIRoot interface
func (s *gumiRoot) Screen() *Screen {
	return s.scr
}
func newGUMIRoot(scr *Screen, under GUMI) GUMIRoot {
	temp := &gumiRoot{
		scr: scr,
	}
	LinkingFrom(temp, under)
	return temp
}
func Root(g GUMI) GUMIRoot {
	if g == nil{
		return nil
	}
	if v, ok := g.(*gumiRoot); ok{
		return v
	}
	return Root(g.Parent())
}
