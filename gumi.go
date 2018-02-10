package gumi

import (
	"image"
	"fmt"
	"github.com/iamGreedy/gumi/gumre"
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
	GUMIRender(frame *image.RGBA)
	GUMISize() gumre.Size
	GUMIClip(r image.Rectangle)
	GUMIUpdate(info *Information, style *Style)
	GUMIInit()
}
type GUMIStructure interface{
	born(gumi GUMI)
	breed(gumi []GUMI)
	Parent() GUMI
	Childrun() []GUMI
}
type GUMICacher interface{
	// TODO GUMICacher
	//Changed()
}
type GUMIEventer interface{
	GUMIHappen(event Event)
}
