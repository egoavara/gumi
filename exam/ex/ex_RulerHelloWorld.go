package ex

import (
	"github.com/iamGreedy/gumi"
	"testing"
)

func RulerHelloWorld(scr *gumi.Screen, theme gumi.Theme) (result testing.BenchmarkResult) {
	scr.Root(gumi.LinkingFrom(

		gumi.NStyle0(theme.ColorLine(2)),
		gumi.NDo0(func() {
			theme.Font.Use()
			defer theme.Font.Release()
			theme.Font.ChangeSize(30)
		}),
		gumi.NDrawing1(gumi.DrawListing(
			//gumi.BuildRuler(
			//	gumi.RULER_DOTS,
			//	20),
			//gumi.BuildRuler(
			//	gumi.RULER_HINT_VERTICAL|gumi.RULER_HINT_HORIZONTAL|gumi.RULER_GRID_HORIZONTAL|gumi.RULER_GRID_VERTICAL,
			//	100),
			gumi.BuildRuler(
				gumi.RULER_SCREEN|gumi.RULER_PROPORTION,
				100),
			//gumi.BuildRuler(
			//	gumi.RULER_GRIDDASH_VERTICAL | gumi.RULER_GRIDDASH_HORIZONTAL,
			//	100),
		)...),
		gumi.NDo0(func() {
			theme.Font.Use()
			defer theme.Font.Release()
			theme.Font.ChangeSize(16)
		}),
		gumi.NStyle0(theme.BackgroundStyle()),
		gumi.NBackground0(),
		gumi.NMargin0(gumi.AUTOSIZE, gumi.RegularBlank(gumi.MinLength(10))),
		gumi.NStyle0(theme.Style(gumi.INTENSE3)),
		gumi.NBackground0(),
		gumi.AText0("Hello, world!", gumi.Align_CENTER)))

	result = testing.Benchmark(func(b *testing.B) {
		scr.Draw(nil)
	})
	return
}
