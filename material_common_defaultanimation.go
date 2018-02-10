package gumi

import "github.com/iamGreedy/gumi/gumre"

type _MaterialAnimation struct {

} 

func (s _MaterialAnimation) Toggle(t float64) float64 {
	return gumre.Animation.Functions.Default(t)
}

func (s _MaterialAnimation) Button(t float64) float64 {
	return gumre.Animation.Functions.Default(t)
}
func (s _MaterialAnimation) Progress(t float64) float64 {
	return gumre.Animation.Functions.Quad.Easing(t)
}

func (s _MaterialAnimation) Radio(t float64) float64 {
	return gumre.Animation.Functions.Default(t)
}
func (s _MaterialAnimation) DropboxStretch(t float64) float64 {
	return gumre.Animation.Functions.Quad.EasingIn(t)
}
func (s _MaterialAnimation) DropboxScrool(t float64) float64 {
	return gumre.Animation.Functions.Linear(t)
}