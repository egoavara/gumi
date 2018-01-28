package gumi

import "math"

type Blank struct {
	L, B, R, T Length
}

func SymmetryBlank(horizontal, vertical Length) Blank {
	return Blank{
		L:horizontal,
		R:horizontal,
		B:vertical,
		T:vertical,
	}
}
func RegularBlank(regular Length) Blank {
	return Blank{
		L:regular,
		R:regular,
		B:regular,
		T:regular,
	}
}
type Size struct {
	Vertical   Length
	Horizontal Length
}

type Length struct {
	Min, Max uint16
}
var AUTOLENGTH = Length{
	Min:0,
	Max:math.MaxUint16,
}

func MinLength(min uint16) Length {
	return Length{
		Min:min,
		Max:math.MaxUint16,
	}
}
func MaxLength(max uint16) Length {
	return Length{
		Min:0,
		Max:max,
	}
}
func FixLength(fix uint16) Length {
	return Length{
		Min:fix,
		Max:fix,
	}
}