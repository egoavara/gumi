package deprecate

import (
	"math"
)

//
const (
	LENGTHMIN = 0
	LENGTHMAX = math.MaxUint16
)

var AUTOLENGTH = Leng{
	Min: 0,
	Max: math.MaxUint16,
}

func MinLength(l uint16) Leng {
	return Leng{l, math.MaxUint16}
}
func FixedLength(l uint16) Leng {
	return Leng{l, l}
}
func MaxLength(l uint16) Leng {
	return Leng{0, l}
}

type Leng struct {
	Min uint16
	Max uint16
}
