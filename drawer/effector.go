package drawer

import "image/draw"

type Effector interface{
	Effect(image draw.Image)
}
