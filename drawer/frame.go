package drawer

import (
	"image"
	"image/draw"
)

type RenderTree struct {
	node  map[ID]*RenderNode
	frame *image.RGBA
	space IDSpace
	root  *RenderNode
}

func NewRenderTree(frame *image.RGBA) *RenderTree {
	temp := &RenderTree{}
	temp.node = make(map[ID]*RenderNode)
	temp.frame = frame
	temp.space = IDSpace{}
	return temp
}
func (s *RenderTree) New(parent *RenderNode) *RenderNode {
	temp := &RenderNode{}
	temp.tree = s
	temp.parent = parent
	temp.childrun = nil
	temp.ID = s.space.New()
	temp.require = true
	// temp.rect
	s.node[temp.ID] = temp
	if parent != nil {
		parent.appendChild(temp)
		temp.parent = parent
	} else {
		if s.root != nil {
			panic(CriticalRootOverride)
		}
		s.root = temp
	}
	return temp
}
func (s *RenderTree) Root() *RenderNode {
	return s.root
}
func (s *RenderTree) Rect() image.Rectangle {
	return s.frame.Rect
}
func (s *RenderTree) Frame() *image.RGBA {
	return s.frame
}

type RenderNode struct {
	tree     *RenderTree
	parent   *RenderNode
	childrun []*RenderNode
	ID       ID
	//
	rect  image.Rectangle
	cache *image.RGBA
	cachevalid bool
	Defer func()
	//
	require  bool
}

func (s *RenderNode) appendChild(node *RenderNode) {
	for _, v := range s.childrun {
		if v.ID == node.ID {
			return
		}
	}
	s.childrun = append(s.childrun, node)
}

//
func (s *RenderNode) FullImage() *image.RGBA {
	return s.tree.Frame()
}
func (s *RenderNode) SubImage() *image.RGBA {
	return s.tree.Frame().SubImage(s.rect).(*image.RGBA)
}
//
func (s *RenderNode) PopCache() {
	draw.Draw(s.tree.frame, s.rect, s.cache, s.cache.Rect.Min, draw.Src)
}
func (s *RenderNode) PushCache() {
	s.cachevalid = true
	draw.Draw(s.cache, s.cache.Rect, s.tree.frame, s.rect.Min, draw.Src)
}
func (s *RenderNode) ValidCache() bool {
	return s.cachevalid
}
func (s *RenderNode) ThrowCache() {
	s.cachevalid = false
}
//
func (s *RenderNode) Size() image.Rectangle {
	return s.rect
}

//
func (s *RenderNode) Parent() *RenderNode {
	return s.parent
}
func (s *RenderNode) DeferDo() {
	for _, v := range s.childrun {
		v.Defer()
		v.DeferDo()
	}
}

//
func (s *RenderNode) SetRect(rectangle image.Rectangle) {
	if s.rect != rectangle {
		s.rect = rectangle
		s.cache = image.NewRGBA(image.Rect(0,0,s.rect.Dx(), s.rect.Dy()))
		s.Require()
	}
}
func (s *RenderNode) GetRect() image.Rectangle {
	return s.rect
}

func (s *RenderNode) Complete() {
	s.require = false
}

func (s *RenderNode) Require() {
	s.require = true
	for _, v := range s.childrun {
		v.Require()
	}
}
func (s *RenderNode) Check() bool {
	if s.recurCheck() {
		s.require = false
		return true
	}
	return false
}
func (s *RenderNode) recurCheck() bool {
	if s.require{
		return true
	}
	any := false
	for _, v := range s.childrun{
		any = any || v.recurCheck()
	}
	return any
}