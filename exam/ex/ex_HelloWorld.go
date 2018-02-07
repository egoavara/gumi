package ex

import (
	"github.com/iamGreedy/gumi"
	"testing"
)

func HelloWorld(scr *gumi.Screen, theme gumi.Theme) (result testing.BenchmarkResult) {
	scr.Root(gumi.LinkingFrom(
		gumi.NStyle0(theme.BackgroundStyle()),
		gumi.NBackground0(),
		gumi.NMargin0(gumi.AUTOSIZE, gumi.RegularBlank(gumi.MinLength(10))),
		gumi.NStyle0(theme.Style(gumi.INTENSE3)),
		gumi.NBackground0(),
		gumi.AText1("Hello, world!", gumi.Align_CENTER)))

	result = testing.Benchmark(func(b *testing.B) {
		scr.Draw(nil)
	})
	return
}
