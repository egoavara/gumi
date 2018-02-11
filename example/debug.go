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
	//r, err := os.Open("./res/172676.jpg")
	r, err := os.Open("./res/cubes_512.png")
	gumre.Assert(err)
	img, ext, err := image.Decode(r)
	gumre.Assert(err)
	fmt.Println(ext)
	var res = img.(draw.Image)
	blur := drawer.NewBlur(20, false, drawer.BlurBox)
	blur.Effect(res)
	gumi.Capture("out", res)
}

