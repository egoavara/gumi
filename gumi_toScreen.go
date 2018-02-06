package gumi

import (
	"fmt"
	"image"
)

type gumiRoot struct {
	SingleStructure
	scr *Screen
}

func (s *gumiRoot) String() string {
	return fmt.Sprintf("%s", "GUMI Root")
}

func (s *gumiRoot) draw(frame *image.RGBA) {
	s.child.draw(frame)
}
func (s *gumiRoot) size() Size {
	return s.child.size()
}
func (s *gumiRoot) rect(r image.Rectangle) {
	s.child.rect(r)
}
func (s *gumiRoot) update(info *Information, style *Style) {
	s.child.update(info, style)
}
func (s *gumiRoot) Occur(event Event) {
	s.child.Occur(event)
}
func newGUMIRoot(scr *Screen, under GUMI) GUMI {
	temp := &gumiRoot{
		scr: scr,
	}

	return LinkingFrom(temp, under)
}

func getScreen(g GUMI) *Screen {
	if g == nil{
		return nil
	}
	if v, ok := g.(*gumiRoot); ok{
		return v.scr
	}
	return getScreen(g.Parent())
}
