package gumi

import (
	"image"
	"fmt"
	"github.com/iamGreedy/gumi/gumre"
)

type ASpacer struct {
	VoidStructure
	//
	verical gumre.Length
	horizontal gumre.Length
}

func (ASpacer) GUMIRender(frame *image.RGBA) {

}
func (s *ASpacer) GUMISize() gumre.Size {
	return gumre.Size{
		Vertical:s.verical,
		Horizontal:s.horizontal,
	}
}
func (ASpacer) GUMIClip(r image.Rectangle) {
}
func (ASpacer) GUMIUpdate(info *Information, style *Style) {
}
func (ASpacer) GUMIHappen(event Event) {
}
func (ASpacer) String() string {
	return fmt.Sprintf("%s", "ASpacer")
}
//
func ASpacer0(horizontal, vertical gumre.Length) *ASpacer {
	return &ASpacer{
		horizontal:horizontal,
		verical:vertical,
	}
}
func ASpacer1(horizontal gumre.Length) *ASpacer {
	return &ASpacer{
		horizontal:horizontal,
		verical:gumre.AUTOLENGTH,
	}
}
func ASpacer2(vertical gumre.Length) *ASpacer {
	return &ASpacer{
		horizontal:gumre.AUTOLENGTH,
		verical:vertical,
	}
}
//
func (s *ASpacer) Get()(horizontal, vertical gumre.Length) {
	return s.horizontal, s.verical
}
func (s *ASpacer) Set(horizontal, vertical gumre.Length) {
	s.horizontal, s.verical = horizontal, vertical
}
func (s *ASpacer) GetHorizontal()(gumre.Length) {
	return s.horizontal
}
func (s *ASpacer) SetHorizontal(horizontal gumre.Length) {
	s.horizontal = horizontal
}
func (s *ASpacer) GetVertical()(gumre.Length) {
	return s.verical
}
func (s *ASpacer) SetVertical(vertical gumre.Length) {
	s.verical = vertical
}
