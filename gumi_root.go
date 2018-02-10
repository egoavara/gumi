package gumi

import (
	"fmt"
	"image"
	"github.com/iamGreedy/gumi/gumre"
)

type gumiRoot struct {
	SingleStructure
	scr *Screen
}

func (s *gumiRoot) String() string {
	return fmt.Sprintf("%s", "GUMI Root")
}

func (s *gumiRoot) GUMIRender(frame *image.RGBA) {
	s.child.GUMIRender(frame)
}
func (s *gumiRoot) GUMISize() gumre.Size {
	return s.child.GUMISize()
}
func (s *gumiRoot) GUMIClip(r image.Rectangle) {
	s.child.GUMIClip(r)
}
func (s *gumiRoot) Screen() *Screen {
	return s.scr
}
func (s *gumiRoot) GUMIUpdate(info *Information, style *Style) {
	s.child.GUMIUpdate(info, style)
}
func (s *gumiRoot) GUMIHappen(event Event) {
	s.child.GUMIHappen(event)
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
