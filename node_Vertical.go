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
	var tempVert = make([]Length, len(s.child))
	var tempHori = make([]Length, len(s.child))

	for i, v := range s.child{
		tempVert[i] = v.size().Vertical
		tempHori[i] = v.size().Horizontal
	}
	dis := s.rule(r.Dy(), tempVert)
	//
	var startat = r.Min.Y
	dx := r.Dx()
	for i, v := range s.child{
		r := image.Rect(
			r.Min.X,
			startat,
			r.Max.X,
			startat + dis[i],
		)
		if int(tempHori[i].Max) < dx{
			r.Max.X = r.Min.X + int(tempHori[i].Max)
		}
		v.rect(r)
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
