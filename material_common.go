package gumi

import "image/color"

type mtColorSingle struct {
	mcl1 MaterialColor
}

func (s *mtColorSingle) GetMaterialColor() MaterialColor {
	return s.mcl1
}
func (s *mtColorSingle) SetMaterialColor(mc MaterialColor)  {
	s.mcl1 = mc
}

type mtColorFromTo struct {
	mcl1 MaterialColor
	mcl2 MaterialColor
}

func (s *mtColorFromTo) GetFromMaterialColor() MaterialColor {
	return s.mcl1
}
func (s *mtColorFromTo) SetFromMaterialColor(mc MaterialColor)  {
	s.mcl1 = mc
}

func (s *mtColorFromTo) GetToMaterialColor() MaterialColor {
	return s.mcl2
}
func (s *mtColorFromTo) SetToMaterialColor(mc MaterialColor)  {
	s.mcl2 = mc
}


func phaseColor(c1, c2 color.Color, handle float64) color.Color {
	r1, g1, b1, a1 := c1.RGBA()
	r2, g2, b2, a2 := c2.RGBA()
	var h2 = handle
	var h1 = 1 - h2
	var res color.RGBA64
	res.R = uint16(float64(r1) * h1 + float64(r2) * h2)
	res.G = uint16(float64(g1) * h1 + float64(g2) * h2)
	res.B = uint16(float64(b1) * h1 + float64(b2) * h2)
	res.A = uint16(float64(a1) * h1 + float64(a2) * h2)
	return res

}
func phasePos(width float64, handle float64) float64 {
	return width*handle
}
