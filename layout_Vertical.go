package gumi

import (
	"image"
	"math"
	"fmt"
	"sync"
)

type LVertical struct {
	MultipleStructure
	rule Distribute
}

func (s *LVertical) draw(frame *image.RGBA) {
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
func (s *LVertical) size() Size {
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
func (s *LVertical) rect(r image.Rectangle) {
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
		inrect := image.Rect(
			r.Min.X,
			startat,
			r.Max.X,
			startat + dis[i],
		)
		if int(tempHori[i].Max) < dx{
			inrect.Max.X = r.Min.X + int(tempHori[i].Max)
		}
		v.rect(inrect)
		startat += dis[i]
	}
}
func (s *LVertical) update(info *Information, style *Style) {
	for _, v := range s.child{
		v.update(info, style)
	}
}
func (s *LVertical) Occur(event Event) {
	for _, v := range s.child{
		go v.Occur(event)
	}
}
func (s *LVertical) String() string{
	return fmt.Sprintf(
		"%s(childrun:%d)", "LVertical", len(s.Childrun()),
	)
}

func LVertical0(rule Distribute, childrun ...GUMI) *LVertical {
	s := &LVertical{
		rule:rule,
	}
	for _, v := range childrun{
		v.Born(s)
	}
	s.Breed(childrun)
	return s
}
func LVertical1(childrun ...GUMI) *LVertical {
	s := &LVertical{
		rule: Distribution.Minimalize,
	}
	for _, v := range childrun{
		v.Born(s)
	}
	s.Breed(childrun)
	return s
}
