package gumi

import "math"

type Blank struct {
	L, B, R, T Length
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