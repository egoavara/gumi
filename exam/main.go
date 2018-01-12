package main

import (
	"flag"
	"fmt"
	"github.com/golang/freetype"
	"github.com/iamGreedy/gumi"
	"io/ioutil"
	"testing"
	. "github.com/iamGreedy/gumi/exam/ex"
)

var (
	OutputFile  = flag.String("OF", "out", "OutputFile : Set output file name")
	Silent      = flag.Bool("SL", false, "SiLent : Do not capture Screen")
	ResultPrint = flag.Bool("RP", true, "ResultPrint : Print benchmark result")
	FontPath    = flag.String("FP", "./res/godoRounded R.ttf", "FontPath : Load font from this")
	Select      = flag.String(
		"S",
		"HelloWorld",
		"Select : "+StringsToString(SelectListStrings()...),
	)
	Resolution = flag.String("R", "VGA",
		"Resolution : "+StringsToString(gumi.DefinedResolutions.Kinds()...))
)

func LoadExampleTheme(fontpath string) gumi.Theme {
	bts, err := ioutil.ReadFile(fontpath)
	if err != nil {
		panic(err)
	}
	f, err := freetype.ParseFont(bts)
	if err != nil {
		panic(err)
	}
	theme := gumi.DefaultDarkTheme.From(gumi.Theme{
		Font: gumi.NewFont(f, 18),
	})
	return theme

}

func main() {
	flag.Set("S", IndexSelectString(5))
	flag.Parse()

	//
	var result testing.BenchmarkResult
	var Width, Height = gumi.DefinedResolutions.Get(*Resolution)
	if Width == 0 && Height == 0 {
		Width = 640
		Height = 480
	}
	var scr = gumi.NewScreen(Width, Height)
	var theme = LoadExampleTheme(*FontPath)
	//
	fn := NameSelectFunc(*Select)
	if fn == nil {
		fmt.Println("Invalid :", *Select)
		return
	}
	result = fn(scr, theme)
	fmt.Println("Run :", *Select)
	fmt.Println()
	if *ResultPrint {
		PrintResult(result)
	}
	if !*Silent {
		gumi.Capture(*OutputFile, scr.Frame())
	}
}
