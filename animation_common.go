package gumi

var Animation _Animation
type _Animation struct {
	Functions _AnimationFunctions
}
type _AnimationFunctions struct {
	//
	Quad  AnimFuncsQuad
	Cubic AnimFuncsCubic
	Quart AnimFuncsQuart
	Quint AnimFuncsQuint
	//
	Material _MaterialAnimation
}

func (_Animation) DeltaByMillis(millis float64) (delta float64) {
	return 1000 / millis
}