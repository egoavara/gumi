package main

import (
	"github.com/iamGreedy/gumi"
	"github.com/iamGreedy/gumi/gutl"
)

func main() {
	scr := gumi.NewScreen(gutl.DefinedResolutions.Get("HVGA"))
	scr.Root(gumi.LinkingFrom(
		gumi.NDrawing1(
			gumi.Drawing.Ruler.Hint.Vertical(100),
			gumi.Drawing.Ruler.Hint.Horizontal(100),
		),
		gumi.NBackground0(),
		gumi.LVertical1(
			gumi.LVertical1(
				gumi.AText1("Test1"),
				gumi.AText1("Test2"),
				gumi.Tool.MarginMinRegular(4, gumi.AText1("Test3")),
			),
			gumi.AText1("Test4"),
		),
	))
	scr.Update(nil, nil)
	scr.Ready()
	scr.Draw()
	gumi.Capture("out", scr.Frame())
}
