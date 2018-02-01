package gumi

import (
	"image"
	"math"
	"fmt"
)

type NVertical struct {
	MultipleStructure
	rule Distribute
}

func (s *NVertical) draw(frame *image.RGBA) {
	for _, v := range s.child{
		v.draw(frame)
	}
}
func (s *NVertical) size() Size {
	var min, max, sum uint16 = 0, math.MaxUint16, 0

	for _, v := range s.child{
		sz := v.size()
		if sz.Horizontal.Min > min{
			min = sz.Horizontal.Min
		}
		if sz.Horizontal.Max < max{
			max = sz.Horizontal.Max
		}
		sum += sz.Vertical.Min
	}
	//
	return Size{
		MinLength(sum),
		Length{Min: min, Max:max},
	}
}
func (s *NVertical) rect(r image.Rectangle) {
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
func (s *NVertical) update(info *Information, style *Style) {
	for _, v := range s.child{
		v.update(info, style)
	}
}
func (s *NVertical) Occur(event Event) {
	for _, v := range s.child{
		v.Occur(event)
	}
}
func (s *NVertical) String() string{
	return fmt.Sprintf(
		"%s(childrun:%d)", "NVertical", len(s.Childrun()),
	)
}

func NVertical0(rule Distribute, childrun ...GUMI) *NVertical {
	s := &NVertical{
		rule:rule,
	}
	for _, v := range childrun{
		v.Born(s)
	}
	s.Breed(childrun)
	return s
}
func NVertical1(childrun ...GUMI) *NVertical {
	s := &NVertical{
		rule: Distribution.Minimalize,
	}
	for _, v := range childrun{
		v.Born(s)
	}
	s.Breed(childrun)
	return s
}
