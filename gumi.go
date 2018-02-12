package gumi

import (
	"image"
	"fmt"
	"github.com/iamGreedy/gumi/gumre"
	"github.com/iamGreedy/gumi/drawer"
)

type GUMI interface{
	GUMIRenderer
	GUMIStructure
	GUMICacher
	GUMIEventer
	fmt.Stringer
}

type GUMIRoot interface{
	GUMI
	Screen() *Screen
}

type GUMIRenderer interface{
	GUMIInit()
	GUMIInfomation(info Information)
	GUMIStyle(style *Style)
	GUMIClip(r image.Rectangle)
	GUMIRender(frame *image.RGBA)
	GUMISize() gumre.Size
	GUMIDraw(frame *image.RGBA)
}
type GUMIStructure interface{
	born(gumi GUMI)
	breed(gumi []GUMI)
	Parent() GUMI
	Childrun() []GUMI
}
type GUMICacher interface{
	GUMIRenderTree(tree *drawer.RenderTree, parentnode *drawer.RenderNode)
	GUMIUpdate()
}
type GUMIEventer interface{
	GUMIHappen(event Event)
}
