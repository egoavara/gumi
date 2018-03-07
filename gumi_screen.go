package gumi

import (
	"image"
	"math/rand"
	"github.com/iamGreedy/gumi/drawer"
)

type Screen struct {
	rtree  *drawer.RenderTree
	rstyle *Style
	//
	_hook  map[uint64]func(event Event) Event
	_defer map[uint64]func(rgba *image.RGBA)
	//

	//
	root GUMIRoot
}

func NewScreen(w, h int) *Screen {
	res := &Screen{
		rtree: drawer.NewRenderTree(image.NewRGBA(image.Rect(0, 0, w, h))),
		rstyle: DefaultStyle(),
		_hook:  make(map[uint64]func(event Event) Event),
		_defer: make(map[uint64]func(rgba *image.RGBA)),
	}
	return res
}

func (s *Screen) Width() int {
	return s.rtree.Rect().Dx()
}
func (s *Screen) Height() int {
	return s.rtree.Rect().Dy()
}
func (s *Screen) Size() (width, height int) {
	return s.Width(), s.Height()
}
func (s *Screen) Root(root GUMI) {
	s.root = newGUMIRoot(s, root)
}
//
func (s *Screen) Event(event Event) {
	for _, v := range s._hook {
		if v != nil {
			event = v(event)
		}
	}
	if event == nil {
		return
	}
	s.root.GUMIHappen(event)
}
//
func (s *Screen) Init() {
	s.root.GUMIInit()
	s.root.GUMIRenderSetup(s.rtree, s.rtree.New(nil))
	//
	s.root.GUMIStyle(s.rstyle)
	s.root.GUMIClip(s.rtree.Rect())
}
func (s *Screen) Update(info Information) {
	s.root.GUMIInfomation(info)
}
func (s *Screen) Draw() {
	s.root.GUMIUpdate()
	for _, v := range s._defer {
		if v != nil {
			v(s.rtree.Frame())
		}
	}
}
func (s *Screen) Frame() image.Image {
	return s.rtree.Frame()
}
func (s *Screen) RGBA() *image.RGBA {
	return s.rtree.Frame()
}
//
func (s *Screen) hookReserve() (id uint64) {
	defer func() {
		s._hook[id] = nil
	}()
	for {
		id = rand.Uint64()
		if id == 0 {
			continue
		}
		if _, ok := s._hook[id]; !ok {
			return id
		}
	}
}
func (s *Screen) hookRequest(id uint64, hooking func(event Event) Event) {
	s._hook[id] = hooking
}
func (s *Screen) deferReserve() (id uint64) {
	defer func() {
		s._defer[id] = nil
	}()
	for {
		id = rand.Uint64()
		if id == 0 {
			continue
		}
		if _, ok := s._defer[id]; !ok {
			return
		}
	}
}
func (s *Screen) deferRequest(id uint64, d func(rgba *image.RGBA)) {
	s._defer[id] = d
}
