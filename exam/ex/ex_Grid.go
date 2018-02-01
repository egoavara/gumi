package ex

import (
	"github.com/iamGreedy/gumi"
	"testing"
)

func Grid(scr *gumi.Screen, theme gumi.Theme) (result testing.BenchmarkResult) {
	scr.Root(gumi.LinkingFrom(
		gumi.NStyle0(theme.BackgroundStyle()),
		gumi.NBackground0(),
		gumi.NMargin0(gumi.AUTOSIZE, gumi.RegularBlank(gumi.MinLength(5))),
		gumi.NGrid(2, 2,
			gumi.LinkingFrom(
				gumi.NStyle0(theme.ColorLine(0)),
				gumi.NMargin0(gumi.AUTOSIZE, gumi.RegularBlank(gumi.MinLength(5))),
				gumi.NBackground0(),
				gumi.AText0("Hello, 0 world!", gumi.Align_CENTER),
			),
			gumi.LinkingFrom(
				gumi.NStyle0(theme.ColorLine(1)),
				gumi.NMargin0(gumi.AUTOSIZE, gumi.RegularBlank(gumi.MinLength(5))),
				gumi.NBackground0(),
				gumi.AText0("Hello, 1 world!", gumi.Align_CENTER),
			),
			gumi.LinkingFrom(
				gumi.NStyle0(theme.ColorLine(1)),
				gumi.NMargin0(gumi.AUTOSIZE, gumi.RegularBlank(gumi.MinLength(5))),
				gumi.NBackground0(),
				gumi.AText0("Hello, 2 world!", gumi.Align_CENTER),
			),
		),
	))

	result = testing.Benchmark(func(b *testing.B) {
		scr.Draw(nil)
	})
	return
}
