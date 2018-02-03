package gumi

type empty struct{}

type synchronizer struct {
	waitcount int
	ch chan empty
}

func newSynchronizer(count int) *synchronizer {
	return &synchronizer{
		ch:make(chan empty, 4),
		waitcount:count,
	}
}
func (s *synchronizer ) Close() {
	for i := 0; i < s.waitcount; i++ {
		<- s.ch
	}
	close(s.ch)
}
func (s *synchronizer ) FinishWork() {
	s.ch <- empty{}
}