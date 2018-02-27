package gumi

import (
	"fmt"
	"image"
	"github.com/iamGreedy/gumi/gumre"
	"github.com/iamGreedy/gumi/drawer"
)

// _::Root
//
// gumiRoot is Special case. private struct
type gumiRoot struct {
	SingleNode
	scr *Screen
}

// GUMIFunction / GUMIInit 					-> SingleNode::Default

// GUMIFunction / GUMIInfomation 			-> Define
func (s *gumiRoot) GUMIInfomation(info Information) {
	s.child.GUMIInfomation(info)
}

// GUMIFunction / GUMIStyle 			-> Define
func (s *gumiRoot) GUMIStyle(style *Style) {
	s.child.GUMIStyle(style)
}

// GUMIFunction / GUMIClip 			-> Define
func (s *gumiRoot) GUMIClip(r image.Rectangle) {
	s.child.GUMIClip(r)
}

// GUMIFunction / GUMIRender 			-> Define
func (s *gumiRoot) GUMIRender(frame *image.RGBA) {

}

// GUMIFunction / GUMISize 			-> Define
func (s *gumiRoot) GUMISize() gumre.Size {
	return s.child.GUMISize()
}


// GUMITree / born 							-> SingleNode::Default

// GUMITree / breed 						-> SingleNode::Default

// GUMITree / Parent()						-> SingleNode::Default

// GUMITree / Childrun()					-> SingleNode::Default


// GUMIRenderer / GUMIRenderSetup			-> Define::Empty
func (s *gumiRoot) GUMIRenderSetup(frame *image.RGBA, tree *drawer.RenderTree, parentnode *drawer.RenderNode) {

}

// GUMIRenderer / GUMIRenderSetup			-> Define
func (s *gumiRoot) GUMIDraw() {
	s.child.GUMIDraw()
}

// GUMIRenderer / GUMIRenderSetup			-> Define
func (s *gumiRoot) GUMIUpdate() {
	// TODO
	panic("implement me")
}

// GUMIEventor
func (s *gumiRoot) GUMIHappen(event Event) {
	s.child.GUMIHappen(event)
}


func (s *gumiRoot) String() string {
	return fmt.Sprintf("%s", "GUMI Root")
}

// GUMIRoot / Screen
func (s *gumiRoot) Screen() *Screen {
	return s.scr
}

// Constructor
func newGUMIRoot(scr *Screen, under GUMI) GUMIRoot {
	temp := &gumiRoot{
		scr: scr,
	}
	LinkingFrom(temp, under)
	return temp
}

// Utility Function
func Root(g GUMI) GUMIRoot {
	if g == nil{
		return nil
	}
	if v, ok := g.(*gumiRoot); ok{
		return v
	}
	return Root(g.Parent())
}
