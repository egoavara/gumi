package drawer

import (
	"image"
	"sync"
)

type Cache [][][STRIDECOUNT]uint8

func NewCache(rgba *image.RGBA) Cache {
	sz := rgba.Bounds().Size()
	temp := make([][][STRIDECOUNT]uint8, sz.X)
	for i := 0; i < sz.X; i++{
		temp[i] = make([][STRIDECOUNT]uint8, sz.Y)
	}
	var wg = new(sync.WaitGroup)
	for x := 0; x < sz.X; x ++{
		wg.Add(1)
		go func(x int) {
			for y := 0; y < sz.Y; y ++{
				offset := rgba.PixOffset(x, y)
				temp[x][y] = [STRIDECOUNT]uint8{
					rgba.Pix[offset + R],
					rgba.Pix[offset + G],
					rgba.Pix[offset + B],
					rgba.Pix[offset + A],
				}
			}
			wg.Done()
		}(x)
	}
	wg.Wait()
	return Cache(temp)
}
func (s Cache ) Commit(rgba *image.RGBA)  {
	var wg = new(sync.WaitGroup)
	for x, line := range s{
		wg.Add(1)
		go func(x int, line [][4]uint8) {
			for y, val := range line{
				offset := rgba.PixOffset(x, y)
				rgba.Pix[offset + R] = val[R]
				rgba.Pix[offset + G] = val[G]
				rgba.Pix[offset + B] = val[B]
				rgba.Pix[offset + A] = val[A]
			}
			wg.Done()
		}(x, line)
	}
	wg.Wait()
}
