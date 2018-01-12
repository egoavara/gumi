package gumi

const (
	EL_Length     ExtendLengthType = iota
	EL_Proportion ExtendLengthType = iota
)

type ExtendLengthType uint8
type ExtendLength struct {
	elt ExtendLengthType
	len Length
}

func (s *ExtendLength) Type() ExtendLengthType {
	return s.elt
}
func (s *ExtendLength) IsProportion() bool {
	return s.elt == EL_Proportion
}
func (s *ExtendLength) Proportion() int {
	return int(s.len.Min)
}
func (s *ExtendLength) IsLength() bool {
	return s.elt == EL_Length
}
func (s *ExtendLength) Length() Length {
	return s.len
}

//
func ELProportion(percentage uint) ExtendLength {
	return ExtendLength{
		elt: EL_Proportion,
		len: FixedLength(uint16(percentage)),
	}
}
func ELLength(l Length) ExtendLength {
	return ExtendLength{
		elt: EL_Length,
		len: l,
	}
}
