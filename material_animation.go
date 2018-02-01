package gumi

var Material _MaterialAnimation
type _MaterialAnimation struct {
}

func (s _MaterialAnimation) Toggle(t float64) float64 {
	return Animation.Default(t)
}

func (s _MaterialAnimation) Button(t float64) float64 {
	return Animation.Default(t)
}
func (s _MaterialAnimation) Progress(t float64) float64 {
	return Animation.Linear(t)
}