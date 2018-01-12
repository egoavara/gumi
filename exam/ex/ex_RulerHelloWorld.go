package ex

import (
	"github.com/iamGreedy/gumi"
	"testing"
)

func RulerHelloWorld(scr *gumi.Screen, theme gumi.Theme) (result testing.BenchmarkResult) {
	scr.Root(gumi.LinkingFrom(

		gumi.NStyle(theme.ColorLine(2)),
		gumi.NDo(func() {
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
		gumi.NDo(func() {
			theme.Font.Use()
			defer theme.Font.Release()
			theme.Font.ChangeSize(16)
		}),
		gumi.NStyle(theme.BackgroundStyle()),
		gumi.NBackground(),
		gumi.NMargin(gumi.AUTOSIZE, gumi.RegularBlank(gumi.MinLength(10))),
		gumi.NStyle(theme.Style(gumi.INTENSE3)),
		gumi.NBackground(),
		gumi.AText("Hello, world!", gumi.Align_CENTER)))

	result = testing.Benchmark(func(b *testing.B) {
		scr.Draw(nil)
	})
	return
}
