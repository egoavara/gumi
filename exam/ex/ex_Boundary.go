package ex

import (
	"github.com/iamGreedy/gumi"
	"testing"
)

func Boundary(scr *gumi.Screen, theme gumi.Theme) (result testing.BenchmarkResult) {
	scr.Root(gumi.LinkingFrom(
		gumi.NStyle(theme.BackgroundStyle()),
		gumi.NBackground(),
		gumi.NMargin(gumi.AUTOSIZE, gumi.RegularBlank(gumi.MinLength(10))),

		gumi.NStyle(theme.Style(gumi.INTENSE3)),
		gumi.NBackground(),
		gumi.NStyle(theme.ColorLine(0)),
		gumi.NBoundary(gumi.BOUNDARY_ALL),
		gumi.NStyle(theme.Style(gumi.INTENSE3)),
		gumi.AText("Hello, world!", gumi.Align_CENTER)))

	result = testing.Benchmark(func(b *testing.B) {
		scr.Draw(nil)
	})
	return
}
