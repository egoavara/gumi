package sdl2example

import (
	"github.com/iamGreedy/gumi"
	"bytes"
	"github.com/iamGreedy/gumi/example/sdl2example/asset"
	"image/jpeg"
	"github.com/iamGreedy/gumi/media"
)

var HelloImage gumi.GUMI

func init() {
	img, err := jpeg.Decode(bytes.NewBuffer(asset.MustAsset("helloImage.jpg")))
	if err != nil {
		panic(err)
	}
	HelloImage = gumi.LinkingFrom(
		//gumi.NDrawing0(gumi.Drawing.FPS()),
		gumi.NBackground0(gumi.Material.Pallette.BackgroundDrawer()),
		gumi.Tool.MarginMinRegular(10,
			gumi.AImage0(media.NewFillup(img, media.FillupNearest)),
		),
	)
}
