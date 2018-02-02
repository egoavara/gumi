package gumi

import (
	"fmt"
	"github.com/fogleman/gg"
	"image"
	"math"
)

const (
	mtEditMinWidth             = 80
	mtEditMinHeight            = 20
	mtEditAnimMillis           = 400
	mtEditDeleteMaxMillisReach = 1000
	mtEditDeleteMaxPerSecond   = 100
)

type MTEdit struct {
	VoidStructure
	BoundStore
	StyleStore
	//
	mtColorSingle
	//
	align Align
	text  string
	//
	onChange MTEditChange
	//
	deleteContinue      int64
	deleteAccum         float64
	deleteMode          bool
	textCursorHandle    int64
	textCursorShow      bool
	cursorEnter, active bool
}
type MTEditChange func(self *MTEdit, text string)

func (s *MTEdit) String() string {
	return fmt.Sprintf("%s(text:%s)", "MTEdit", s.text)
}
func (s *MTEdit) draw(frame *image.RGBA) {
	var ctx = GGContextRGBASub(frame, s.bound)
	var w, h = float64(ctx.Width()), float64(ctx.Height())
	var radius = h / 2
	s.style.useContext(ctx)
	s.style.releaseContext(ctx)
	// string position make
	var drawtext = s.text
	if s.active && s.textCursorShow {
		drawtext += "_"
	}
	var expectw, expecth = ctx.MeasureString(drawtext)
	var stringposX, stringposY float64
	var vert, hori = ParseAlign(s.align)

	switch vert {
	case Align_BOTTOM:
		stringposY = h
	case Align_VCENTER:
		stringposY = h/2 + expecth/2
	case Align_TOP:
		stringposY = 0 + expecth
	}
	switch hori {
	case Align_RIGHT:
		stringposX = w - radius - expectw
	case Align_HCENTER:
		stringposX = w/2 - expectw/2
	case Align_LEFT:
		stringposX = radius
	}
	//
	ctx.SetColor(s.style.Material.PalletteColor(s.mtColorSingle.mcl1)[0])
	ctx.DrawArc(radius, radius, radius, gg.Radians(90), gg.Radians(270))
	ctx.DrawRectangle(radius, 0, w-radius*2, h)
	ctx.DrawArc(w-radius, radius, radius, gg.Radians(-90), gg.Radians(90))
	ctx.Fill()
	//
	ctx.SetColor(s.style.Material.PalletteColor(s.mtColorSingle.mcl1)[1])
	if s.active {
		ctx.DrawArc(radius, radius, radius, gg.Radians(90), gg.Radians(270))
		ctx.DrawLine(radius, 0, w-radius, 0)
		ctx.DrawLine(radius, h, w-radius, h)
		ctx.Stroke()
		ctx.DrawArc(w-radius, radius, radius, gg.Radians(-90), gg.Radians(90))
		ctx.Stroke()
	}
	ctx.DrawString(drawtext, stringposX, stringposY)
	ctx.Stroke()
}
func (s *MTEdit) size() Size {
	return Size{
		Vertical:   MinLength(mtEditMinHeight),
		Horizontal: MinLength(mtEditMinWidth),
	}
}
func (s *MTEdit) rect(r image.Rectangle) {
	s.bound = r
}
func (s *MTEdit) update(info *Information, style *Style) {
	s.style = style
	//
	if s.active {
		s.textCursorHandle += info.Dt
		if s.textCursorHandle/mtEditAnimMillis > 0 {
			s.textCursorShow = !s.textCursorShow
			s.textCursorHandle = s.textCursorHandle % mtEditAnimMillis
		}
		if s.deleteMode {
			if s.deleteContinue != mtEditDeleteMaxMillisReach {
				s.deleteContinue += info.Dt
				if s.deleteContinue > mtEditDeleteMaxMillisReach {
					s.deleteContinue = mtEditDeleteMaxMillisReach
				}
			}
			s.deleteAccum += (mtEditDeleteMaxPerSecond * float64(s.deleteContinue) / mtEditDeleteMaxMillisReach) * (float64(info.Dt) / 1000)
			delaccumfloor := math.Floor(s.deleteAccum)
			if delaccumfloor >= 1 {
				s.deleteAccum -= delaccumfloor
				s.text = StringDeleteBack(s.text, int(delaccumfloor))
				if s.onChange != nil {
					s.onChange(s, s.text)
				}
			}
		}
	}
}
func StringDeleteBack(str string, count int) string {
	temp := []rune(str)
	templen := len(temp)
	if count > templen {
		count = templen
	}
	return string(temp[:templen-count])

}
func (s *MTEdit) Occur(event Event) {
	switch ev := event.(type) {
	case EventKeyPress:
		switch ev.Key {
		case KEY_BACKSPACE:
			if s.active {
				if !s.deleteMode {
					s.text = StringDeleteBack(s.text, 1)
					if s.onChange != nil {
						s.onChange(s, s.text)
					}
				}
				s.deleteMode = true
			}
		}
	case EventKeyRelease:
		switch ev.Key {
		case KEY_BACKSPACE:
			if s.active {
				s.deleteMode = false
				s.deleteContinue = 0
				s.deleteAccum = 0
			}

		case KEY_MOUSE1:
			if s.cursorEnter {
				s.active = true
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
	case EventRune:
		if s.active {
			s.text += string(ev.Rune)
			if s.onChange != nil {
				s.onChange(s, s.text)
			}
		}
	}

}
//
func MTEdit0(mcl MaterialColor, str string, align Align) *MTEdit {
	tem := &MTEdit{
		text:  str,
		align: align,
	}
	tem.SetMaterialColor(mcl)
	return tem
}
func MTEdit1(str string, align Align) *MTEdit {
	return &MTEdit{
		text:  str,
		align: align,
	}
}
func MTEdit2(str string) *MTEdit {
	return &MTEdit{
		text:  str,
		align: Align_LEFT | Align_VCENTER,
	}
}
func MTEdit3() *MTEdit {
	return &MTEdit{
		text:  "",
		align: Align_LEFT | Align_VCENTER,
	}
}

func (s *MTEdit) Set(str string) {
	s.SetText(str)
}
func (s *MTEdit) Get() string {
	return s.GetText()
}
func (s *MTEdit) SetText(str string) {
	s.text = str
	if s.onChange != nil {
		s.onChange(s, s.text)
	}
}
func (s *MTEdit) GetText() string {
	return s.text
}
func (s *MTEdit) SetAlign(align Align) {
	s.align = align
}
func (s *MTEdit) GetAlign() Align {
	return s.align
}
func (s *MTEdit) OnChange(callback MTEditChange) {
	s.onChange = callback
}
func (s *MTEdit) ReferChange() MTEditChange {
	return s.onChange
}
