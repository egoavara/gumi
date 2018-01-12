package ex

import (
	"github.com/iamGreedy/gumi"
	"testing"
	"strconv"
)

func HLayoutCase(scr *gumi.Screen, theme gumi.Theme) (result testing.BenchmarkResult) {
	root := gumi.NStyle(theme.BackgroundStyle())
	bgd1 := gumi.NBackground()
	spc := gumi.NMargin(gumi.AUTOSIZE, gumi.RegularBlank(gumi.Length{5, gumi.LENGTHMAX}))
	txt := gumi.AText("Hello, world!", gumi.Align_CENTER)
	lay := gumi.NHorizontal(
		gumi.LinkingFrom(
			gumi.NStyle(theme.Style(gumi.INTENSE3)),
			gumi.NBackground(),
			gumi.NMargin(gumi.AUTOSIZE, gumi.RegularBlank(gumi.Length{5, gumi.LENGTHMAX})),
			txt,
		),
		gumi.LinkingFrom(
			gumi.NMargin(gumi.AUTOSIZE.HModify(gumi.Length{5, gumi.LENGTHMAX}), gumi.NOBLANK),
			gumi.AEmpty(),
		),
		gumi.LinkingFrom(
			gumi.NStyle(theme.ColorLine(2)),
			gumi.NBackground(),
			gumi.NMargin(gumi.AUTOSIZE, gumi.RegularBlank(gumi.Length{5, gumi.LENGTHMAX})),
			gumi.AText("Hello, world!", gumi.Align_LEFT|gumi.Align_BOTTOM),
		),
	)
	scr.Root(gumi.LinkingFrom(root, bgd1, spc, lay))

	txt.Set("F : 1")
	result = testing.Benchmark(func(b *testing.B) {
		txt.Set("o : " + strconv.FormatInt(int64(b.N), 10))
		scr.Draw(nil)
	})
	return
}
