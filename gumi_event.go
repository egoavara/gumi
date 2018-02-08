package gumi

const (
	KEYPRESS   EventKind = iota
	KEYRELEASE EventKind = iota
	CURSOR     EventKind = iota
	SCROLL     EventKind = iota
	RUNE       EventKind = iota
	RUNEEDIT       EventKind = iota
	// RESIZE     EventKind = iota
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

type EventScroll struct {
	X, Y int32
}

func (EventScroll) Kind() EventKind {
	return SCROLL
}

type EventKeyPress struct {
	Key uint8
}

func (EventKeyPress) Kind() EventKind {
	return KEYPRESS
}

type EventKeyRelease struct {
	Key uint8
}

func (EventKeyRelease) Kind() EventKind {
	return KEYRELEASE
}

type EventRuneComplete struct {
	Rune rune
}

func (EventRuneComplete) Kind() EventKind {
	return RUNE
}

type EventRuneEdit struct {
	Rune rune
}

func (EventRuneEdit) Kind() EventKind {
	return RUNEEDIT
}