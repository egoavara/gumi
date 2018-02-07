package gumi

import (
	"image"
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
	var minMax, sum uint16 = 0, 0
	for _, v := range s.child{
		sz := v.size()
		if sz.Vertical.Min > minMax{
			minMax = sz.Vertical.Min
		}
		sum += sz.Horizontal.Min
	}
	return Size{
		MinLength(minMax),
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
	for i, v := range s.child{
		r := image.Rect(
			startat,
			r.Min.Y,
			startat + dis[i],
			r.Max.Y,
		)
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
	for _, v := range childrun{
		v.Born(s)
	}
	s.Breed(childrun)
	return s
}
func LHorizontal1(childrun ...GUMI) *LHorizontal {
	s := &LHorizontal{
		rule: Distribution.Minimalize,
	}
	for _, v := range childrun{
		v.Born(s)
	}
	s.Breed(childrun)
	return s
}
