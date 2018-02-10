package gumi

import (
	"image"
	"fmt"
	"github.com/iamGreedy/gumi/gumre"
)

type NDo struct {
	GUMI
	SingleStructure
	fnDraw   NDoDrawing
	fnRect   func(rect *image.Rectangle)
	fnUpdate func(info *Information, style *Style) (*Information, *Style)
	fnOccur  func(event Event) Event
}
type NDoDrawing struct{ Before, After func(frame *image.RGBA) }

func (s *NDo) String() string {
	temp := ""
	if s.fnDraw.Before != nil{
		temp += "Draw.Before, "
	}
	if s.fnDraw.After != nil{
		temp += "Draw.After, "
	}
	if s.fnOccur != nil{
		temp += "GUMIHappen, "
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

func (s *NDo) GUMIRender(frame *image.RGBA) {
	if s.fnDraw.Before != nil {
		s.fnDraw.Before(frame)
	}
	if s.fnDraw.After != nil {
		defer s.fnDraw.After(frame)
	}
	s.child.GUMIRender(frame)
}

func (s *NDo) GUMISize() gumre.Size {
	return s.child.GUMISize()
}

func (s *NDo) GUMIClip(r image.Rectangle) {
	if s.fnRect != nil {
		s.fnRect(&r)
	}
	s.child.GUMIClip(r)
}

func (s *NDo) GUMIUpdate(info *Information, style *Style) {
	if s.fnUpdate != nil {
		temp1, temp2 := s.fnUpdate(info, style)
		if temp1 != nil {
			info = temp1
		}
		if temp2 != nil {
			style = temp2
		}
	}
	s.child.GUMIUpdate(info, style)
}

func (s *NDo) GUMIHappen(event Event) {
	if s.fnOccur != nil {
		event = s.fnOccur(event)
		if event == nil {
			return
		}
	}
	s.child.GUMIHappen(event)
}

func NDo0(
	fnDraw *NDoDrawing,
	fnRect func(rect *image.Rectangle),
	fnUpdate func(info *Information, style *Style) (*Information, *Style),
	fnOccur func(event Event) Event,
) *NDo {

	if fnDraw == nil {
		fnDraw = &NDoDrawing{
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
