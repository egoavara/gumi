package gumi

import (
	"fmt"
	"image"
)

const (
	MTDropboxMinWidth             = 80
	MTDropboxMinHeight            = 20
	MTDropboxAnimMillis           = 400
	MTDropboxDeleteMaxMillisReach = 1000
	MTDropboxDeleteMaxPerSecond   = 100
)

type MTDropbox struct {
	VoidStructure
	BoundStore
	StyleStore
	//
	mtColorSingle
	//
	align    Align
	elems    []GUMI
	selected int
	inactive bool
	//
	onChange MTDropboxChange
	//
	cursorEnter, active bool
}
type MTDropboxChange func(self *MTDropbox, selected GUMI)

func (s *MTDropbox) String() string {
	return fmt.Sprintf("%s(select:%v)", "MTDropbox", s.elems[s.selected])
}
func (s *MTDropbox) draw(frame *image.RGBA) {
	var ctx = GGContextRGBASub(frame, s.bound)
	var w, h = float64(ctx.Width()), float64(ctx.Height())
	var radius = h / 2
	s.style.useContext(ctx)
	defer s.style.releaseContext(ctx)
	// string position make
}
func (s *MTDropbox) size() Size {
	return Size{
		Vertical:   MinLength(MTDropboxMinHeight),
		Horizontal: MinLength(MTDropboxMinWidth),
	}
}
func (s *MTDropbox) rect(r image.Rectangle) {
	s.bound = r
}
func (s *MTDropbox) update(info *Information, style *Style) {
	s.style = style
	//
	if s.active {
	}
}
func (s *MTDropbox) Occur(event Event) {
	switch ev := event.(type) {
	case EventKeyRelease:
		switch ev.Key {
		case KEY_MOUSE1:
			if s.cursorEnter {
				if !s.inactive {
					s.active = true
				}
			} else {
				s.active = false
			}
		}

	case EventCursor:
		x := int(ev.X)
		y := int(ev.Y)
		if (s.bound.Min.X <= x && x < s.bound.Max.X) && (s.bound.Min.Y <= y && y < s.bound.Max.Y) {
			s.cursorEnter = true
		} else {
			s.cursorEnter = false
		}
	}

}
func (s *MTDropbox) elemsVLength() Length{
	l := AUTOLENGTH
	for _, v := range s.elems{
		sz := v.size()
		l.Min += sz.Vertical.Min
	}
	return l
}

//
func MTDropbox0(mcl MaterialColor, str string, align Align) *MTDropbox {
	tem := &MTDropbox{
		text:  str,
		align: align,
	}
	tem.SetMaterialColor(mcl)
	return tem
}
func MTDropbox1(str string, align Align) *MTDropbox {
	return &MTDropbox{
		text:  str,
		align: align,
	}
}
func MTDropbox2(str string) *MTDropbox {
	return &MTDropbox{
		text:  str,
		align: Align_LEFT | Align_VCENTER,
	}
}
func MTDropbox3() *MTDropbox {
	return &MTDropbox{
		text:  "",
		align: Align_LEFT | Align_VCENTER,
	}
}

func (s *MTDropbox) OnChange(callback MTDropboxChange) {
	s.onChange = callback
}
func (s *MTDropbox) ReferChange() MTDropboxChange {
	return s.onChange
}
