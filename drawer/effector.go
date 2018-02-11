package drawer

import "image/draw"

type Effector interface{
	Effect(target draw.Image)
}
