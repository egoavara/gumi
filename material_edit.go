package gumi

import (
	"fmt"
	"github.com/fogleman/gg"
	"image"
)

const (
	mtEditMinWidth                    = 80
	mtEditMinHeight                   = 20
	mtEditAnimationTextCursorInterval = 400
	mtEditDeleteInterval              = 50
	mtEditDeleteThereshold            = 200
)
const (
	mtEditAnimationTextCursor = iota
	//
	mtEditAnimationLength = iota
)

type MTEdit struct {
	VoidStructure
	boundStore
	styleStore
	//
	mtColorSingle
	studio                *AnimationStudio
	textCursor            *AnimationSwitch
	deleteSum             int64
	deleteOn              bool
	deleteCount           int
	deleteTheresholdStack int64
	deleteRequestCount    uint
	ketCTRL               bool
	editingRune           rune
	editingNow            bool
	//
	align    Align
	text     string
	inactive bool
	//
	onChange            MTEditChange
	cursorEnter, active bool
}
type MTEditChange func(self *MTEdit, text string)

func (s *MTEdit) String() string {
	return fmt.Sprintf("%s(text:%s)", "MTEdit", s.text)
}
func (s *MTEdit) init() {
	s.studio = NewAnimationStudio(mtEditAnimationLength)
	s.textCursor = s.studio.Set(mtEditAnimationTextCursor, &AnimationSwitch{
		Interval: mtEditAnimationTextCursorInterval,
	}).(*AnimationSwitch)
}
func (s *MTEdit) draw(frame *image.RGBA) {
	var baseColor, mainColor = s.GetMaterialColor().Color()
	var ctx = GGContextRGBASub(frame, s.bound)
	var w, h = float64(ctx.Width()), float64(ctx.Height())
	var radius = h / 2
	//
	s.style.useContext(ctx)
	defer s.style.releaseContext(ctx)
	// string position make
	var drawtext = s.text
	if s.editingRune != 0 {
		drawtext += string(s.editingRune)
	} else {
		if s.active && s.textCursor.Switch {
			drawtext += "_"
		}
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
	ctx.SetColor(baseColor)
	ctx.DrawArc(radius, radius, radius, gg.Radians(90), gg.Radians(270))
	ctx.DrawRectangle(radius, 0, w-radius*2, h)
	ctx.DrawArc(w-radius, radius, radius, gg.Radians(-90), gg.Radians(90))
	ctx.Fill()
	//
	ctx.SetColor(mainColor)
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
	s.studio.Animate(info)
	if s.deleteRequestCount >0 {
		if !s.editingNow{
			if s.ketCTRL {
				s.text = StringControlBackSpace(s.text)
			} else {
				s.text = StringBackSpace(s.text, 1)
			}
		}else {
			if s.editingRune == 0{
				s.editingNow = false
			}
		}
		s.deleteRequestCount = 0
	}
	s.deleteSum += info.Dt
	s.deleteTheresholdStack += info.Dt
	temp := s.deleteSum / mtEditDeleteInterval
	if temp >= 1 {
		s.deleteSum -= temp * mtEditDeleteInterval
		if s.deleteOn && s.deleteTheresholdStack > mtEditDeleteThereshold {
			if !s.editingNow{
				for i := int64(0); i < temp; i++ {
					s.deleteCount += 1
					if s.ketCTRL {
						s.text = StringControlBackSpace(s.text)
					} else {
						s.text = StringBackSpace(s.text, 1)
					}
				}
			}else {
				if s.editingRune == 0{
					s.editingNow = false
				}
			}
		}
	}
}
func (s *MTEdit) deleteRequest(count uint) {
	s.deleteRequestCount += count
}
func (s *MTEdit) Occur(event Event) {
	switch ev := event.(type) {
	case EventKeyPress:
		switch ev.Key {
		case KEY_CONTROL:
			s.ketCTRL = true
		case KEY_BACKSPACE:
			if s.active {
				s.deleteOn = true
				s.deleteCount = 0
				s.deleteTheresholdStack = 0
			}
		}
	case EventKeyRelease:
		switch ev.Key {
		case KEY_CONTROL:
			s.ketCTRL = false
		case KEY_BACKSPACE:
			if s.active {
				if s.deleteCount == 0 {
					s.deleteRequest(1)
				}
				s.deleteOn = false
				s.deleteCount = 0
				s.deleteTheresholdStack = 0
			}
		case KEY_MOUSE1:
			if s.cursorEnter {
				if !s.inactive {
					s.active = true
				}

			} else {
				s.active = false
				s.deleteOn = false
				s.deleteCount = 0
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
	case EventRuneEdit:
		if s.active {
			s.editingRune = ev.Rune
			s.editingNow = true
		}
	case EventRuneComplete:
		if s.active {
			s.editingRune = 0
			s.editingNow = false
			s.text += string(ev.Rune)
			if s.onChange != nil {
				s.onChange(s, s.text)
			}
		}
	}

}

// Constructors for MTEdit
func MTEdit0() *MTEdit {
	temp := &MTEdit{
		text:  "",
		align: Align_LEFT | Align_VCENTER,
	}
	temp.SetMaterialColor(Material.Pallette.White)
	return temp
}

// Constructors for MTEdit
func MTEdit1(str string) *MTEdit {
	temp := &MTEdit{
		text:  str,
		align: Align_LEFT | Align_VCENTER,
	}
	temp.SetMaterialColor(Material.Pallette.White)
	return temp
}

// Constructors for MTEdit
func MTEdit2(str string, align Align) *MTEdit {
	temp := &MTEdit{
		text:  str,
		align: align,
	}
	temp.SetMaterialColor(Material.Pallette.White)
	return temp
}

// Constructors for MTEdit
func MTEdit3(mcl *MaterialColor, str string, align Align) *MTEdit {
	temp := &MTEdit{
		text:  str,
		align: align,
	}
	temp.SetMaterialColor(mcl)
	return temp
}

// ===============================================================================================
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
func (s *MTEdit) GetActive() bool {
	return !s.inactive
}
func (s *MTEdit) SetActive(active bool) {
	s.inactive = !active
}
func (s *MTEdit) OnChange(callback MTEditChange) {
	s.onChange = callback
}
func (s *MTEdit) ReferChange() MTEditChange {
	return s.onChange
}
