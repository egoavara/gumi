package gumi

import (
	"image"
	"math"
)

type nHorizontal struct {
	MultipleStructure
	rule Dispensor
}

func (s *nHorizontal) draw(frame *image.RGBA) {
	for _, v := range s.child{
		v.draw(frame)
	}
}

func (s *nHorizontal) size() Size {
	var min, max uint16 = 0, math.MaxUint16
	for _, v := range s.child{
		sz := v.size()
		if sz.Vertical.Min > min{
			min = sz.Vertical.Min
		}
		if sz.Vertical.Max < max{
			max = sz.Vertical.Max
		}
	}
	return Size{
		Length{Min:min, Max:max},
		MAXLENGTH,
	}
}


func (s *nHorizontal) rect(r image.Rectangle) {
	//
	var tempVert = make([]Length, len(s.child))
	var tempHori = make([]Length, len(s.child))

	for i, v := range s.child{
		tempVert[i] = v.size().Vertical
		tempHori[i] = v.size().Horizontal
	}
	dis := s.rule(r.Dx(), tempHori)
	//
	var startat = r.Min.X
	dy := r.Dy()
	for i, v := range s.child{
		r := image.Rect(
			startat,
			r.Min.Y,
			startat + dis[i],
			r.Max.Y,
		)
		if int(tempVert[i].Max) < dy{
			r.Max.Y = r.Min.Y + int(tempVert[i].Max)
		}
		v.rect(r)
		startat += dis[i]
	}
}

func (s *nHorizontal) update(info *Information, style *Style) {
	for _, v := range s.child{
		v.update(info, style)
	}
}

func (s *nHorizontal) Occur(event Event) {
	for _, v := range s.child{
		v.Occur(event)
	}
}
func NHorizontal(rule Dispensor, childrun ...GUMI) *nHorizontal {
	s := &nHorizontal{
		rule:rule,
	}
	s.Breed(childrun)
	return s
}
func NHorizontal1(childrun ...GUMI) *nHorizontal{
	s := &nHorizontal{
		rule: DispensorMinimalize,
	}
	s.Breed(childrun)
	return s
}
