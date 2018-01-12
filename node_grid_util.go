package gumi

import "sync"

type gridNode struct {
	wcount, hcount   int
	wlength, hlength []ExtendLength
	//
	wcache, hcache   int
	wlcache, hlcache []int
	mtx *sync.RWMutex
}

func GridNode(wc, hc int) gridNode {
	temp := gridNode{
		wcount:  wc,
		hcount:  hc,
		wlength: make([]ExtendLength, wc),
		hlength: make([]ExtendLength, hc),
		//
		wlcache: make([]int, wc),
		hlcache: make([]int, hc),
		//
		wcache: -1,
		hcache: -1,
		mtx: new(sync.RWMutex),
	}
	for i := range temp.wlength {
		temp.wlength[i] = ELProportion(1)
	}
	for i := range temp.hlength {
		temp.hlength[i] = ELProportion(1)
	}
	return temp
}
func (s *gridNode) Size() (w, h int) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return s.wcount, s.hcount
}
func (s *gridNode) Width(x int) ExtendLength {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return s.wlength[x]
}
func (s *gridNode) Height(y int) ExtendLength {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return s.hlength[y]
}
func (s *gridNode) SetWidth(x int, l ExtendLength) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.wlength[x] = l
}
func (s *gridNode) SetHeight(y int, l ExtendLength) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.hlength[y] = l
}
func (s *gridNode) StoreCache(w, h int) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.wcache, s.hcache = w, h
	distributeGN(w, s.wcount, s.wlength, s.wlcache)
	distributeGN(h, s.hcount, s.hlength, s.hlcache)

	return
}
func distributeGN(length, cnt int, lens []ExtendLength, cch []int) {
	var normalcount = 0
	var normallength = 0
	var addall = 0
	for x := 0; x < cnt; x++ {
		switch lens[x].Type() {
		case EL_Length:
			temp := lens[x].Length()
			cch[x] = int(temp.Min)
			if addall > length {
				cch[x] = 0
			} else if addall+int(temp.Min) > length {
				cch[x] = length - addall
			} else {
				cch[x] = int(temp.Min)
			}
			addall += int(temp.Min)
		case EL_Proportion:
			normalcount += lens[x].Proportion()
		}
	}
	if addall > length {
		normallength = 0
	} else {
		if normalcount == 0 {
			normallength = 0
		} else {
			normallength = (length - addall) / normalcount
		}

	}
	for i, v := range lens {
		if v.IsProportion() {

		}
		switch v.Type() {
		case EL_Length:
		case EL_Proportion:
			cch[i] = normallength * v.Proportion()
		default:
			cch[i] = 0
		}
	}
	var sum = 0
	for i, v := range cch{
		sum += v
		cch[i] = sum
	}
}
func (s *gridNode) LoadCache(w, h int) (ok bool, wl, hl []int) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	if w == s.wcache && h == s.hcache {
		wl = make([]int, s.wcount)
		hl = make([]int, s.hcount)
		for i, v := range s.wlcache {
			wl[i] = v
		}
		for i, v := range s.hlcache {
			hl[i] = v
		}
		return true, wl, hl
	} else {
		return false, nil, nil
	}
}
func (s *gridNode) Indexize(x, y int) (int){
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return x + y * s.wcount
}