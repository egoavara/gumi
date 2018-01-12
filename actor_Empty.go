package gumi

type aEmpty struct {
	GUMILINK_EMPTY
}

func (s *aEmpty) size(drawing *Drawing, style *Style) Size {
	return AUTOSIZE
}
func (s *aEmpty) draw(drawing *Drawing, style *Style, frame Frame) {
}

//
func AEmpty() *aEmpty {
	return &aEmpty{}
}