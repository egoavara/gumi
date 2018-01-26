package gumi

//import "image"
//
//type nVertical struct {
//	MultipleStructure
//	BoundStore
//}
//
//func (s *nVertical) draw(frame *image.RGBA) {
//	for _, v := range s.child{
//		v.draw(frame)
//	}
//}
//
//func (s *nVertical) size() Size {
//	return Size{
//		AUTOLVARIABLE,
//		AUTOLVARIABLE,
//	}
//}
//
//func helperLength(szs []Length, have int) (res []int) {
//	res = make([]int, len(szs))
//	var countratio = 0
//	var minSum = 0
//	var maxSum = 0
//	//
//	for _, v := range szs{
//		switch tp := v.(type) {
//		case LRatio:
//			countratio += int(tp.Ratio)
//		case LFixed:
//			minSum += int(tp.Fixed)
//			maxSum += int(tp.Fixed)
//		case LVariable:
//			minSum += int(tp.Min)
//			maxSum += int(tp.Max)
//		}
//	}
//	//
//	if maxSum <= have{
//		//모든게 남음
//		rat := have - maxSum
//		for i, v := range szs{
//			switch tp := v.(type) {
//			case LRatio:
//				res[i] = rat * int(tp.Ratio)
//			case LFixed:
//				res[i] = int(tp.Fixed)
//			case LVariable:
//				res[i] = int(tp.Max)
//			}
//		}
//	}else if minSum <= have{
//		// 최소크기는 만족하지만 최대크기는 만족안함
//		rat := have - minSum
//		for i, v := range szs{
//			switch tp := v.(type) {
//			case LRatio:
//				res[i] = rat * int(tp.Ratio)
//			case LFixed:
//				res[i] = int(tp.Fixed)
//			case LVariable:
//				res[i] = int(tp.Min)
//			}
//		}
//	}else {
//		//최소크기도 만족 못함
//		rat := have / minSum
//		for i, v := range szs{
//			switch tp := v.(type) {
//			case LRatio:
//				res[i] = 0
//			case LFixed:
//				res[i] = rat * int(tp.Fixed)
//			case LVariable:
//				res[i] = rat * int(tp.Min)
//			}
//		}
//	}
//	return
//}
//func (s *nVertical) rect(r image.Rectangle) {
//	s.bound = r
//	//
//	var temp = make([]Length, len(s.child))
//	for i, v := range s.child{
//		temp[i] = v.size().Vertical
//	}
//	res := helperLength(temp, r.Dy())
//	//
//	var startat = 0
//	for i, v := range s.child{
//		v.rect(image.Rect(
//			0,
//			startat,
//			r.Dx(),
//			startat + res[i],
//		))
//		startat += res[i]
//	}
//}
//
//func (s *nVertical) update(info *Information, style *Style) {
//	for _, v := range s.child{
//		v.update(info, style)
//	}
//}
//
//func (s *nVertical) Occur(event Event) {
//	for _, v := range s.child{
//		v.Occur(event)
//	}
//}

//func (s *nVertical) size(drawing *Drawing, style *Style) Size {
//	size := Size{
//		Horizontal: AUTOLENGTH,
//		Vertical:   AUTOLENGTH,
//	}
//	for _, v := range s.child {
//		temp := v.(GUMILink).size(drawing, style)
//		size.Vertical.Min += temp.Vertical.Min
//	}
//	return size
//}
//func (s *nVertical) draw(drawing *Drawing, style *Style, frame Frame) {
//	var start uint16 = 0
//	var snc = make(chan struct{})
//	var cnt = len(s.child) - 1
//	//
//	for i, v := range s.child {
//		v2 := v.(GUMILink)
//		temp := v2.size(drawing, style)
//		sf := frame.SubFrame(
//			image.Rect(0, int(start), int(temp.Horizontal.Max), int(start+temp.Vertical.Min)),
//		)
//		if !sf.Bounds().Empty() {
//			go func(a int) {
//				v2.draw(drawing, style, sf)
//				if cnt == a {
//					close(snc)
//				}
//			}(i)
//		}
//		start += temp.Vertical.Min
//	}
//	for range snc {
//	}
//}
//func NVertical(childrun ...GUMI) *nVertical {
//	s := &nVertical{}
//	s.Breed(childrun)
//	return s
//}
