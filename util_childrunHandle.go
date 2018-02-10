package gumi

import "github.com/iamGreedy/gumi/gumre"

func saveGUMIChildrun(dst *[]GUMI, mode gumre.Mode, index gumre.Index, src ...GUMI) (writen int) {
	var ln = len(*dst)
	var idx int
	switch mode {
	default:
		fallthrough
	case gumre.REPLACE:
		for i, v := range src {
			idx = (index + gumre.Index(i)).Indexize(ln)
			if idx == gumre.IndexNotExist {
				break
			}
			(*dst)[idx] = v
			writen++
		}
	case gumre.PUSHBACKWARD:
		idx = (index + gumre.Index(1)).Indexize(ln)
		if idx == gumre.IndexNotExist{
			(*dst) = append(
				(*dst),
				src...,
			)
		}else {
			dstforw := (*dst)[:idx]
			dstback := (*dst)[idx:]
			(*dst) = append(
				dstforw,
				append(
					src,
					dstback...,
				)...,
			)
		}
	case gumre.PUSHONWARD:
		idx = index.Indexize(ln)
		dstforw := (*dst)[:idx]
		dstback := (*dst)[idx:]
		(*dst) = append(
			dstforw,
			append(
				src,
				dstback...,
			)...,
		)
		writen = len(src)
	}
	return writen
}
func loadGUMIChildrun (dst []GUMI, index gumre.Index, count int) (res []GUMI) {
	var ln = len(dst)
	var idx = index.Indexize(ln)
	var sz int
	if idx == gumre.IndexNotExist{
		return
	}
	if idx + count <= ln{
		sz = count
	}else {
		sz = ln - idx
	}
	res = make([]GUMI, sz)
	for i := 0; i < sz; i ++ {
		res[i] = dst[idx + i]
	}
	return res
}