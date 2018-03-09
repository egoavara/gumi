package sdl2example

import "github.com/iamGreedy/gumi"

var HelloWorld gumi.GUMI

func init()  {
	HelloWorld = gumi.LinkingFrom(
		//gumi.NDrawing0(gumi.Drawing.FPS()),
		gumi.NBackground0(gumi.Material.Pallette.BackgroundDrawer()),
		gumi.LCenter0(
			gumi.AText0("Hello, World!"),
		),
	)
}