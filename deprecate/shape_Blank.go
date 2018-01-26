package deprecate

import (
	"fmt"
	"math"
)

var NOBLANK = Blank{
	Leng{0, 0},
	Leng{0, 0},
	Leng{0, 0},
	Leng{0, 0},
}
var MAXBLANK = Blank{
	FixedLength(math.MaxUint16),
	FixedLength(math.MaxUint16),
	FixedLength(math.MaxUint16),
	FixedLength(math.MaxUint16),
}

type Blank struct {
	L, B, R, T Leng
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

func SymmetryBlank(h, v Leng) Blank {
	return MakeBlank(h, v, h, v)
}
func RegularBlank(r Leng) Blank {
	return MakeBlank(r, r, r, r)
}

func MakeBlank(l, b, r, t Leng) Blank {
	return Blank{
		l, b, r, t,
	}
}
