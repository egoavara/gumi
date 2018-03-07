package sdl2example

import (
	"github.com/iamGreedy/gumi"
	"fmt"
)

var HelloButton gumi.GUMI

func init() {
	HelloButton = gumi.LinkingFrom(
		//gumi.NDrawing0(gumi.Drawing.FPS()),
		gumi.NBackground0(gumi.Material.Pallette.BackgroundDrawer()),
		gumi.LCenter0(
			gumi.MTButton0("Hello, Button!", func(self *gumi.MTButton) {
				fmt.Println("Hello, Button!")
			}),
		),
	)
}
