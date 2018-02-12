package drawer

import (
	"image"
	"sync/atomic"
)

type RenderTree struct {
	node map[ID]*RenderNode
	space IDSpade
}
type IDSpade struct {
	top uint32
}
func (s *IDSpade) New() ID {
	return ID(atomic.AddUint32(&s.top, 1))
}
type ID uint32

func (s *RenderTree) Add(id ID, node RenderNode) {
	s.node[id] = &node
}
type RenderNode struct {
	Cache Cache
	Rect image.Rectangle
	Depth int
}


