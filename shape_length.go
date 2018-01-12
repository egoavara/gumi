package gumi

import (
	"math"
)

//
const (
	LENGTHMIN = 0
	LENGTHMAX = math.MaxUint16
)

var AUTOLENGTH = Length{
	Min: 0,
	Max: math.MaxUint16,
}

func MinLength(l uint16) Length {
	return Length{l, math.MaxUint16}
}
func FixedLength(l uint16) Length {
	return Length{l, l}
}
func MaxLength(l uint16) Length {
	return Length{0, l}
}

type Length struct {
	Min uint16
	Max uint16
}
