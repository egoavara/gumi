package drawer

import (
	"image/draw"
	"math/rand"
	"math"
	"sync"
)

type Noise struct {
	Intense float64
}

func NewNoise(intense float64) *Noise {
	return &Noise{
		Intense:intense,
	}
}

func (s Noise) Draw(target draw.Image) {
	var dst, dstStride, startX, endY = startEditPix(target)
	var dstW, dstH = sizeEidtPix(len(dst), dstStride)
	defer endEditPix(target, dst, dstStride, startX, endY )
	var wg = new(sync.WaitGroup)
	//
	for y := 0; y < dstH; y++ {
		wg.Add(1)
		go func(innerY int) {
			for x := 0; x < dstW; x++ {
				offset := offsetEditPix(x,innerY,dstStride)
				delta :=  noiseGausian(s.Intense)
				dst[offset+R] = uint8(clamp(delta + float64(dst[offset+R]), 0, math.MaxUint8))
				dst[offset+G] = uint8(clamp(delta + float64(dst[offset+G]), 0, math.MaxUint8))
				dst[offset+B] = uint8(clamp(delta + float64(dst[offset+B]), 0, math.MaxUint8))
			}
			wg.Done()
		}(y)
	}
	wg.Wait()
}
func noiseGausian(intense float64) float64 {
	return rand.NormFloat64()*intense
}

