package main

import (
	"github.com/iamGreedy/gumi"
	"image"
	"runtime"
	"image/jpeg"
	"bytes"
	"github.com/iamGreedy/gumi/example/sdl2example/asset"
	"github.com/iamGreedy/gumi/drawer"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	img, err := jpeg.Decode(bytes.NewBuffer(asset.MustAsset("helloImage.jpg")))
	if err != nil {
		panic(err)
	}
	drw := media.NewFillup(img, media.FillupNearest)

	save := image.NewRGBA(image.Rect(0,0,800, 600))
	sub := save.SubImage(image.Rect(100, 100, 700, 500)).(*image.RGBA)
	//
	drw.Draw(sub)
	gumi.Capture("out", save)

}
