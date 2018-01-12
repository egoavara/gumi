package gumi

import (
	. "image"
	"image/color"
	"image/draw"
)

type Frame interface {
	draw.Image
	SubFrame(r Rectangle) Frame
	Rect() Rectangle
	Real() Rectangle
	ReferPix() []uint8
	Merge(frame Frame, startAt Point)
	Pix(pix []uint8)
	Copy() Frame
	rgbaCompatible() *RGBA
}
type gumiFrame struct {
	pix []uint8
	//
	stride   int
	bound    Rectangle
	rect     Rectangle
	fullRect Rectangle
}

func NewFrame(w, h int) Frame {
	return &gumiFrame{
		pix:      make([]uint8, 4*w*h),
		stride:   4 * w,
		bound:    Rect(0, 0, w, h),
		rect:     Rect(0, 0, w, h),
		fullRect: Rect(0, 0, w, h),
	}
}
func (p *gumiFrame) ColorModel() color.Model {
	return color.RGBAModel
}
func (p *gumiFrame) Bounds() Rectangle {
	return p.bound
}
func (p *gumiFrame) At(x, y int) color.Color {
	if !(Point{x, y}.In(p.bound)) {
		return color.RGBA{}
	}
	i := p.PixOffset(x, y)
	return color.RGBA{p.pix[i+0], p.pix[i+1], p.pix[i+2], p.pix[i+3]}
}
func (p *gumiFrame) PixOffset(x, y int) int {
	return (y+p.rect.Min.Y)*p.stride + (x+p.rect.Min.X)*4
}
func (p *gumiFrame) Set(x, y int, c color.Color) {
	if !(Point{x, y}.In(p.bound)) {
		return
	}
	i := p.PixOffset(x, y)
	c1 := color.RGBAModel.Convert(c).(color.RGBA)
	p.pix[i+0] = c1.R
	p.pix[i+1] = c1.G
	p.pix[i+2] = c1.B
	p.pix[i+3] = c1.A
}
func (p *gumiFrame) SubFrame(r Rectangle) Frame {
	r = r.Add(p.rect.Min).Intersect(p.rect)
	if r.Empty() {
		return &gumiFrame{}
	}

	return &gumiFrame{
		pix:      p.pix,
		stride:   p.stride,
		bound:    Rect(0, 0, r.Dx(), r.Dy()),
		rect:     r,
		fullRect: p.fullRect,
	}
}
func (p *gumiFrame) Rect() Rectangle {
	return p.rect
}
func (p *gumiFrame) Real() Rectangle {
	return p.fullRect
}
func (p *gumiFrame) ReferPix() []uint8 {
	return p.pix[p.PixOffset(0, 0):]
}
func (p *gumiFrame) Merge(frame Frame, startAt Point) {
	draw.Draw(
		p,
		frame.Bounds().Add(startAt).Intersect(p.rect),
		frame,
		frame.Bounds().Min,
		draw.Src,
	)
}
func (p *gumiFrame) Copy() Frame {
	result := &gumiFrame{
		pix:      make([]uint8, len(p.pix)),
		rect:     p.rect,
		stride:   p.stride,
		fullRect: p.fullRect,
	}
	copy(result.pix, p.pix)
	//
	return result
}
func (p *gumiFrame) Pix(pix []uint8) {
	p.pix = pix
}
func (p *gumiFrame) rgbaCompatible() *RGBA {
	return &RGBA{
		Stride: p.stride,
		Pix:    p.pix[p.PixOffset(0, 0):],
		Rect:   p.bound,
	}
}
