package main

import (
	"fmt"
	"github.com/iamGreedy/gumi"
	"github.com/iamGreedy/gumi/drawer"
	"github.com/iamGreedy/gumi/gumre"
	"image"
	"os"
	"runtime"
	"image/draw"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println(runtime.GOMAXPROCS(0))
	r, err := os.Open("./res/cubes_512.png")
	//r, err := os.Open("./res/172676.jpg")
	gumre.Assert(err)
	img, ext, err := image.Decode(r)
	gumre.Assert(err)
	rgba := image.NewRGBA(image.Rect(0,0,1024, 1024))
	draw.Draw(rgba, rgba.Rect, img, img.Bounds().Min, draw.Src)
	fmt.Println(ext)
	resizer := drawer.NewFillup(img, drawer.FillupNearest)
	resizer.Draw(rgba)
	//noise := drawer.NewNoise(32)
	//noise.Draw(rgba)
	//blur := drawer.NewBlur(20, false, drawer.BlurBox)
	//blur.Draw(rgba)
	gumi.Capture("out", rgba)
}

