package gumi

type FrameCache struct {
	originframe Frame
	node        []*FrameCacheNode
}
type FrameCacheNode struct {
	frame Frame
	style *Style
	elem  GUMIElem
}

func NewCache(w, h int) *FrameCache {
	return &FrameCache{
		originframe: NewFrame(w, h),
	}
}
func (s *FrameCache) Build(drawing *Drawing) Frame {
	temp := s.originframe.Copy()
	tempPix := temp.ReferPix()
	for _, v := range s.node {
		v.frame.Pix(tempPix)
		v.elem.draw(drawing, v.style, v.frame)
	}
	return temp
}
