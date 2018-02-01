package gumi

import "fmt"

func LinkingFrom(lks ...GUMI) (from GUMI) {
	r := lks[0].(GUMI)
	p := r
	for _, v := range lks {
		if p != r{
			v.Born(p)
		}
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
func GUMIParents(e GUMI) string {
	if e.Parent() == nil{
		return fmt.Sprint(e)
	}
	return GUMIParents(e.Parent()) + fmt.Sprint(".", e)
}