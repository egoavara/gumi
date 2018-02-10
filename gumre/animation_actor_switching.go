package gumre


type Switching struct {
	Switch            bool
	Current, Interval float64
}

func (s *Switching) Reset() {
	s.Switch = false
	s.Current = 0
}
func (s *Switching) Animate(delta float64) {
	s.Current += delta
	s.Switch = (int(s.Current)/int(s.Interval))%2 == 1
}
