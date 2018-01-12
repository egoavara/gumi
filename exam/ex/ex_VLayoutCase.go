package ex

import (
	"github.com/iamGreedy/gumi"
	"testing"
)

func VLayoutCase(scr *gumi.Screen, theme gumi.Theme) (result testing.BenchmarkResult) {
	scr.Root(
		gumi.LinkingFrom(
			gumi.NStyle(theme.BackgroundStyle()),
			gumi.NBackground(),
			gumi.NMargin(gumi.AUTOSIZE, gumi.RegularBlank(gumi.Length{5, gumi.LENGTHMAX})),
			gumi.NVertical(
				gumi.LinkingFrom(
					gumi.NStyle(theme.ColorLine(0)),
					gumi.NBackground(),
					gumi.NMargin(gumi.AUTOSIZE, gumi.RegularBlank(gumi.Length{5, gumi.LENGTHMAX})),
					gumi.AText("1 : Hello, world!", gumi.Align_CENTER),
				),
				gumi.LinkingFrom(
					gumi.NStyle(theme.ColorLine(1)),
					gumi.NBackground(),
					gumi.NMargin(gumi.AUTOSIZE, gumi.RegularBlank(gumi.Length{5, gumi.LENGTHMAX})),
					gumi.AText("2 : Hello, world!", gumi.Align_CENTER),
				),
				gumi.LinkingFrom(
					gumi.NStyle(theme.ColorLine(2)),
					gumi.NBackground(),
					gumi.NMargin(gumi.AUTOSIZE, gumi.RegularBlank(gumi.Length{5, gumi.LENGTHMAX})),
					gumi.AText("3 : Hello, world!", gumi.Align_CENTER),
				),
				gumi.LinkingFrom(
					gumi.NStyle(theme.ColorLine(3)),
					gumi.NBackground(),
					gumi.NMargin(gumi.AUTOSIZE, gumi.RegularBlank(gumi.Length{5, gumi.LENGTHMAX})),
					gumi.AText("4 : Hello, world!", gumi.Align_CENTER),
				),
				gumi.LinkingFrom(
					gumi.NStyle(theme.ColorLine(4)),
					gumi.NBackground(),
					gumi.NMargin(gumi.AUTOSIZE, gumi.RegularBlank(gumi.Length{5, gumi.LENGTHMAX})),
					gumi.AText("5 : Hello, world!", gumi.Align_CENTER),
				),
			),
		),
	)
	result = testing.Benchmark(func(b *testing.B) {
		scr.Draw(nil)
	})
	return
}
