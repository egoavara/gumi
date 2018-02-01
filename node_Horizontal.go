package gumi

import (
	"image"
	"math"
	"fmt"
)

type NHorizontal struct {
	MultipleStructure
	rule Distribute
}

func (s *NHorizontal) String() string {
	return fmt.Sprintf("%s(childrun:%d)", "NHorizontal", len(s.Childrun()))
}
func (s *NHorizontal) draw(frame *image.RGBA) {
	for _, v := range s.child{
		v.draw(frame)
	}
}
func (s *NHorizontal) size() Size {
	var min, max, sum uint16 = 0, math.MaxUint16, 0
	for _, v := range s.child{
		sz := v.size()
		if sz.Vertical.Min > min{
			min = sz.Vertical.Min
		}
		if sz.Vertical.Max < max{
			max = sz.Vertical.Max
		}
		sum += sz.Horizontal.Min
	}
	return Size{
		Length{Min:min, Max:max},
		MinLength(sum),
	}
}
func (s *NHorizontal) rect(r image.Rectangle) {
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
func (s *NHorizontal) update(info *Information, style *Style) {
	for _, v := range s.child{
		v.update(info, style)
	}
}
func (s *NHorizontal) Occur(event Event) {
	for _, v := range s.child{
		v.Occur(event)
	}
}
func NHorizontal0(rule Distribute, childrun ...GUMI) *NHorizontal {
	s := &NHorizontal{
		rule:rule,
	}
	s.Breed(childrun)
	return s
}
func NHorizontal1(childrun ...GUMI) *NHorizontal {
	s := &NHorizontal{
		rule: Distribution.Minimalize,
	}
	s.Breed(childrun)
	return s
}
