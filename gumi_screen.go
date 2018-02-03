package gumi

import (
	"image"
)

type Screen struct {
	frame *image.RGBA
	//
	//
	root GUMI
}

func NewScreen(w, h int) *Screen {
	res := &Screen{
		frame: image.NewRGBA(image.Rect(0, 0, w, h)),
	}
	return res
}
func (s *Screen) Width() int {
	return s.frame.Rect.Dx()
}
func (s *Screen) Height() int {
	return s.frame.Rect.Dy()
}
func (s *Screen) Size() (width, height int) {
	return s.Width(), s.Height()
}
func (s *Screen) Resize(w, h int) {
	s.frame = image.NewRGBA(image.Rect(0, 0, w, h))
}
func (s *Screen) Root(root GUMI) {
	s.root = root
}
func (s *Screen) Event(event Event) {
	if event == nil {
		return
	}
	s.root.Occur(event)
}
func (s *Screen) Draw() {
	s.root.draw(s.frame)
}
func (s *Screen) Ready() {
	if s.root == nil {
		return
	}
	s.root.rect(s.frame.Rect)
}
func (s *Screen) Update(info *Information, style *Style) {
	if info == nil {
		info = DefaultInformation()
	}
	if style == nil {
		style = DefaultStyle()
	}
	s.root.update(info, style)
}
func (s *Screen) Frame() image.Image {
	return s.frame
}
func (s *Screen) RGBA() *image.RGBA {
	return s.frame
}
