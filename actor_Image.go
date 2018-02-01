package gumi

import (
	"image"
	"image/draw"
	"fmt"
)

type AImage struct {
	VoidStructure
	BoundStore
	//
	img image.Image
}

func (s *AImage) String() string {
	return fmt.Sprintf("%s", "AImage")
}

func (s *AImage) draw(frame *image.RGBA) {
	draw.Draw(frame, s.bound, s.img, s.img.Bounds().Min, draw.Over)
}

func (s AImage) size() Size {
	bd := s.img.Bounds()
	return Size{
		Horizontal: MinLength(uint16(bd.Dx())),
		Vertical:   MinLength(uint16(bd.Dy())),
	}
}
func (s *AImage) rect(rect image.Rectangle) {
	s.bound = rect
}

func (s *AImage) update(info *Information, style *Style) {

}

func (s *AImage) Occur(event Event) {
}

func AImage0(img image.Image) *AImage {
	return &AImage{
		img: img,
	}
}
