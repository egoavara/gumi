package uniplm

import (
	"image"
)

type nHorizontal struct {
	GUMILINK_MULTIPLE
}

func (s *nHorizontal) size(drawing *Drawing, style *Style) Size {
	size := Size{
		Horizontal: AUTOLENGTH,
		Vertical:   AUTOLENGTH,
	}
	for _, v := range s.child {
		temp := v.(GUMILink).size(drawing, style)
		size.Horizontal.Min += temp.Horizontal.Min
	}
	return size
}
func (s *nHorizontal) draw(drawing *Drawing, style *Style, frame Frame) {
	var start uint16 = 0
	var snc = make(chan struct{})
	var cnt = len(s.child) - 1
	//
	for i, v := range s.child {
		v2 := v.(GUMILink)
		temp := v2.size(drawing, style)
		sf := frame.SubFrame(
			image.Rect(int(start), 0, int(start+temp.Horizontal.Min), int(temp.Vertical.Max)),
		)
		if !sf.Bounds().Empty() {
			go func(a int) {
				v2.draw(drawing, style, sf)
				if cnt == a {
					close(snc)
				}
			}(i)
		}
		start += temp.Horizontal.Min
	}
	for range snc {
	}
}

func NHorizontal(childrun ...GUMILinker) *nHorizontal {
	s := &nHorizontal{}
	s.Link(childrun...)
	return s
}
