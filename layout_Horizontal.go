package gumi

import (
	"image"
	"math"
	"fmt"
	"sync"
)

type LHorizontal struct {
	MultipleStructure
	rule Distribute
}

func (s *LHorizontal) String() string {
	return fmt.Sprintf("%s(childrun:%d)", "LHorizontal", len(s.Childrun()))
}
func (s *LHorizontal) draw(frame *image.RGBA) {
	wg := new(sync.WaitGroup)
	wg.Add(len(s.child))
	defer wg.Wait()
	for _, v := range s.child{
		go func(elem GUMI) {
			elem.draw(frame)
			wg.Done()
		}(v)
	}
}
func (s *LHorizontal) size() Size {
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
func (s *LHorizontal) rect(r image.Rectangle) {
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
func (s *LHorizontal) update(info *Information, style *Style) {
	for _, v := range s.child{
		v.update(info, style)
	}
}
func (s *LHorizontal) Occur(event Event) {
	for _, v := range s.child{
		v.Occur(event)
	}
}
func LHorizontal0(rule Distribute, childrun ...GUMI) *LHorizontal {
	s := &LHorizontal{
		rule:rule,
	}
	s.Breed(childrun)
	return s
}
func LHorizontal1(childrun ...GUMI) *LHorizontal {
	s := &LHorizontal{
		rule: Distribution.Minimalize,
	}
	s.Breed(childrun)
	return s
}
