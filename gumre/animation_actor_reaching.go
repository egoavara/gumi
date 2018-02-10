package gumre



type Reaching struct {
	off                       bool
	Current, Delta, To, Range float64
	Fn                        AnimationFunction
}

func (s *Reaching) Start() {
	s.off = false
}
func (s *Reaching) Pause() {
	s.off = true
}
func (s *Reaching) Stop() {
	s.Current = 0
	s.off = true
}
func (s *Reaching) Reset() {
	s.off = false
	s.Current = 0
}
func (s *Reaching) Function(Fn AnimationFunction) {
	s.Fn = Fn
}
func (s *Reaching) Animate(delta float64) {
	if s.off {
		return
	}
	if s.Current == s.To {
		return
	}
	if s.Current < s.To {
		s.Current += s.Delta * delta / 1000
		if s.Current > s.To {
			s.Current = s.To
		}
	} else if s.Current > s.To {
		s.Current -= s.Delta * delta / 1000
		if s.Current < s.To {
			s.Current = s.To
		}
	}

}
func (s *Reaching) Value() float64 {
	if s.Range == 0 || s.Current == 0 {
		return 0
	}
	return s.Range * s.Fn(s.Current/s.Range)
}
func (s *Reaching) Percent() float64 {
	if s.Range == 0 || s.Current == 0 {
		return 0
	}
	return s.Value() / s.Range
}