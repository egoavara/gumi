package ex

import (
	"github.com/iamGreedy/gumi"
	"testing"
	"strconv"
)

func HLayoutCase(scr *gumi.Screen, theme gumi.Theme) (result testing.BenchmarkResult) {
	root := gumi.NStyle0(theme.BackgroundStyle())
	bgd1 := gumi.NBackground0()
	spc := gumi.NMargin0(gumi.AUTOSIZE, gumi.RegularBlank(gumi.Length{5, gumi.LENGTHMAX}))
	txt := gumi.AText0("Hello, world!", gumi.Align_CENTER)
	lay := gumi.LHorizontal0(
		gumi.LinkingFrom(
			gumi.NStyle0(theme.Style(gumi.INTENSE3)),
			gumi.NBackground0(),
			gumi.NMargin0(gumi.AUTOSIZE, gumi.RegularBlank(gumi.Length{5, gumi.LENGTHMAX})),
			txt,
		),
		gumi.LinkingFrom(
			gumi.NMargin0(gumi.AUTOSIZE.HModify(gumi.Length{5, gumi.LENGTHMAX}), gumi.NOBLANK),
			gumi.AEmpty0(),
		),
		gumi.LinkingFrom(
			gumi.NStyle0(theme.ColorLine(2)),
			gumi.NBackground0(),
			gumi.NMargin0(gumi.AUTOSIZE, gumi.RegularBlank(gumi.Length{5, gumi.LENGTHMAX})),
			gumi.AText0("Hello, world!", gumi.Align_LEFT|gumi.Align_BOTTOM),
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
