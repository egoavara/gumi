package main

import (
	"github.com/iamGreedy/gumi"
	"github.com/iamGreedy/gumi/gumre"
	"image"
)

type Test struct {
	gumi.VoidNode
}

func (s *Test) draw(frame *image.RGBA) {

}

func (Test) size() gumi.Size {
	return gumi.Size{
		gumi.AUTOLENGTH,
		gumi.AUTOLENGTH,
	}
}

func (Test) rect(r image.Rectangle) {

}

func (Test) update(info *gumi.Information, style *gumi.Style) {
}

func (Test) Occur(event gumi.Event) {
}

func (Test) String() string {
	return "test"
}

func main() {
	scr := gumi.NewScreen(gumre.DefinedResolutions.Get("HVGA"))
	t := &Test{}
	scr.Root(gumi.LinkingFrom(
		gumi.NDrawing1(
			gumi.Drawing.Ruler.Hint.Vertical(100),
			gumi.Drawing.Ruler.Hint.Horizontal(100),
		),
		gumi.NBackground0(gumi.Material.Pallette.BackgroundImage()),
		gumi.LVertical1(
			gumi.LVertical1(
				gumi.AText0("Test1"),
				t,
				gumi.Tool.MarginMinRegular(4, gumi.AText0("Test3")),
			),
			gumi.AText0("Test4"),
		),
	))
	scr.Init()

	scr.Update(nil, nil)
	scr.Ready()
	scr.Draw()
	gumi.Capture("out", scr.Frame())
}
