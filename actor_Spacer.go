package gumi

import (
	"image"
	"fmt"
)

type ASpacer struct {
	VoidStructure
	//
	verical Length
	horizontal Length
}

func (ASpacer) draw(frame *image.RGBA) {

}
func (s *ASpacer) size() Size {
	return Size{
		Vertical:s.verical,
		Horizontal:s.horizontal,
	}
}
func (ASpacer) rect(r image.Rectangle) {
}
func (ASpacer) update(info *Information, style *Style) {
}
func (ASpacer) Occur(event Event) {
}
func (ASpacer) String() string {
	return fmt.Sprintf("%s", "ASpacer")
}
//
func ASpacer0(horizontal, vertical Length) *ASpacer {
	return &ASpacer{
		horizontal:horizontal,
		verical:vertical,
	}
}
func ASpacer1(horizontal Length) *ASpacer {
	return &ASpacer{
		horizontal:horizontal,
		verical:AUTOLENGTH,
	}
}
func ASpacer2(vertical Length) *ASpacer {
	return &ASpacer{
		horizontal:AUTOLENGTH,
		verical:vertical,
	}
}
//
func (s *ASpacer) Get()(horizontal, vertical Length) {
	return s.horizontal, s.verical
}
func (s *ASpacer) Set(horizontal, vertical Length) {
	s.horizontal, s.verical = horizontal, vertical
}
func (s *ASpacer) GetHorizontal()(Length) {
	return s.horizontal
}
func (s *ASpacer) SetHorizontal(horizontal Length) {
	s.horizontal = horizontal
}
func (s *ASpacer) GetVertical()(Length) {
	return s.verical
}
func (s *ASpacer) SetVertical(vertical Length) {
	s.verical = vertical
}
