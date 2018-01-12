package gumi

const (
	Align_TOP     = 0x00
	Align_VCENTER = 0x01
	Align_BOTTOM  = 0x02
	Align_LEFT    = 0x00
	Align_HCENTER = 0x10
	Align_RIGHT   = 0x20
	//

	Align_CENTER = Align_HCENTER | Align_VCENTER
)

type Align uint8

func ParseAlign(a Align) (v uint8, h uint8) {
	if a&Align_BOTTOM == Align_BOTTOM {
		v = Align_BOTTOM
	} else if a&Align_VCENTER == Align_VCENTER {
		v = Align_VCENTER
	} else {
		v = Align_TOP
	}
	if a&Align_RIGHT == Align_RIGHT {
		h = Align_RIGHT
	} else if a&Align_HCENTER == Align_HCENTER {
		h = Align_HCENTER
	} else {
		h = Align_LEFT
	}
	return

}
