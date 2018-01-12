package gumi

import (
	"image"
)

type nVertical struct {
	GUMILINK_MULTIPLE
}

func (s *nVertical) size(drawing *Drawing, style *Style) Size {
	size := Size{
		Horizontal: AUTOLENGTH,
		Vertical:   AUTOLENGTH,
	}
	for _, v := range s.child {
		temp := v.(GUMILink).size(drawing, style)
		size.Vertical.Min += temp.Vertical.Min
	}
	return size
}
func (s *nVertical) draw(drawing *Drawing, style *Style, frame Frame) {
	var start uint16 = 0
	var snc = make(chan struct{})
	var cnt = len(s.child) - 1
	//
	for i, v := range s.child {
		v2 := v.(GUMILink)
		temp := v2.size(drawing, style)
		sf := frame.SubFrame(
			image.Rect(0, int(start), int(temp.Horizontal.Max), int(start+temp.Vertical.Min)),
		)
		if !sf.Bounds().Empty() {
			go func(a int) {
				v2.draw(drawing, style, sf)
				if cnt == a {
					close(snc)
				}
			}(i)
		}
		start += temp.Vertical.Min
	}
	for range snc {
	}
}

func NVertical(childrun ...GUMILinker) *nVertical {
	s := &nVertical{}
	s.Link(childrun...)
	return s
}
