package gumi

func LinkingFrom(lks ...GUMILinker) (from GUMILink) {
	r := lks[0].(GUMILink)
	p := r.(GUMILinker)
	for _, v := range lks {
		p.Link(v)
		p = v
	}
	return r
}
func LinkingTo(lks ...GUMILinker) (to GUMILink) {
	r := lks[0].(GUMILink)
	p := r.(GUMILinker)
	for _, v := range lks {
		p.Link(v)
		p = v
	}
	return lks[len(lks)-1].(GUMILink)
}

func DrawListing(fns ...[]DrawFunc) []DrawFunc {
	var temp []DrawFunc
	for _, fn := range fns {
		temp = append(temp, fn...)
	}
	return temp
}
