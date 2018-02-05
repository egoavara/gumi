package gumi

import (
	"fmt"
	"github.com/fogleman/gg"
	"image"
)

const (
	MTDropboxMinWidth             = 80
	MTDropboxScroolWidth          = 10
	MTDropboxMinHeight            = 20
	MTDropboxAnimStretchPerSecond = 300
)

type MTDropbox struct {
	VoidStructure
	BoundStore
	StyleStore
	//
	mtColorSingle
	//
	elems    []mtDropboxElem
	selected int
	inactive bool
	//
	onChange MTDropboxChange
	//
	cursorEnter, active bool
	box, boxTo          uint16
	scrool, scroolTo    uint16
}
type mtDropboxElem struct {
	elem          string
	width, height int
}

type MTDropboxChange func(self *MTDropbox, selected string)

func (s *MTDropbox) String() string {
	return fmt.Sprintf("%s(select:%s)", "MTDropbox", s.elems[s.selected])
}
func (s *MTDropbox) draw(frame *image.RGBA) {
	bd := s.bound
	//
	var ctx = GGContextRGBASub(frame, s.bound)
	s.style.useContext(ctx)
	defer s.style.releaseContext(ctx)
	//
	ctx.SetColor(s.style.Material.PalletteColor(s.mtColorSingle.mcl1)[0])
	var w, h = float64(ctx.Width()), float64(ctx.Height())
	var radius = float64(s.bound.Dy()) / 2
	//
	if s.active{
		ctx.DrawArc(radius, radius, radius, gg.Radians(180), gg.Radians(270))
		ctx.DrawRectangle(radius, 0, w - radius * 2, h)
		ctx.DrawArc(w-radius, radius, radius, gg.Radians(270), gg.Radians(360))
	}else{
		var elemsum = 0
		var elemcut = 0
		for _, v := range s.elems {
			elemsum += v.height
		}
		bd.Max.Y += elemsum
		if bd.Max.Y > frame.Rect.Max.Y {
			elemcut = bd.Max.Y - frame.Rect.Max.Y
			bd.Max.Y = frame.Rect.Max.Y
		}
	}


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
	if s.style != style {
		style.Default.Font.Use()
		defer style.Default.Font.Release()
		for _, v := range s.elems {
			v.width, v.height = style.Default.Font.CalculateSize(v.elem)
		}
	}
	s.style = style
	//
	delta := float64(MTDropboxAnimStretchPerSecond) / 1000 * float64(info.Dt)
	if s.box != s.boxTo {
		fbox := float64(s.box)
		if s.box > s.boxTo {
			fbox = fbox - delta
			if fbox < float64(s.boxTo) {
				fbox = float64(s.boxTo)
			}
		} else {
			fbox = fbox + delta
			if fbox > float64(s.boxTo) {
				fbox = float64(s.boxTo)
			}
		}
		s.box = uint16(fbox)
	}
	if s.scrool != s.scroolTo {
		fscrool := float64(s.scrool)
		if s.scrool > s.scroolTo {
			fscrool = fscrool - delta
			if fscrool < float64(s.scroolTo) {
				fscrool = float64(s.scroolTo)
			}
		} else {
			fscrool = fscrool + delta
			if fscrool > float64(s.scroolTo) {
				fscrool = float64(s.scroolTo)
			}
		}
		s.scrool = uint16(fscrool)
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
func (s *MTDropbox) IndexElem(idx int) string {
	if idx < 0 {
		return ""
	}else if idx >= len(s.elems){
		return ""
	}
	return s.elems[idx].elem
}
func (s *MTDropbox) ExistElem(idx int) bool{
	if idx < 0 {
		return ""
	}else if idx >= len(s.elems){
		return ""
	}
	return s.elems[idx].elem
}
//
func (s *MTDropbox) OnChange(callback MTDropboxChange) {
	s.onChange = callback
}
func (s *MTDropbox) ReferChange() MTDropboxChange {
	return s.onChange
}
