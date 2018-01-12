package gumi

type nStyle struct {
	GUMILINK_SINGLE
	s *Style
}

func NStyle(s *Style) *nStyle {
	if s == nil {
		s = DefaultStyle
	}
	return &nStyle{
		s: s,
	}
}
func (s *nStyle) size(drawing *Drawing, style *Style) Size {
	return s.child.(GUMIElem).size(drawing, s.s)
}
func (s *nStyle) draw(drawing *Drawing, style *Style, frame Frame) {
	s.child.(GUMIElem).draw(drawing, s.s, frame)
}
func (s *nStyle) Style(st *Style) {
	s.s = st
}
