package sdl2example

import (
	"fmt"
	"github.com/iamGreedy/gumi"
)

var Buttons gumi.GUMI

func init() {
	Buttons = gumi.LinkingFrom(
		gumi.NBackground0(gumi.Material.Pallette.BackgroundDrawer()),
		gumi.LCenter0(
			gumi.LVertical1(
				gumi.Tool.MarginMinRegular(4,
					gumi.MTButton1(gumi.Material.Pallette.White, "Button 0", func(self *gumi.MTButton) {
						fmt.Println(self.Get(), ", Material Color : ", self.GetMaterialColor().String())
					}),
				),
				gumi.Tool.MarginMinRegular(4,
					gumi.MTButton1(gumi.Material.Pallette.Red, "Button 1", func(self *gumi.MTButton) {
						fmt.Println(self.Get(), ", Material Color : ", self.GetMaterialColor().String())
					}),
				),
				gumi.Tool.MarginMinRegular(4,
					gumi.MTButton1(gumi.Material.Pallette.Green, "Button 2", func(self *gumi.MTButton) {
						fmt.Println(self.Get(), ", Material Color : ", self.GetMaterialColor().String())
					}),
				),
				gumi.Tool.MarginMinRegular(4,
					gumi.MTButton1(gumi.Material.Pallette.Blue, "Button 3", func(self *gumi.MTButton) {
						fmt.Println(self.Get(), ", Material Color : ", self.GetMaterialColor().String())
					}),
				),
				gumi.Tool.MarginMinRegular(4,
					gumi.MTButton1(gumi.Material.Pallette.Yellow, "Button 4", func(self *gumi.MTButton) {
						fmt.Println(self.Get(), ", Material Color : ", self.GetMaterialColor().String())
					}),
				),
			),
		),
	)
}
