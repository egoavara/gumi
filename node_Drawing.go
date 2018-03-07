package gumi

import (
	"image"
	"fmt"
	"github.com/iamGreedy/gumi/gumre"
	"github.com/iamGreedy/gumi/drawer"
)

type NDrawing struct {
	SingleNode
	styleStore
	rendererStore
	//
	drawfuncs []Drawer
}

func (s *NDrawing) GUMIInfomation(info Information) {
	var changed bool
	for _, v := range s.drawfuncs{
		if v2, ok := v.(DrawerWithInformation); ok{
			changed = changed || v2.Inform(info)
		}
	}
	if changed{
		s.rnode.Require()
	}
	s.child.GUMIInfomation(info)
}
func (s *NDrawing) GUMIStyle(style *Style) {
	s.style = style
	s.child.GUMIStyle(style)
}
func (s *NDrawing) GUMIClip(r image.Rectangle) {
	s.rnode.SetRect(r)
	s.child.GUMIClip(r)
}
func (s *NDrawing) GUMIRender(frame *image.RGBA) {
}
func (s *NDrawing) GUMIRenderSetup(tree *drawer.RenderTree, parentnode *drawer.RenderNode) {
	s.rtree = tree
	s.rnode = tree.New(parentnode)
	s.child.GUMIRenderSetup(tree, s.rnode)
}
func (s *NDrawing) GUMIUpdate() {
	s.child.GUMIUpdate()
	var ctx = createContext(s.rnode.SubImage())
	for _, f := range s.drawfuncs {
		ctx.Push()
		f.Draw(ctx, s.style)
		ctx.Pop()
	}
}

func (s *NDrawing) GUMIHappen(event Event) {
	s.child.GUMIHappen(event)
}
func (s *NDrawing) GUMISize() gumre.Size {
	return s.child.GUMISize()
}
func (s *NDrawing) String() string {
	return fmt.Sprintf("%s(drawing:%d GUMIRender)", "NDrawing", len(s.drawfuncs))
}
//
func NDrawing0(drawFuncs ...Drawer) *NDrawing {
	return &NDrawing{
		drawfuncs: drawFuncs,
	}
}
func NDrawing1(drawFuncs ...Drawer) *NDrawing {
	return &NDrawing{
		drawfuncs: drawFuncs,
	}
}
