package gumi

func Proportion(have, a, b uint16) (pa, pb uint16) {
	ab := float64(a + b)
	fh := float64(have)
	pa = uint16(fh / ab * float64(a))
	pb = uint16(fh / ab * float64(b))
	return
}
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

