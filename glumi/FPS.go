package glumi

import "time"

type FPS struct {
	interval time.Duration
	ticker   *time.Ticker
}
func (s *FPS) Start()  {
	s.ticker = time.NewTicker(s.interval)
}
func (s *FPS) Wait() time.Time {
	return <- s.ticker.C
}
func (s *FPS) Stop()  {
	s.ticker.Stop()
	s.ticker = nil
}
func (s *FPS) SetInterval(interval time.Duration)  {
	s.interval = interval
}
func (s *FPS) GetInterval()(interval time.Duration)  {
	return s.interval
}