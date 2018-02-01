package ex

import (
	"github.com/iamGreedy/gumi"
	"testing"
)

func VLayoutCase(scr *gumi.Screen, theme gumi.Theme) (result testing.BenchmarkResult) {
	scr.Root(
		gumi.LinkingFrom(
			gumi.NStyle0(theme.BackgroundStyle()),
			gumi.NBackground0(),
			gumi.NMargin0(gumi.AUTOSIZE, gumi.RegularBlank(gumi.Length{5, gumi.LENGTHMAX})),
			gumi.NVertical0(
				gumi.LinkingFrom(
					gumi.NStyle0(theme.ColorLine(0)),
					gumi.NBackground0(),
					gumi.NMargin0(gumi.AUTOSIZE, gumi.RegularBlank(gumi.Length{5, gumi.LENGTHMAX})),
					gumi.AText0("1 : Hello, world!", gumi.Align_CENTER),
				),
				gumi.LinkingFrom(
					gumi.NStyle0(theme.ColorLine(1)),
					gumi.NBackground0(),
					gumi.NMargin0(gumi.AUTOSIZE, gumi.RegularBlank(gumi.Length{5, gumi.LENGTHMAX})),
					gumi.AText0("2 : Hello, world!", gumi.Align_CENTER),
				),
				gumi.LinkingFrom(
					gumi.NStyle0(theme.ColorLine(2)),
					gumi.NBackground0(),
					gumi.NMargin0(gumi.AUTOSIZE, gumi.RegularBlank(gumi.Length{5, gumi.LENGTHMAX})),
					gumi.AText0("3 : Hello, world!", gumi.Align_CENTER),
				),
				gumi.LinkingFrom(
					gumi.NStyle0(theme.ColorLine(3)),
					gumi.NBackground0(),
					gumi.NMargin0(gumi.AUTOSIZE, gumi.RegularBlank(gumi.Length{5, gumi.LENGTHMAX})),
					gumi.AText0("4 : Hello, world!", gumi.Align_CENTER),
				),
				gumi.LinkingFrom(
					gumi.NStyle0(theme.ColorLine(4)),
					gumi.NBackground0(),
					gumi.NMargin0(gumi.AUTOSIZE, gumi.RegularBlank(gumi.Length{5, gumi.LENGTHMAX})),
					gumi.AText0("5 : Hello, world!", gumi.Align_CENTER),
				),
			),
		),
	)
	result = testing.Benchmark(func(b *testing.B) {
		scr.Draw(nil)
	})
	return
}
