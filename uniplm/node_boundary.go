package uniplm

import (
	"github.com/fogleman/gg"
)

type BoundarySide uint8

const (
	BOUNDARY_TOP    BoundarySide = 1 << iota
	BOUNDARY_BOTTOM BoundarySide = 1 << iota
	BOUNDARY_LEFT   BoundarySide = 1 << iota
	BOUNDARY_RIGHT  BoundarySide = 1 << iota
	BOUNDARY_ALL                 = BOUNDARY_TOP | BOUNDARY_BOTTOM | BOUNDARY_LEFT | BOUNDARY_RIGHT
)

type nBoundary struct {
	GUMILINK_SINGLE
	bs BoundarySide
}

func NBoundary(bs BoundarySide) *nBoundary {
	return &nBoundary{
		bs: bs,
	}
}
func (s *nBoundary) size(drawing *Drawing, style *Style) Size {
	return s.child.(GUMIElem).size(drawing, style)
}
func (s *nBoundary) draw(drawing *Drawing, style *Style, frame Frame) {
	rect := frame.Bounds()
	ctx := gg.NewContextForRGBA(frame.rgbaCompatible())
	ctx.SetLineWidth(style.LineWidth)
	ctx.SetColor(style.Line.At(0, 0))
	if s.bs&BOUNDARY_TOP == BOUNDARY_TOP {
		ctx.DrawLine(
			0,
			style.LineWidth/2,
			float64(rect.Max.X),
			style.LineWidth/2)
	}
	if s.bs&BOUNDARY_BOTTOM == BOUNDARY_BOTTOM {
		ctx.DrawLine(
			0,
			float64(rect.Max.Y)-style.LineWidth/2,
			float64(rect.Max.X),
			float64(rect.Max.Y)-style.LineWidth/2)
	}
	if s.bs&BOUNDARY_LEFT == BOUNDARY_LEFT {
		ctx.DrawLine(
			0+style.LineWidth/2,
			0,
			0+style.LineWidth/2,
			float64(rect.Max.Y))
	}
	if s.bs&BOUNDARY_RIGHT == BOUNDARY_RIGHT {
		ctx.DrawLine(
			float64(rect.Max.X)-style.LineWidth/2,
			0,
			float64(rect.Max.X)-style.LineWidth/2,
			float64(rect.Max.Y))
	}
	ctx.Stroke()
	s.child.(GUMIElem).draw(drawing, style, frame)
}
