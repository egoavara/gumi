package gumi

import (
	"image"
	"image/draw"
)

type aImage struct {
	VoidStructure
	BoundStore
	//
	img image.Image
}

func (s *aImage) draw(frame *image.RGBA) {
	draw.Draw(frame, s.bound, s.img, s.img.Bounds().Min, draw.Over)
}

func (s aImage) size() Size {
	bd := s.img.Bounds()
	return Size{
		Horizontal: MinLength(uint16(bd.Dx())),
		Vertical:   MinLength(uint16(bd.Dy())),
	}
}
func (s *aImage) rect(rect image.Rectangle) {
	s.bound = rect
}

func (s *aImage) update(info *Information, style *Style) {
}

func (s *aImage) Occur(event Event) {
}

func AImage(img image.Image) *aImage {
	return &aImage{
		img: img,
	}
}
