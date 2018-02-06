package gumi

type Axis uint8

const (
	AxisVertical   Axis = 1 << iota
	AxisHorizontal Axis = 1 << iota
	AxisAll = AxisVertical | AxisHorizontal
)
