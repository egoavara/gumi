package gumi

type GUMIElem interface {
	// Build using
	size(drawing *Drawing, style *Style) Size
	//
	draw(drawing *Drawing, style *Style, frame Frame)
	//drawCaching(drawing *Drawing, style *Style, frame Frame)
}

type GUMILink interface {
	GUMIElem
	GUMILinker
}
type GUMILinker interface {
	born(GUMILinker)
	Parent() GUMILinker
	Link(elem ...GUMILinker)
	Childrun() []GUMILinker
}

type GUMILINK_MULTIPLE struct {
	parent GUMILinker
	child  []GUMILinker
}

func (s *GUMILINK_MULTIPLE) born(parent GUMILinker) {
	s.parent = parent
}
func (s *GUMILINK_MULTIPLE) Parent() GUMILinker {
	return s.parent
}
func (s *GUMILINK_MULTIPLE) Link(elem ...GUMILinker) {
	for _, e := range elem {
		e.born(s)
	}
	s.child = elem
}
func (s *GUMILINK_MULTIPLE) Childrun() []GUMILinker {
	temp := make([]GUMILinker, len(s.child))
	copy(temp, s.child)
	return temp
}

type GUMILINK_SINGLE struct {
	parent GUMILinker
	child  GUMILinker
}

func (s *GUMILINK_SINGLE) born(parent GUMILinker) {
	s.parent = parent
}
func (s *GUMILINK_SINGLE) Parent() GUMILinker {
	return s.parent
}
func (s *GUMILINK_SINGLE) Link(elem ...GUMILinker) {
	elem[0].born(s)
	s.child = elem[0]
}
func (s *GUMILINK_SINGLE) Childrun() []GUMILinker {
	return []GUMILinker{s.child}
}

type GUMILINK_EMPTY struct {
	parent GUMILinker
}

func (s *GUMILINK_EMPTY) born(parent GUMILinker) {
	s.parent = parent
}
func (s *GUMILINK_EMPTY) Parent() GUMILinker {
	return s.parent
}
func (s *GUMILINK_EMPTY) Link(elem ...GUMILinker) {
}
func (s *GUMILINK_EMPTY) Childrun() []GUMILinker {
	return nil
}
