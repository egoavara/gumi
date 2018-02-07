package gumi

type _MaterialAnimation struct {

} 

func (s _MaterialAnimation) Toggle(t float64) float64 {
	return Animation.Functions.Default(t)
}

func (s _MaterialAnimation) Button(t float64) float64 {
	return Animation.Functions.Default(t)
}
func (s _MaterialAnimation) Progress(t float64) float64 {
	return Animation.Functions.Quad.Easing(t)
}

func (s _MaterialAnimation) Radio(t float64) float64 {
	return Animation.Functions.Default(t)
}
func (s _MaterialAnimation) DropboxStretch(t float64) float64 {
	return Animation.Functions.Quad.EasingIn(t)
}
func (s _MaterialAnimation) DropboxScrool(t float64) float64 {
	return Animation.Functions.Linear(t)
}