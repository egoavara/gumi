package gumi

import (
	"image"
	"fmt"
	"github.com/iamGreedy/gumi/gumre"
)

type ALModal struct {
	SingleStructure
	boundStore
	styleStore
	//
	lastCursorEvent EventCursor
	//
	modal GUMI
	show bool
}

func (s *ALModal) GUMIInit() {
	s.modal.GUMIInit()
	s.child.GUMIInit()
}
func (s *ALModal) GUMIRender(frame *image.RGBA) {
	s.child.GUMIRender(frame)
	if s.show{
		s.modal.GUMIRender(frame)
	}
}
func (s *ALModal) GUMISize() gumre.Size {
	return s.child.GUMISize()
}
func (s *ALModal) GUMIClip(r image.Rectangle) {
	s.bound = r
	if s.show{
		s.modal.GUMIClip(r)
	}else {
		s.child.GUMIClip(r)
	}
}
func (s *ALModal) GUMIUpdate(info *Information, style *Style) {
	s.style = style
	s.child.GUMIUpdate(info, style)
	s.modal.GUMIUpdate(info, style)
}
func (s *ALModal) GUMIHappen(event Event) {
	if s.show{
		s.modal.GUMIHappen(event)
	}else {
		if event.Kind() == EVENT_CURSOR {
			s.lastCursorEvent = event.(EventCursor)
		}
		s.child.GUMIHappen(event)
	}
}
func (s *ALModal) String() string {
	return fmt.Sprintf("%s", "ALModal")
}
//
func ALModal0() *ALModal {
	temp := &ALModal{}
	return temp
}
func ALModal1(modal GUMI) *ALModal {
	temp := &ALModal{
		modal:modal,
	}
	temp.modal.born(temp)
	return temp
}

//
func (s *ALModal ) SetModal(modal GUMI)  {
	s.modal = modal
	s.modal.born(s)
}
func (s *ALModal ) GetModal() GUMI {
	return s.modal
}
func (s *ALModal ) SetShow(show bool)  {
	s.show = show
	s.modal.GUMIHappen(s.lastCursorEvent)
}
func (s *ALModal ) GetShow() bool {
	return s.show
}