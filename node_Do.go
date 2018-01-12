package gumi

type nDo struct {
	GUMILINK_SINGLE
	fn func()
}

func (s *nDo) size(drawing *Drawing, style *Style) Size {
	return s.child.(GUMIElem).size(drawing, style)
}
func (s *nDo) draw(drawing *Drawing, style *Style, frame Frame) {
	s.fn()
	s.child.(GUMIElem).draw(drawing, style, frame)
}
func NDo(fn func()) *nDo {
	return &nDo{
		fn: fn,
	}
}
