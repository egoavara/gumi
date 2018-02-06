package gumi

func (_Tool) MarginMinRegular(min uint16, gumi GUMI) GUMI {
	temp := NMargin0(RegularBlank(MinLength(min)))
	temp2 := LinkingFrom(temp, gumi)
	return temp2
}
func (_Tool) MarginMinSymmetry(minHorizontal, minVertical uint16, gumi GUMI) GUMI {
	temp := NMargin0(SymmetryBlank(MinLength(minHorizontal), MinLength(minVertical)))
	return LinkingFrom(temp, gumi)
}
