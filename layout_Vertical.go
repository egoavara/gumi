package gumi

import (
	"image"
	"fmt"
	"sync"
	"github.com/iamGreedy/gumi/gumre"
	"github.com/iamGreedy/gumi/drawer"
)

// Layout::Horizontal
//
// Horizontal align
type LVertical struct {
	MultipleNode
	rendererStore
	rule gumre.Distribute
}

// GUMIFunction / GUMIInit 					-> SingleNode::Default

// GUMIFunction / GUMIInfomation 			-> Define
func (s *LVertical) GUMIInfomation(info Information) {
	for _, v := range s.child{
		v.GUMIInfomation(info)
	}
}

// GUMIFunction / GUMIStyle 				-> Define
func (s *LVertical) GUMIStyle(style *Style) {
	for _, v := range s.child{
		v.GUMIStyle(style)
	}
}

// GUMIFunction / GUMIClip 					-> Define
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

// GUMIFunction / GUMIRender 				-> Define::Empty
func (s *LVertical) GUMIRender(frame *image.RGBA) {

}

// GUMIFunction / GUMISize 					-> Define
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

// GUMITree / born 							-> MultipleNode::Default

// GUMITree / breed 						-> MultipleNode::Default

// GUMITree / parent()						-> MultipleNode::Default

// GUMITree / childrun()					-> MultipleNode::Default

// GUMIRenderer / GUMIRenderSetup 			-> Define::Empty
func (s *LVertical) GUMIRenderSetup(tree *drawer.RenderTree, parentnode *drawer.RenderNode) {
	s.rtree = tree
	s.rnode = tree.New(parentnode)
	for _, v := range s.child{
		v.GUMIRenderSetup(s.rtree, s.rnode)
	}
}

// GUMIRenderer / GUMIUpdate 				-> Define
func (s *LVertical) GUMIUpdate() {
	if s.rnode.Check(){
		s.rnode.Require()
		wg := new(sync.WaitGroup)
		wg.Add(len(s.child))
		defer wg.Wait()
		for _, v := range s.child{
			go func(elem GUMI) {
				elem.GUMIUpdate()
				wg.Done()
			}(v)
		}
		s.rnode.Complete()
	}
}

// GUMIEventer / GUMIHappen					-> Define
func (s *LVertical) GUMIHappen(event Event) {
	for _, v := range s.child{
		go v.GUMIHappen(event)
	}
}

// fmt.Stringer / String					-> Define
func (s *LVertical) String() string{
	return fmt.Sprintf(
		"%s(childrun:%d)", "LVertical", len(s.Childrun()),
	)
}

// Constructor 0
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

// Constructor 1
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

// Get Elems
func (s *LVertical) LoadElements(index gumre.Index, count int) []GUMI {
	return loadGUMIChildrun(s.child, index, count)
}

// Count Elems
func (s *LVertical) SizeElements() int {
	return len(s.child)
}

// Set Elems
func (s *LVertical) SaveElements(mode gumre.Mode, index gumre.Index, elem ...GUMI) (input int) {
	return saveGUMIChildrun(&s.child, mode, index, elem...)
}