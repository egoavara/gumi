package gumi

import "fmt"

func LinkingFrom(lks ...GUMI) (from GUMI) {
	for i := 0 ; i < len(lks); i ++{
		if i != 0{
			lks[i].Born(lks[i - 1])
		}
		if i < len(lks) - 1{
			lks[i].Breed([]GUMI{lks[i + 1]})
		}
	}
	return lks[0]
}
//func LinkingTo(lks ...GUMI) (to GUMI) {
//	r := lks[0].(GUMI)
//	p := r.(GUMI)
//	for _, v := range lks {
//		p.Breed([]GUMI{v})
//		p = v
//	}
//	return lks[len(lks)-1].(GUMI)
//}
func GUMIParents(e GUMI) string {
	if e.Parent() == nil{
		return fmt.Sprint(e)
	}
	return GUMIParents(e.Parent()) + fmt.Sprint(".", e)
}