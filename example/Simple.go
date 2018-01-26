package main

import (
	"github.com/iamGreedy/gumi"
	"github.com/iamGreedy/gumi/gutl"
)

func main() {
	scr := gumi.NewScreen(gutl.DefinedResolutions.Get("VGA"))
	scr.Root(gumi.LinkingFrom(
		gumi.NDrawing1(gumi.BuildRuler(
			gumi.RULER_HINT_VERTICAL |gumi.RULER_HINT_HORIZONTAL,
			100,
		)...),
		gumi.NBackground(),
		gumi.NMargin(gumi.Blank{
			gumi.MinLength(5),
			gumi.MinLength(5),
			gumi.MinLength(5),
			gumi.MinLength(5),
		}),
		gumi.NStyle(gumi.DefaultDarkTheme.Style(gumi.INTENSE3)),
		gumi.NBackground(),
		gumi.AText("Hello, World!", gumi.Align_CENTER),
	))
	scr.Update(nil, nil)
	scr.Ready()
	scr.Draw()
	gumi.Capture("out", scr.Frame())
}
