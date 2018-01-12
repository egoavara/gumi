package gumi

import (
	"image"
)

type Screen struct {
	w, h int
	//
	frame Frame
	//
	cache *FrameCache
	//
	root GUMIElem
}

func NewScreen(w, h int) *Screen {
	return &Screen{
		w:     w,
		h:     h,
		frame: NewFrame(w, h),
	}
}
func (s *Screen) Root(root GUMIElem) {
	s.root = root
}
func (s *Screen) Draw(drawing *Drawing) {
	if drawing == nil {
		drawing = &Drawing{
			Dt:         0,
			ThrowCache: false,
			Level: 0,
		}
	}
	if s.root == nil {
		return
	}
	s.root.draw(drawing, DefaultStyle, s.frame)
}

func (s *Screen) Frame() image.Image {
	return s.frame
}
