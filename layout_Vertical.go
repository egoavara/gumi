package gumi

import (
	"image"
	"fmt"
	"sync"
	"github.com/iamGreedy/gumi/gumre"
	"github.com/iamGreedy/gumi/drawer"
)

type LVertical struct {
	MultipleNode
	rule gumre.Distribute
}

func (s *LVertical) GUMIInfomation(info Information) {
	for _, v := range s.child{
		v.GUMIInfomation(info)
	}
}
func (s *LVertical) GUMIStyle(style *Style) {
	for _, v := range s.child{
		v.GUMIStyle(style)
	}
}
func (s *LVertical) GUMIClip(r image.Rectangle) {
	var tempVert = make([]gumre.Length, len(s.child))
	var tempHori = make([]gumre.Length, len(s.child))

	for i, v := range s.child{
		tempVert[i] = v.GUMISize().Vertical
		tempHori[i] = v.GUMISize().Horizontal
	}
	dis := s.rule(r.Dy(), tempVert)
	//
	var startat = r.Min.Y
	for i, v := range s.child{
		inrect := image.Rect(
			r.Min.X,
			startat,
			r.Max.X,
			startat + dis[i],
		)
		v.GUMIClip(inrect)
		startat += dis[i]
	}
}
func (s *LVertical) GUMIRender(frame *image.RGBA) {

}
func (s *LVertical) GUMIDraw(frame *image.RGBA) {
	wg := new(sync.WaitGroup)
	wg.Add(len(s.child))
	defer wg.Wait()
	for _, v := range s.child{
		go func(elem GUMI) {
			elem.GUMIDraw(frame)
			wg.Done()
		}(v)
	}
}

func (s *LVertical) GUMIRenderTree(tree *drawer.RenderTree, parentnode *drawer.RenderNode) {
	panic("implement me")
}
func (s *LVertical) GUMIUpdate() {
	panic("implement me")
}

func (s *LVertical) GUMIHappen(event Event) {
	for _, v := range s.child{
		go v.GUMIHappen(event)
	}
}
func (s *LVertical) GUMISize() gumre.Size {
	var minMax, sum uint16 = 0, 0
	for _, v := range s.child{
		sz := v.GUMISize()
		if sz.Horizontal.Min > minMax{
			minMax = sz.Horizontal.Min
		}
		sum += sz.Vertical.Min
	}
	return gumre.Size{
		gumre.MinLength(sum),
		gumre.MinLength(minMax),
	}
}
func (s *LVertical) String() string{
	return fmt.Sprintf(
		"%s(childrun:%d)", "LVertical", len(s.Childrun()),
	)
}

func LVertical0(rule gumre.Distribute, childrun ...GUMI) *LVertical {
	s := &LVertical{
		rule:rule,
	}
	for _, v := range childrun{
		v.born(s)
	}
	s.breed(childrun)
	return s
}
func LVertical1(childrun ...GUMI) *LVertical {
	s := &LVertical{
		rule: gumre.Distribution.Minimalize,
	}
	for _, v := range childrun{
		v.born(s)
	}
	s.breed(childrun)
	return s
}

func (s *LVertical) LoadElements(index gumre.Index, count int) []GUMI {
	return loadGUMIChildrun(s.child, index, count)
}
func (s *LVertical) SizeElements() int {
	return len(s.child)
}
func (s *LVertical) SaveElements(mode gumre.Mode, index gumre.Index, elem ...GUMI) (input int) {
	return saveGUMIChildrun(&s.child, mode, index, elem...)
}