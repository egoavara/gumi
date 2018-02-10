package drawer

import (
	"image/draw"
	"image"
)

type Drawer interface{
	image.Image
	Draw(dst draw.Image)
}