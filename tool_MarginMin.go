package gumi

import "github.com/iamGreedy/gumi/gumre"

func (_Tool) MarginMinRegular(min uint16, gumi GUMI) GUMI {
	temp := NMargin0(gumre.RegularBlank(gumre.MinLength(min)))
	temp2 := LinkingFrom(temp, gumi)
	return temp2
}
func (_Tool) MarginMinSymmetry(minHorizontal, minVertical uint16, gumi GUMI) GUMI {
	temp := NMargin0(gumre.SymmetryBlank(gumre.MinLength(minHorizontal), gumre.MinLength(minVertical)))
	return LinkingFrom(temp, gumi)
}
