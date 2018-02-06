package gumi

type AnimationStudio struct {
	ani []AnimationActor
}

func NewAnimationStudio(anicount int) *AnimationStudio {
	temp := &AnimationStudio{
		ani: make([]AnimationActor, anicount),
	}
	return temp
}
func (s *AnimationStudio) Get(idx int) AnimationActor {
	return s.ani[idx]
}
func (s *AnimationStudio) Set(idx int, actor AnimationActor) AnimationActor {
	s.ani[idx] = actor
	return actor
}
func (s *AnimationStudio) Reset() {
	for _, v := range s.ani {
		v.Reset()
	}
}
func (s *AnimationStudio) Animate(info *Information) {
	for _, v := range s.ani {
		v.Animate(info)
	}
}

type AnimationActor interface {
	Reset()
	Animate(info *Information)
}

type AnimationPercent struct {
	From, Current, Delta, To float64
	Fn                       AnimationFunction
}
func (s *AnimationPercent) Reset() {
	s.From = 0
	s.Current = 0
	s.To = 0
}
func (s *AnimationPercent) Function(Fn AnimationFunction) {
	s.Fn = Fn
}
func (s *AnimationPercent) Request(to float64) {
	if s.To != to {
		s.From = s.Current
		s.To = to
	}
}
func (s *AnimationPercent) Animate(info *Information) {
	var dt = float64(info.Dt)
	if s.Current == s.To {
		return
	}
	if s.To > s.From {
		s.Current += s.Delta * dt / 1000
		if s.Current > s.To {
			s.Current = s.To
		}
	}
	if s.To < s.From {
		s.Current -= s.Delta * dt / 1000
		if s.Current < s.To {
			s.Current = s.To
		}
	}
}

type AnimationSwitch struct {
	Switch            bool
	Current, Interval float64
}
func (s *AnimationSwitch) Reset() {
	s.Switch = false
	s.Current = 0
}
func (s *AnimationSwitch) Animate(info *Information) {
	var dt = float64(info.Dt)
	s.Current += dt
	s.Switch = (int(s.Current) / int(s.Interval)) % 2 == 1
}

type AnimationReaching struct {
	on                 bool
	Current, Delta, To float64
	Fn                 AnimationFunction
}
func (s *AnimationReaching) Start() {
	s.on = true
}
func (s *AnimationReaching) Pause() {
	s.on = false
}
func (s *AnimationReaching) Stop() {
	s.Current = 0
	s.on = false
}
func (s *AnimationReaching) Reset() {
	s.on = false
	s.Current = 0
}
func (s *AnimationReaching) Function(Fn AnimationFunction) {
	s.Fn = Fn
}
func (s *AnimationReaching) Animate(info *Information) {
	if !s.on {
		return
	}
	if s.Current == s.To {
		return
	}
	var dt = float64(info.Dt)
	s.Current += s.To * s.Delta * dt / 1000
	if s.Current > s.To {
		s.Current = s.To
	}
}
