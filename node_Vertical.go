package gumi

import (
	"image"
)

type nVertical struct {
	MultipleStructure
	rule Dispensor
}

func (s *nVertical) draw(frame *image.RGBA) {
	for _, v := range s.child{
		v.draw(frame)
	}
}

func (s *nVertical) size() Size {
	return Size{
		AUTOLENGTH,
		AUTOLENGTH,
	}
}


func (s *nVertical) rect(r image.Rectangle) {
	//
	var temp = make([]Length, len(s.child))
	for i, v := range s.child{
		temp[i] = v.size().Vertical
	}
	dis := s.rule(r.Dy(), temp)
	//
	var startat = r.Min.Y
	for i, v := range s.child{
		v.rect(image.Rect(
			r.Min.X,
			startat,
			r.Max.X,
			startat + dis[i],
		))
		startat += dis[i]
	}
}

func (s *nVertical) update(info *Information, style *Style) {
	for _, v := range s.child{
		v.update(info, style)
	}
}

func (s *nVertical) Occur(event Event) {
	for _, v := range s.child{
		v.Occur(event)
	}
}
func NVertical(rule Dispensor, childrun ...GUMI) *nVertical {
	s := &nVertical{
		rule:rule,
	}
	s.Breed(childrun)
	return s
}
func NVertical1(childrun ...GUMI) *nVertical {
	s := &nVertical{
		rule: DispensorMinimalize,
	}
	s.Breed(childrun)
	return s
}
