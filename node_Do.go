package gumi

import (
	"image"
	"fmt"
)

type NDo struct {
	SingleStructure
	fnDraw   struct{ Before, After func(frame *image.RGBA) }
	fnRect   func(rect *image.Rectangle)
	fnUpdate func(info *Information, style *Style) (*Information, *Style)
	fnOccur  func(event Event) Event
}

func (s *NDo) String() string {
	temp := ""
	if s.fnDraw.Before != nil{
		temp += "Draw.Before, "
	}
	if s.fnDraw.After != nil{
		temp += "Draw.After, "
	}
	if s.fnOccur != nil{
		temp += "Occur, "
	}
	if s.fnRect != nil{
		temp += "Rect, "
	}
	if s.fnUpdate != nil{
		temp += "Update, "
	}
	if len(temp) == 0{
		return fmt.Sprintf("%s", "NDo")
	}
	temp = temp[:len(temp) - 2]
	return fmt.Sprintf("%s(%s)", "NDo", temp)
}

func (s *NDo) draw(frame *image.RGBA) {
	if s.fnDraw.Before != nil {
		s.fnDraw.Before(frame)
	}
	if s.fnDraw.After != nil {
		defer s.fnDraw.After(frame)
	}
	s.child.draw(frame)
}

func (s *NDo) size() Size {
	return s.child.size()
}

func (s *NDo) rect(r image.Rectangle) {
	if s.fnRect != nil {
		s.fnRect(&r)
	}
	s.child.rect(r)
}

func (s *NDo) update(info *Information, style *Style) {
	if s.fnUpdate != nil {
		temp1, temp2 := s.fnUpdate(info, style)
		if temp1 != nil {
			info = temp1
		}
		if temp2 != nil {
			style = temp2
		}
	}
	s.child.update(info, style)
}

func (s *NDo) Occur(event Event) {
	if s.fnOccur != nil {
		event = s.fnOccur(event)
		if event == nil {
			return
		}
	}
	s.child.Occur(event)
}

func NDo0(
	fnDraw *struct{ Before, After func(frame *image.RGBA) },
	fnRect func(rect *image.Rectangle),
	fnUpdate func(info *Information, style *Style) (*Information, *Style),
	fnOccur func(event Event) Event,
) *NDo {

	if fnDraw == nil {
		fnDraw = &struct{ Before, After func(frame *image.RGBA) }{
			After:  nil,
			Before: nil,
		}
	}
	return &NDo{
		fnRect:   fnRect,
		fnUpdate: fnUpdate,
		fnDraw:   *fnDraw,
		fnOccur:  fnOccur,
	}
}
