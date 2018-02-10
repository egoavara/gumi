package gumi

import (
	"image"
	"fmt"
	"github.com/iamGreedy/gumi/gumre"
)

type AEmpty struct {
	VoidStructure
}

func (AEmpty) GUMIRender(frame *image.RGBA) {
}

func (AEmpty) GUMISize() gumre.Size {
	return gumre.Size{
		gumre.AUTOLENGTH,
		gumre.AUTOLENGTH,
	}
}

func (AEmpty) GUMIClip(image.Rectangle) {
}

func (AEmpty) GUMIUpdate(info *Information, style *Style) {
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
