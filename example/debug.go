package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"fmt"
)

func main() {
	var version sdl.Version
	sdl.GetVersion(&version)
	fmt.Println(version)
}
