package gumi

import (
	"fmt"
	"github.com/fogleman/gg"
	"image"
)

const (
	MTDropboxMinWidth             = 80
	MTDropboxScroolWidth          = 8
	MTDropboxElemMargin           = 8
	MTDropboxElemUnderline        = 2
	MTDropboxMinHeight            = 20
	MTDropboxAnimStretchPerSecond = 400
	MTDropboxScroolModify         = 16
)

type MTDropbox struct {
	VoidStructure
	boundStore
	styleStore
	//
	mtColorSingle
	//
	Elems    mtDropboxElemList
	selected int

	inactive bool
	maxbox   int
	cutbox   int
	//
	scr     *Screen
	deferid uint64
	hookid  uint64
	//
	onChange MTDropboxChange
	//
	hover               int
	cursorEnter, active bool
	box, boxTo          int
	scrool, scroolTo    int
}

type MTDropboxChange func(self *MTDropbox, selected string)

func (s *MTDropbox) init() {
	s.scr = getScreen(s)
	s.deferid = s.scr.deferReserve()
	s.hookid = s.scr.hookReserve()
}
func (s *MTDropbox) String() string {
	return fmt.Sprintf("%s(select:%s)", "MTDropbox", s.Elems[s.selected])
}
func (s *MTDropbox) draw(frame *image.RGBA) {
	var baseColor, mainColor = s.GetMaterialColor().Color()
	if s.box > 0 && s.Elems.Length() > 0 {
		s.scr.deferRequest(s.deferid, func(rgba *image.RGBA) {
			var radius = float64(s.bound.Dy()) / 2
			var box int
			if s.maxbox > 0 && s.box > s.maxbox {
				s.cutbox = s.box - s.maxbox
				box = s.maxbox
			} else {
				box = s.box
			}
			if s.cutbox < s.scroolTo {
				s.scroolTo = s.cutbox
				s.scrool = s.cutbox
			}
			if s.bound.Max.Y+box+int(radius) > frame.Rect.Max.Y {
				diff := s.bound.Max.Y + box + int(radius) - frame.Rect.Max.Y
				box -= diff
				s.cutbox += diff
			}

			// selecte, background, scrool
			func() {
				bd := s.bound
				bd.Max.Y += box + int(radius)
				var ctx = GGContextRGBASub(rgba, bd)
				s.style.useContext(ctx)
				defer s.style.releaseContext(ctx)
				//
				var w, h = float64(ctx.Width()), float64(ctx.Height())
				// background
				ctx.SetColor(baseColor)
				ctx.DrawArc(radius, radius, radius, gg.Radians(180), gg.Radians(270))
				ctx.DrawArc(radius, h-radius, radius, gg.Radians(90), gg.Radians(180))
				ctx.DrawRectangle(radius-1, 0, w-radius*2+1, h)
				ctx.DrawArc(w-radius, radius, radius, gg.Radians(-90), gg.Radians(0))
				ctx.DrawArc(w-radius, h-radius, radius, gg.Radians(0), gg.Radians(90))
				ctx.Fill()
				// outline
				ctx.SetColor(Scale.Color(baseColor, mainColor, float64(s.box)/float64(s.boxTo)))
				ctx.DrawArc(radius, radius, radius, gg.Radians(180), gg.Radians(270))
				ctx.DrawLine(radius, 0, w-radius, 0)
				ctx.DrawArc(w-radius, radius, radius, gg.Radians(270), gg.Radians(360))
				ctx.DrawLine(w, radius, w, h-radius)
				ctx.DrawArc(w-radius, h-radius, radius, gg.Radians(0), gg.Radians(90))
				ctx.DrawLine(w-radius, h, radius, h)
				ctx.DrawArc(radius, h-radius, radius, gg.Radians(90), gg.Radians(180))
				ctx.DrawLine(0, h-radius, 0, radius)
				ctx.Stroke()
				// selected underline
				ctx.SetColor(mainColor)
				ctx.Push()
				ctx.SetLineWidth(.25)
				ctx.DrawLine(radius, float64(s.bound.Dy()), w-2*radius, float64(s.bound.Dy()))
				ctx.Stroke()
				//

				ctx.Pop()
				// selected
				selectedElem := s.Elems.getForDraw(s.selected)
				if len(selectedElem.content) > 0 {
					ctx.DrawString(selectedElem.content, radius, (float64(s.bound.Dy())-float64(selectedElem.h))/2+float64(selectedElem.h))
					ctx.Stroke()
				}
				// scroll
				percent := float64(box) / float64(s.Elems.heightSum()+MTDropboxElemMargin*s.Elems.Length())
				scroolPercent := float64(s.scrool) / float64(box+s.cutbox)
				if percent < 0 {
					percent = 0
				}
				if percent > 1 {
					percent = 1
				}

				ctx.DrawArc(w-radius, radius+(scroolPercent)*(h-radius*2), MTDropboxScroolWidth/2, gg.Radians(180), gg.Radians(360))
				ctx.DrawRectangle(w-radius-MTDropboxScroolWidth/2, radius+(scroolPercent)*(h-radius*2), MTDropboxScroolWidth, percent*(h-radius*2))
				ctx.DrawArc(w-radius, radius+(scroolPercent)*(h-radius*2)+percent*(h-radius*2), MTDropboxScroolWidth/2, gg.Radians(0), gg.Radians(180))
				ctx.Fill()
			}()
			// elems, hover
			func() {
				bd := s.bound
				bd.Min.Y = s.bound.Max.Y
				bd.Max.Y = s.bound.Max.Y + box + int(radius)
				var ctx = GGContextRGBASub(rgba, bd)
				s.style.useContext(ctx)
				defer s.style.releaseContext(ctx)
				sum := MTDropboxElemMargin
				ctx.SetColor(mainColor)
				for i, v := range s.Elems.refer() {
					drawY := float64(sum+v.h) - float64(s.scrool)
					ctx.DrawString(v.content, radius, drawY)
					ctx.Stroke()
					if i == s.hover {
						ctx.DrawLine(radius, drawY+MTDropboxElemUnderline, radius+float64(v.w), drawY+MTDropboxElemUnderline)
						ctx.Stroke()
					}
					sum += v.h + MTDropboxElemMargin
				}
			}()
		})
	} else {
		s.scr.deferRequest(s.deferid, nil)
		var ctx = GGContextRGBASub(frame, s.bound)
		s.style.useContext(ctx)
		defer s.style.releaseContext(ctx)
		//
		var w, h = float64(ctx.Width()), float64(ctx.Height())
		var radius = float64(s.bound.Dy()) / 2
		//
		ctx.SetColor(baseColor)
		ctx.DrawArc(radius, radius, radius, gg.Radians(90), gg.Radians(270))
		ctx.DrawRectangle(radius, 0, w-radius*2, h)
		ctx.DrawArc(w-radius, radius, radius, gg.Radians(-90), gg.Radians(90))
		ctx.Fill()
		//
		ctx.SetColor(mainColor)
		elem := s.Elems.getForDraw(s.selected)
		if len(elem.content) > 0 {
			ctx.DrawString(elem.content, radius, (h-float64(elem.h))/2+float64(elem.h))
			ctx.Stroke()
		}
		ctx.DrawCircle(w-radius, radius, MTDropboxScroolWidth/2)
		ctx.Fill()
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
	if s.style != style || s.Elems.needUpdate() {
		s.Elems.update(style)
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
		s.box = int(fbox)
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
		s.scrool = int(fscrool)
	}
}
func (s *MTDropbox) Occur(event Event) {
	if s.inactive {
		s.cursorEnter = false
		s.active = false
		return
	}
	switch ev := event.(type) {
	case EventKeyRelease:
		switch ev.Key {
		case KEY_MOUSE1:
			if s.cursorEnter {
				// 정지상태인 경우 클릭 무시
				if !s.active {
					// 선택상태가 아니였을 경우 선택상태로 전환, 이벤트 후킹 실시
					s.active = true
					s.boxTo = s.Elems.heightSum() + (s.Elems.Length())*MTDropboxElemMargin
					s.scr.hookRequest(s.hookid, func(event Event) Event {
						if v, ok := event.(EventCursor); ok {
							bd := s.bound
							bd.Max.Y += s.box
							if bd.Min.X <= int(v.X) && int(v.X) < bd.Max.X && bd.Min.Y <= int(v.Y) && int(v.Y) < bd.Max.Y {
								s.Occur(event)
								return nil
							}
						}
						return event
					})
				} else {
					// 선택상태, 커서 진입인 경우 selected, 선택하기
					if s.hover >= 0 {
						s.selected = s.hover
						s.active = false
						s.boxTo = 0
						s.scr.hookRequest(s.hookid, nil)
						if s.onChange != nil {
							s.onChange(s, s.Elems.Get(s.selected))
						}
					}
				}
			} else {
				// 다른 곳 클릭시 선택 취소
				s.active = false
				s.boxTo = 0
				s.scr.hookRequest(s.hookid, nil)
			}
		}

	case EventCursor:
		x := int(ev.X)
		y := int(ev.Y)
		bd := s.bound
		bd.Max.Y += s.box
		if (bd.Min.X <= x && x < bd.Max.X) && (bd.Min.Y <= y && y < bd.Max.Y) {
			s.cursorEnter = true
			sum := s.bound.Max.Y + MTDropboxElemMargin
			if y >= s.bound.Max.Y {
				for i, elem := range s.Elems.refer() {
					if sum <= y+s.scrool && y+s.scrool < sum+elem.h+MTDropboxElemMargin {
						s.hover = i
						break
					}
					sum += elem.h + MTDropboxElemMargin
				}
			} else {
				s.hover = -1
			}
		} else {
			s.hover = -1
			s.cursorEnter = false
		}
	case EventScroll:
		s.scroolTo += MTDropboxScroolModify * int(ev.Y)
		if s.scroolTo < 0 {
			s.scroolTo = 0
		}
	}
}

//
func MTDropbox0() *MTDropbox {
	res := &MTDropbox{
		Elems:    mtDropboxElemList{},
		selected: 0,
		hover:    -1,
	}
	res.SetMaterialColor(Material.Pallette.White)
	return res
}
func MTDropbox1(maxboxlen uint16) *MTDropbox {
	res := &MTDropbox{
		Elems:    mtDropboxElemList{},
		selected: 0,
		hover:    -1,
		maxbox:   int(maxboxlen),
	}
	res.SetMaterialColor(Material.Pallette.White)
	return res
}
func MTDropbox2(change MTDropboxChange) *MTDropbox {
	res := &MTDropbox{
		Elems:    mtDropboxElemList{},
		selected: 0,
		hover:    -1,
	}
	res.SetMaterialColor(Material.Pallette.White)
	res.OnChange(change)
	return res
}
func MTDropbox3(change MTDropboxChange, elems ...string) *MTDropbox {
	res := &MTDropbox{
		Elems:    mtDropboxElemList{},
		selected: 0,
		hover:    -1,
	}
	res.SetMaterialColor(Material.Pallette.White)
	res.OnChange(change)
	for i, v := range elems {
		res.Elems.Set(i, v)
	}
	return res
}
func (s *MTDropbox) OnChange(callback MTDropboxChange) {
	s.onChange = callback
}
func (s *MTDropbox) ReferChange() MTDropboxChange {
	return s.onChange
}

func (s *MTDropbox) GetMaxboxLength() uint16 {
	return uint16(s.maxbox)
}
func (s *MTDropbox) SetMaxboxLength(l uint16) {
	s.maxbox = int(l)
}
