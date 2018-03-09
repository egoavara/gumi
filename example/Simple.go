package main

import (
	"github.com/iamGreedy/gumi"
	"github.com/iamGreedy/gumi/gumre"
	"github.com/iamGreedy/gumi/example/sdl2example"
)


func main() {
	scr := gumi.NewScreen(gcore.DefinedResolutions.Get("HVGA"))
	scr.Root(sdl2example.HelloWorld)
	scr.Init()
	scr.Update(gumi.Information{Dt:0})
	scr.Draw()
	gumi.Capture("out", scr.Frame())
}
