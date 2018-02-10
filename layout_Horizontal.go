package gumi

import (
	"fmt"
	"github.com/iamGreedy/gumi/gumre"
	"image"
	"sync"
)

type LHorizontal struct {
	MultipleStructure
	rule gumre.Distribute
}

func (s *LHorizontal) String() string {
	return fmt.Sprintf("%s(childrun:%d)", "LHorizontal", len(s.Childrun()))
}
func (s *LHorizontal) GUMIRender(frame *image.RGBA) {
	wg := new(sync.WaitGroup)
	wg.Add(len(s.child))
	defer wg.Wait()
	for _, v := range s.child {
		go func(elem GUMI) {
			elem.GUMIRender(frame)
			wg.Done()
		}(v)
	}
}
func (s *LHorizontal) GUMISize() gumre.Size {
	var minMax, sum uint16 = 0, 0
	for _, v := range s.child {
		sz := v.GUMISize()
		if sz.Vertical.Min > minMax {
			minMax = sz.Vertical.Min
		}
		sum += sz.Horizontal.Min
	}
	return gumre.Size{
		gumre.MinLength(minMax),
		gumre.MinLength(sum),
	}
}
func (s *LHorizontal) GUMIClip(r image.Rectangle) {
	//
	var tempVert = make([]gumre.Length, len(s.child))
	var tempHori = make([]gumre.Length, len(s.child))

	for i, v := range s.child {
		tempVert[i] = v.GUMISize().Vertical
		tempHori[i] = v.GUMISize().Horizontal
	}
	dis := s.rule(r.Dx(), tempHori)
	//
	var startat = r.Min.X
	for i, v := range s.child {
		r := image.Rect(
			startat,
			r.Min.Y,
			startat+dis[i],
			r.Max.Y,
		)
		v.GUMIClip(r)
		startat += dis[i]
	}
}
func (s *LHorizontal) GUMIUpdate(info *Information, style *Style) {
	for _, v := range s.child {
		v.GUMIUpdate(info, style)
	}
}
func (s *LHorizontal) GUMIHappen(event Event) {
	for _, v := range s.child {
		v.GUMIHappen(event)
	}
}

func LHorizontal0(rule gumre.Distribute, childrun ...GUMI) *LHorizontal {
	s := &LHorizontal{
		rule: rule,
	}
	for _, v := range childrun {
		v.born(s)
	}
	s.breed(childrun)
	return s
}
func LHorizontal1(childrun ...GUMI) *LHorizontal {
	s := &LHorizontal{
		rule: gumre.Distribution.Minimalize,
	}
	for _, v := range childrun {
		v.born(s)
	}
	s.breed(childrun)
	return s
}

func (s *LHorizontal) LoadElements(index gumre.Index, count int) []GUMI {
	return loadGUMIChildrun(s.child, index, count)
}
func (s *LHorizontal) SizeElements() int {
	return len(s.child)
}
func (s *LHorizontal) SaveElements(mode gumre.Mode, index gumre.Index, elem ...GUMI) (input int) {
	return saveGUMIChildrun(&s.child, mode, index, elem...)
}


