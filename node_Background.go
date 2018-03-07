package gumi

import (
	"fmt"
	"image"
	"github.com/iamGreedy/gumi/gumre"
	"github.com/iamGreedy/gumi/drawer"
)

type NBackground struct {
	SingleNode
	rendererStore
	//
	drawer drawer.Drawer
}

func (s *NBackground) GUMIInfomation(info Information) {
	s.child.GUMIInfomation(info)
}
func (s *NBackground) GUMIStyle(style *Style) {
	s.child.GUMIStyle(style)
}
func (s *NBackground) GUMIClip(rect image.Rectangle) {
	s.rnode.SetRect(rect)
	s.child.GUMIClip(rect)
}
func (s *NBackground) GUMIRender(frame *image.RGBA) {
	s.drawer.Draw(frame)

}
func (s NBackground) GUMISize() gumre.Size {
	return s.child.GUMISize()
}

func (s *NBackground) GUMIRenderSetup(tree *drawer.RenderTree, parentnode *drawer.RenderNode) {
	s.rtree = tree
	s.rnode = parentnode
	s.child.GUMIRenderSetup(tree, s.rnode)
}
func (s *NBackground) GUMIUpdate() {
	if s.rnode.Check(){
		if s.rnode.ValidCache(){
			s.rnode.PopCache()
		}else {
			s.GUMIRender(s.rnode.SubImage())
			s.rnode.PushCache()
		}
	}
	s.child.GUMIUpdate()
}

func (s *NBackground) GUMIHappen(event Event) {
	s.child.GUMIHappen(event)
}

func (s *NBackground) String() string {
	return fmt.Sprintf("%s", "NBackground")
}

func NBackground0(draw drawer.Drawer) *NBackground {
	return &NBackground{
		drawer:draw,
	}
}

func (s *NBackground) Set(dw drawer.Drawer) {
	s.SetDrawer(dw)
}
func (s *NBackground) Get() drawer.Drawer {
	return s.GetDrawer()
}

func (s *NBackground) SetDrawer(dw drawer.Drawer) {
	s.drawer = dw
}

func (s *NBackground) GetDrawer() drawer.Drawer {
	return s.drawer
}