# GUMI
![Version](https://img.shields.io/badge/version-0.0.1-green.svg)
![Warn](https://img.shields.io/badge/status-alpha-red.svg)
![License](https://img.shields.io/badge/license-BSD%202--clause-blue.svg)

#### Simple GUI Renderer writen by go
If you want to render **Game GUI**
Or you need **Native go gui** renderer, it perfect for you 



## Getting Start
Install by `go get github.com/iamGreedy/gumi`

## Example
```go
package main

import (
	"github.com/iamGreedy/gumi"
)

func main() {
	var scr = gumi.NewScreen(gumi.DefinedResolutions.Get("VGA"))
	// above line mean equal : var scr = gumi.NewScreen(640, 480)
	scr.Root(gumi.LinkingFrom(
		gumi.NStyle(gumi.DefaultDarkTheme.BackgroundStyle()),
		gumi.NBackground(),
		gumi.NMargin(gumi.AUTOSIZE, gumi.RegularBlank(gumi.MinLength(10))),
		gumi.NStyle(gumi.DefaultDarkTheme.Style(gumi.INTENSE3)),
		gumi.NBackground(),
		gumi.AText("Hello, world!", gumi.Align_CENTER),
	))
	scr.Draw(nil)
	gumi.Capture("out", scr.Frame())
}
```

There is also some example in `./exam`.

You can build and run by follow command  
```
cd ./exam
go build
```
You can inform by run `./exam -help`
and you can change example by use `-S` option.

If you want to run `HelloWorld` example, you just type 
```
./exam -S HelloWorld
```
Want to know all available example? just use `-help` option. 

Or you can find SelectList on `./exam/ex/SelectList.go` like
```go
var SelectList = []func(*gumi.Screen, gumi.Theme) testing.BenchmarkResult{...}
```

## Documentation
_Working in progress..._

## History
You can read history [here](https://github.com/iamGreedy/gumi/blob/master/HISTORY.md)
## License
[Copyright (c) 2017 iamGreedy BSD 2-clause](https://github.com/iamGreedy/gumi/blob/master/LICENSE.md)