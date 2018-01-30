package gumi

func LinkingFrom(lks ...GUMI) (from GUMI) {
	r := lks[0].(GUMI)
	p := r
	for _, v := range lks {
		p.Breed([]GUMI{v})
		p = v
	}
	return r
}
func LinkingTo(lks ...GUMI) (to GUMI) {
	r := lks[0].(GUMI)
	p := r.(GUMI)
	for _, v := range lks {
		p.Breed([]GUMI{v})
		p = v
	}
	return lks[len(lks)-1].(GUMI)
}