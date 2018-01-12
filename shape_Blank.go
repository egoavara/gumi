package gumi

import (
	"fmt"
	"math"
)

var NOBLANK = Blank{
	Length{0, 0},
	Length{0, 0},
	Length{0, 0},
	Length{0, 0},
}
var MAXBLANK = Blank{
	FixedLength(math.MaxUint16),
	FixedLength(math.MaxUint16),
	FixedLength(math.MaxUint16),
	FixedLength(math.MaxUint16),
}

type Blank struct {
	L, B, R, T Length
}

func (s *Blank) VMin() uint16 {
	return s.T.Min + s.B.Min
}
func (s *Blank) HMin() uint16 {
	return s.L.Min + s.R.Min
}
func (s Blank) String() string {
	return fmt.Sprintf(
		"Blank(L : %d, R : %d, T : %d, B : %d)", s.L, s.R, s.T, s.B,
	)
}

func SymmetryBlank(h, v Length) Blank {
	return MakeBlank(h, v, h, v)
}
func RegularBlank(r Length) Blank {
	return MakeBlank(r, r, r, r)
}

func MakeBlank(l, b, r, t Length) Blank {
	return Blank{
		l, b, r, t,
	}
}
