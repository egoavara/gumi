package main

import (
	"github.com/iamGreedy/gumi"
)

func main() {
	var scr = gumi.NewScreen(gumi.DefinedResolutions.Get("VGA"))
	// above line mean equal : var scr = gumi.NewScreen(640, 480)
	scr.Root(gumi.LinkingFrom(
		gumi.NStyle(gumi.DefaultDarkTheme.BackgroundStyle()),
		gumi.NBackground(),
		gumi.NMargin(gumi.AUTOSIZE, gumi.RegularBlank(gumi.MinLength(10))),
		gumi.NStyle(gumi.DefaultDarkTheme.Style(gumi.INTENSE3)),
		gumi.NBackground(),
		gumi.AText("Hello, world!", gumi.Align_CENTER),
	))
	scr.Draw(nil)
	gumi.Capture("out", scr.Frame())
}
