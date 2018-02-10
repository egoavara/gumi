package gumi

import (
	"image"
	"image/draw"
	"fmt"
	"github.com/iamGreedy/gumi/gumre"
)

type AImage struct {
	VoidStructure
	boundStore
	//
	img image.Image
}

func (s *AImage) String() string {
	return fmt.Sprintf("%s", "AImage")
}

func (s *AImage) GUMIRender(frame *image.RGBA) {
	draw.Draw(frame, s.bound, s.img, s.img.Bounds().Min, draw.Over)
}

func (s AImage) GUMISize() gumre.Size {
	bd := s.img.Bounds()
	return gumre.Size{
		Horizontal: gumre.MinLength(uint16(bd.Dx())),
		Vertical:   gumre.MinLength(uint16(bd.Dy())),
	}
}
func (s *AImage) GUMIClip(rect image.Rectangle) {
	s.bound = rect
}

func (s *AImage) GUMIUpdate(info *Information, style *Style) {

}

func (s *AImage) GUMIHappen(event Event) {
}

func AImage0(img image.Image) *AImage {
	return &AImage{
		img: img,
	}
}
