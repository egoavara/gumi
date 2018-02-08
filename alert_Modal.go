package gumi

import (
	"image"
	"fmt"
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

func (s *ALModal) init() {
	s.modal.init()
	s.child.init()
}
func (s *ALModal) draw(frame *image.RGBA) {
	s.child.draw(frame)
	if s.show{
		s.modal.draw(frame)
	}
}
func (s *ALModal) size() Size {
	return s.child.size()
}
func (s *ALModal) rect(r image.Rectangle) {
	s.bound = r
	if s.show{
		s.modal.rect(r)
	}else {
		s.child.rect(r)
	}
}
func (s *ALModal) update(info *Information, style *Style) {
	s.style = style
	s.child.update(info, style)
	s.modal.update(info, style)
}
func (s *ALModal) Occur(event Event) {
	if s.show{
		s.modal.Occur(event)
	}else {
		if event.Kind() == EVENT_CURSOR {
			s.lastCursorEvent = event.(EventCursor)
		}
		s.child.Occur(event)
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
	temp.modal.Born(temp)
	return temp
}

//
func (s *ALModal ) SetModal(modal GUMI)  {
	s.modal = modal
	s.modal.Born(s)
}
func (s *ALModal ) GetModal() GUMI {
	return s.modal
}
func (s *ALModal ) SetShow(show bool)  {
	s.show = show
	s.modal.Occur(s.lastCursorEvent)
}
func (s *ALModal ) GetShow() bool {
	return s.show
}