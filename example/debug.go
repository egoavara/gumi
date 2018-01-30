package main

import (
	"github.com/iamGreedy/gumi"
	"image"
	"github.com/fogleman/gg"
	"image/color"
)

func main() {
	ori := image.NewRGBA(image.Rect(0,0,800,600))
	img := ori.SubImage(image.Rect(400,300, 800, 600)).(*image.RGBA)
	img.Rect = image.Rect(0,0,img.Rect.Dx(), img.Rect.Dy())

	ctx := gg.NewContextForRGBA(img)
	ctx.SetLineWidth(3)
	var w, h = float64(ctx.Width()), float64(ctx.Height())
	radius := h/ 2
	ctx.SetColor(color.RGBA{255, 0, 0, 255})
	ctx.DrawArc(radius, radius, radius, gg.Radians(90), gg.Radians(270))

	ctx.DrawArc(w - radius, radius, radius, gg.Radians(-90), gg.Radians(90))
	ctx.DrawRectangle(radius, 0 ,w - radius * 2, h)
	ctx.Fill()
	//ctx.Stroke()
	gumi.Capture("out1", ori)
}