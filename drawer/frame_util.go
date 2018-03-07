package drawer

import "sync/atomic"

type ID uint32

type IDSpace struct {
	top uint32
}
func (s *IDSpace) New() ID {
	return ID(atomic.AddUint32(&s.top, 1))
}
