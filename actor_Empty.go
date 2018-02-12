package gumi

import (
	"fmt"
	"github.com/iamGreedy/gumi/gumre"
	"image"
	"github.com/iamGreedy/gumi/drawer"
)

type AEmpty struct {
	VoidStructure
}

func (s AEmpty) GUMIInfomation(info Information) {
}
func (s AEmpty) GUMIStyle(style *Style) {
}
func (AEmpty) GUMIClip(image.Rectangle) {
}
func (AEmpty) GUMIRender(frame *image.RGBA) {
}
func (s AEmpty) GUMIDraw(frame *image.RGBA) {
}

func (s AEmpty) GUMIRenderTree(tree *drawer.RenderTree, parentnode *drawer.RenderNode) {
	// TODO
	panic("implement me")
}

func (s AEmpty) GUMIUpdate() {
	// TODO
	panic("implement me")
}


func (AEmpty) GUMISize() gumre.Size {
	return gumre.Size{
		gumre.AUTOLENGTH,
		gumre.AUTOLENGTH,
	}
}
func (AEmpty) GUMIHappen(event Event) {
}
func (s AEmpty) String() string {
	return fmt.Sprintf("%s", "AEmpty")
}

//
func AEmpty0() *AEmpty {
	return &AEmpty{}
}
