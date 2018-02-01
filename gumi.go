package gumi

import (
	"image"
	"fmt"
)

type GUMI interface{
	GUMIRenderer
	GUMIStructure
	GUMICacher
	GUMIEventer
	fmt.Stringer
}

type GUMIRenderer interface{
	draw(frame *image.RGBA)
	size() Size
	rect(r image.Rectangle)
	update(info *Information, style *Style)
}
type GUMIStructure interface{
	Born(gumi GUMI)
	Breed(gumi []GUMI)
	Parent() GUMI
	Childrun() []GUMI
}
type GUMICacher interface{
	// TODO GUMICacher
	//Changed()
}
type GUMIEventer interface{
	Occur(event Event)
}
