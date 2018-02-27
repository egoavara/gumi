// gumi.go define GUMI interface
// GUMI is elements of GUI
// Every elements can render(graphic), affect(event), update(information) must implements this interface
package gumi

import (
	"fmt"
	"github.com/iamGreedy/gumi/drawer"
	"github.com/iamGreedy/gumi/gumre"
	"image"
)

// GUMI is a collection of basic elements
type GUMI interface {
	GUMIFunction
	GUMITree
	GUMIRenderer
	GUMIEventer

	// TODO : GUMIRelay

	fmt.Stringer
}

// GUMI Root is special case of GUMI
// GUMI Root help to find Screen
// Mostly root locate root position on GUMI Tree
// But it is not necessary
type GUMIRoot interface {
	GUMI
	Screen() *Screen
}

type GUMIFunction interface {
	GUMIInit()                       // TODO : Relay
	GUMIInfomation(info Information) // TODO : Relay
	GUMIStyle(style *Style)          // TODO : Relay
	GUMIClip(r image.Rectangle)      // TODO : Relay
	GUMIRender(frame *image.RGBA)
	GUMISize() gumre.Size
}
type GUMITree interface {
	born(gumi GUMI)
	breed(gumi []GUMI)
	Parent() GUMI
	Childrun() []GUMI
}
type GUMIRenderer interface {
	GUMIRenderSetup(frame *image.RGBA, tree *drawer.RenderTree, parentnode *drawer.RenderNode)
	GUMIUpdate()
	GUMIDraw()
}
type GUMIEventer interface {
	GUMIHappen(event Event)
}

// TODO : Relay
type GUMIRelay interface {
	GUMIRelayInit()
	GUMIRelayInfomation(info Information)
	GUMIRelayStyle(style *Style)
	GUMIRelayClip(r image.Rectangle)
}
