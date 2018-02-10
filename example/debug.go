package main

import (
	"image"
	"github.com/iamGreedy/gumi"
	"os"
	"github.com/iamGreedy/gumi/gumre"
	"fmt"
	"github.com/iamGreedy/gumi/drawer"
)

func main() {
	r, err := os.Open("./res/square.png")
	gumre.Assert(err)
	img, ext, err := image.Decode(r)
	gumre.Assert(err)
	fmt.Println(ext)
	drwer := drawer.NewFillup(img, drawer.FillupGausian)
	rgba := image.NewRGBA(image.Rect(0,0, 1024,1024))
	drwer.Draw(rgba)
	gumi.Capture("out", rgba)
}

func exper(dst *[]int, src ... int) {
	*dst = append(*dst, src...)
}

