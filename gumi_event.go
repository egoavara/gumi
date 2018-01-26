package gumi

const (
	KEY    EventKind = iota
	CURSOR EventKind = iota
	RUNE   EventKind = iota
)

type EventKind uint8
type Event interface {
	Kind() EventKind
}

type EventCursor struct {
	X, Y uint16
}
func (EventCursor) Kind() EventKind {
	return CURSOR
}

type EventKey struct {
	Key uint8
}
func (EventKey) Kind() EventKind {
	return KEY
}

type EventRune struct {
	Rune rune
}
func (EventRune) Kind() EventKind {
	return RUNE
}
