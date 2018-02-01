package gumi

import (
	"image"
	"fmt"
)

type AEmpty struct {
	VoidStructure
}

func (AEmpty) draw(frame *image.RGBA) {
}

func (AEmpty) size() Size {
	return Size{
		AUTOLENGTH,
		AUTOLENGTH,
	}
}

func (AEmpty) rect(image.Rectangle) {
}

func (AEmpty) update(info *Information, style *Style) {
}

func (AEmpty) Occur(event Event) {
}
func (s AEmpty) String() string {
	return fmt.Sprintf("%s", "AEmpty")
}
//
func AEmpty0() *AEmpty {
	return &AEmpty{}
}
