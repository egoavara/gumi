package gumi

import "image"

type nDo struct {
	SingleStructure
	fnDraw   struct{ Before, After func(frame *image.RGBA) }
	fnRect   func(rect *image.Rectangle)
	fnUpdate func(info *Information, style *Style) (*Information, *Style)
	fnOccur  func(event Event) Event
}

func (s *nDo) draw(frame *image.RGBA) {
	if s.fnDraw.Before != nil {
		s.fnDraw.Before(frame)
	}
	if s.fnDraw.After != nil {
		defer s.fnDraw.After(frame)
	}
	s.child.draw(frame)
}

func (s *nDo) size() Size {
	return s.child.size()
}

func (s *nDo) rect(r image.Rectangle) {
	if s.fnRect != nil {
		s.fnRect(&r)
	}
	s.child.rect(r)
}

func (s *nDo) update(info *Information, style *Style) {
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

func (s *nDo) Occur(event Event) {
	if s.fnOccur != nil {
		event = s.fnOccur(event)
		if event == nil {
			return
		}
	}
	s.child.Occur(event)
}

func NDo(
	fnDraw *struct{ Before, After func(frame *image.RGBA) },
	fnRect func(rect *image.Rectangle),
	fnUpdate func(info *Information, style *Style) (*Information, *Style),
	fnOccur func(event Event) Event,
) *nDo {

	if fnDraw == nil {
		fnDraw = &struct{ Before, After func(frame *image.RGBA) }{
			After:  nil,
			Before: nil,
		}
	}
	return &nDo{
		fnRect:   fnRect,
		fnUpdate: fnUpdate,
		fnDraw:   *fnDraw,
		fnOccur:  fnOccur,
	}
}
