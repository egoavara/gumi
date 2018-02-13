package res

import (
	"image"
	"github.com/fogleman/gg"
	"math"
	"image/color"
)
var ImageHexagon image.Image

func init() {
	ctx := gg.NewContext(120, 120)
	//
	ctx.SetRGBA(0,0,0, 0)
	ctx.SetColor(color.White)
	centerX, centerY := float64(ctx.Width()) / 2, float64(ctx.Height()) / 2
	radius := 40.
	var sin, cos float64
	//
	sin = math.Sin(math.Pi / 3 * 0)
	cos = math.Cos(math.Pi / 3 * 0)
	ctx.MoveTo(centerX + radius * cos, centerY + radius * sin)
	for i := 1 ; i <= 6; i++{
		sin = math.Sin(math.Pi / 3 * float64(i))
		cos = math.Cos(math.Pi / 3 * float64(i))
		ctx.LineTo(centerX + radius * cos, centerY + radius * sin)
	}
	ctx.Fill()
	ImageHexagon = ctx.Image()
}