package gumi

import "image"

type aEmpty struct {
	VoidStructure
}

func (aEmpty) draw(frame *image.RGBA) {
}

func (aEmpty) size() Size {
	return Size{
		AUTOLENGTH,
		AUTOLENGTH,
	}
}

func (aEmpty) rect(image.Rectangle) {
}

func (aEmpty) update(info *Information, style *Style) {
}

func (aEmpty) Occur(event Event) {
}

//
func AEmpty() *aEmpty {
	return &aEmpty{}
}
