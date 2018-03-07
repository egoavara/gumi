package temp

import (
	"fmt"
	"github.com/iamGreedy/gumi/gumre"
	"image"
	"sync"
	"github.com/iamGreedy/gumi/drawer"
)

// Layout::Horizontal
//
// Horizontal align
type LHorizontal struct {
	MultipleNode
	rule gumre.Distribute
}

// GUMIFunction / GUMIInit 					-> SingleNode::Default

// GUMIFunction / GUMIInfomation 			-> Define
func (s *LHorizontal) GUMIInfomation(info Information) {
	for _, v := range s.child {
		v.GUMIInfomation(info)
	}
}

// GUMIFunction / GUMIStyle 				-> Define
func (s *LHorizontal) GUMIStyle(style *Style) {
	for _, v := range s.child {
		v.GUMIStyle(style)
	}
}

// GUMIFunction / GUMIClip 					-> Define
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

// GUMIFunction / GUMIRender 				-> Define::Empty
func (s *LHorizontal) GUMIRender(frame *image.RGBA) {
}

// GUMIFunction / GUMISize 					-> Define
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

// GUMITree / born 							-> MultipleNode::Default

// GUMITree / breed 						-> MultipleNode::Default

// GUMITree / parent()						-> MultipleNode::Default

// GUMITree / childrun()					-> MultipleNode::Default

// GUMIRenderer / GUMIRenderSetup 			-> Define
func (s *LHorizontal) GUMIRenderSetup(frame *image.RGBA, tree *drawer.RenderTree, parentnode *drawer.RenderNode) {
	for _, v := range s.child {
		v.GUMIRenderSetup(frame, tree, parentnode)
	}
}

// GUMIRenderer / GUMIDraw 					-> Define
func (s *LHorizontal) GUMIDraw() {
	wg := new(sync.WaitGroup)
	wg.Add(len(s.child))
	defer wg.Wait()
	for _, v := range s.child {
		go func(elem GUMI) {
			elem.GUMIDraw()
			wg.Done()
		}(v)
	}
}

// GUMIRenderer / GUMIUpdate				-> Define
func (s *LHorizontal) GUMIUpdate() {
	panic("implement me")
}

// GUMIEventer / GUMIHappen					-> Define
func (s *LHorizontal) GUMIHappen(event Event) {
	for _, v := range s.child {
		v.GUMIHappen(event)
	}
}

// fmt.Stringer / String					-> Define
func (s *LHorizontal) String() string {
	return fmt.Sprintf("%s(childrun:%d)", "LHorizontal", len(s.Childrun()))
}


// Constructor 0
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

// Constructor 1
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


