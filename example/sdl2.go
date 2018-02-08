package main

import (
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/golang/freetype/truetype"
	"github.com/iamGreedy/gumi"
	"github.com/iamGreedy/gumi/glumi"
	"github.com/iamGreedy/gumi/gutl"
	"github.com/iamGreedy/gumi/res"
	"github.com/veandco/go-sdl2/sdl"
	"runtime"
	"time"
)

func main() {
	var width, height = gutl.DefinedResolutions.Get("HD")
	var err error

	// init go:runtime
	GoRuntimeInit()
	// Init SDL2
	SDL2Init()
	wnd, ctx := SDL2Window(width, height)
	sdl.GLMakeCurrent(wnd, ctx)
	defer sdl.GLDeleteContext(ctx)
	// Init GL
	GLInit()
	fmt.Println("OpenGL version : ", gl.GoStr(gl.GetString(gl.VERSION)))
	// Init GUMI
	GUMIInit()
	// Init GLumi
	GLUMIInit()
	// GLumi Object allocate
	lumi := glumi.NewGLUMI()
	// window build

	//window.SetCharModsCallback(func(w *glfw.Window, char rune, mods glfw.ModifierKey) {
	//	fmt.Println(string([]rune{char}), mods)
	//})
	//
	modal := gumi.ALModal0()
	modal.SetModal(gumi.LinkingFrom(
		gumi.NBackground0(gumi.Material.Pallette.SilluetImage()),
		gumi.LCenter0(gumi.NMinimum0(gumi.AxisBoth, gumi.MTButton0("Modal, Hello!", func(self *gumi.MTButton) {
			modal.SetShow(!modal.GetShow())
		}))),
	))
	scr := gumi.NewScreen(width, height)
	scr.Root(gumi.LinkingFrom(
		modal,
		gumi.NBackground0(gumi.Material.Pallette.BackgroundImage()),
		gumi.NDrawing1(
			gumi.Drawing.FPS(),
			//gumi.Drawing.Ruler.Screen(),
			gumi.Drawing.Ruler.Hint.Vertical(100),
			gumi.Drawing.Ruler.Hint.Horizontal(100),
		),
		gumi.NMargin0(gumi.RegularBlank(gumi.MinLength(20))),
		gumi.LVertical1(
			gumi.Tool.MarginMinRegular(4, gumi.MTButton1(gumi.Material.Pallette.Red, "Close", func(self *gumi.MTButton) {
				err = glumi.Stop
			})),
			//
			gumi.ASpacer2(gumi.MinLength(20)),
			//
			gumi.LVertical1(ToggleProgress...),
			gumi.LHorizontal1(Radios...),
			//
			gumi.ASpacer2(gumi.MinLength(20)),
			gumi.Tool.MarginMinRegular(4, gumi.MTButton0("Reset", func(self *gumi.MTButton) {
				for _, v := range VerticalProgress {
					v.Childrun()[0].(*gumi.MTProgress).Set(0)
				}
			})),
			gumi.Tool.MarginMinRegular(4, gumi.MTButton0("Run", func(self *gumi.MTButton) {
				for i, v := range VerticalProgress {
					go func(index int, g gumi.GUMI) {
						t := time.NewTimer(200 * time.Millisecond * time.Duration(index))
						defer t.Stop()
						<-t.C
						g.Childrun()[0].(*gumi.MTProgress).Set(1)
					}(i, v)
				}
			})),
			gumi.LinkingFrom(
				gumi.NSize0(gumi.Size{Horizontal: gumi.AUTOLENGTH, Vertical: gumi.MinLength(180)}),
				gumi.LHorizontal1(VerticalProgress...),
			),
			//
			gumi.Tool.MarginMinRegular(4, gumi.MTButton0("Modal", func(self *gumi.MTButton) { modal.SetShow(!modal.GetShow()) })),
			gumi.Tool.MarginMinRegular(4, gumi.MTDropbox3(func(self *gumi.MTDropbox, selected string) {
				fmt.Printf("MTDropbox %6s : %s\n", self.GetMaterialColor(), selected)
			}, DropboxElems...)),
			gumi.LinkingFrom(
				gumi.NMargin0(gumi.RegularBlank(gumi.MinLength(4))),
				gumi.MTEdit0(),
			),
			gumi.AImage0(res.ImageHexagon),
			gumi.AText1("Hello, World!", gumi.Align_CENTER),
			gumi.AText1("안녕!", gumi.Align_CENTER),
		),
	))

	// GLumi Screen Setup
	lumi.SetScreen(scr)
	err = lumi.Init(60)
	if err != nil {
		panic(err)
	}

	lumi.Loop(func(lumi *glumi.GLUMIFullScreen) error {
		if err != nil {
			return err
		}
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			err = lumi.Event.Handle(event)
			if err != nil {
				return err
			}
		}
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		return nil
	}, func(lumi *glumi.GLUMIFullScreen) error {
		sdl.GLSwapWindow(wnd)
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}

var ToggleProgress = []gumi.GUMI{
	gumi.Tool.MarginMinRegular(4,
		gumi.LHorizontal1(
			gumi.MTToggle1(gumi.Material.Pallette.White, gumi.Material.Pallette.White, func(self *gumi.MTToggle, active bool) {
				pbar := self.Parent().Childrun()[1].Childrun()[0].Childrun()[0].Childrun()[0].(*gumi.MTProgress)
				if active {
					pbar.Set(1)
				} else {
					pbar.Set(0)
				}
			}),
			gumi.LCenter0(
				gumi.LinkingFrom(
					gumi.NSize0(gumi.Size{
						Vertical:   gumi.MINLENGTH,
						Horizontal: gumi.MAXLENGTH,
					}),
					gumi.NMargin0(gumi.SymmetryBlank(gumi.MinLength(4), gumi.AUTOLENGTH)),
					gumi.MTProgress1(gumi.Material.Pallette.White, gumi.Material.Pallette.White),
				),
			),
		),
	),
	gumi.Tool.MarginMinRegular(4,
		gumi.LHorizontal1(
			gumi.MTToggle1(gumi.Material.Pallette.White, gumi.Material.Pallette.Red, func(self *gumi.MTToggle, active bool) {
				pbar := self.Parent().Childrun()[1].Childrun()[0].Childrun()[0].Childrun()[0].(*gumi.MTProgress)
				if active {
					pbar.Set(1)
				} else {
					pbar.Set(0)
				}
			}),
			gumi.LCenter0(
				gumi.LinkingFrom(
					gumi.NSize0(gumi.Size{
						Vertical:   gumi.MINLENGTH,
						Horizontal: gumi.MAXLENGTH,
					}),
					gumi.NMargin0(gumi.SymmetryBlank(gumi.MinLength(4), gumi.AUTOLENGTH)),
					gumi.MTProgress1(gumi.Material.Pallette.White, gumi.Material.Pallette.Red),
				),
			),
		),
	),
	gumi.Tool.MarginMinRegular(4,
		gumi.LHorizontal1(
			gumi.MTToggle1(gumi.Material.Pallette.White, gumi.Material.Pallette.Green, func(self *gumi.MTToggle, active bool) {
				pbar := self.Parent().Childrun()[1].Childrun()[0].Childrun()[0].Childrun()[0].(*gumi.MTProgress)
				if active {
					pbar.Set(1)
				} else {
					pbar.Set(0)
				}
			}),
			gumi.LCenter0(
				gumi.LinkingFrom(
					gumi.NSize0(gumi.Size{
						Vertical:   gumi.MINLENGTH,
						Horizontal: gumi.MAXLENGTH,
					}),
					gumi.NMargin0(gumi.SymmetryBlank(gumi.MinLength(4), gumi.AUTOLENGTH)),
					gumi.MTProgress1(gumi.Material.Pallette.White, gumi.Material.Pallette.Green),
				),
			),
		),
	),
	gumi.Tool.MarginMinRegular(4,
		gumi.LHorizontal1(
			gumi.MTToggle1(gumi.Material.Pallette.White, gumi.Material.Pallette.Blue, func(self *gumi.MTToggle, active bool) {
				pbar := self.Parent().Childrun()[1].Childrun()[0].Childrun()[0].Childrun()[0].(*gumi.MTProgress)
				if active {
					pbar.Set(1)
				} else {
					pbar.Set(0)
				}
			}),
			gumi.LCenter0(
				gumi.LinkingFrom(
					gumi.NSize0(gumi.Size{
						Vertical:   gumi.MINLENGTH,
						Horizontal: gumi.MAXLENGTH,
					}),
					gumi.NMargin0(gumi.SymmetryBlank(gumi.MinLength(4), gumi.AUTOLENGTH)),
					gumi.MTProgress1(gumi.Material.Pallette.White, gumi.Material.Pallette.Blue),
				),
			),
		),
	),
	gumi.Tool.MarginMinRegular(4,
		gumi.LHorizontal1(
			gumi.MTToggle1(gumi.Material.Pallette.White, gumi.Material.Pallette.Yellow, func(self *gumi.MTToggle, active bool) {
				pbar := self.Parent().Childrun()[1].Childrun()[0].Childrun()[0].Childrun()[0].(*gumi.MTProgress)
				if active {
					pbar.Set(1)
				} else {
					pbar.Set(0)
				}
			}),
			gumi.LCenter0(
				gumi.LinkingFrom(
					gumi.NSize0(gumi.Size{
						Vertical:   gumi.MINLENGTH,
						Horizontal: gumi.MAXLENGTH,
					}),
					gumi.NMargin0(gumi.SymmetryBlank(gumi.MinLength(4), gumi.AUTOLENGTH)),
					gumi.MTProgress1(gumi.Material.Pallette.White, gumi.Material.Pallette.Yellow),
				),
			),
		),
	),
}
var Radios = []gumi.GUMI{
	gumi.Tool.MarginMinRegular(4, gumi.MTRadio1(
		gumi.Material.Pallette.White,
		gumi.Material.Pallette.White,
		func(self *gumi.MTRadio, active bool) {
			fmt.Printf("MTRadio %6s : %v\n", self.GetToMaterialColor(), active)
		},
	)),
	gumi.Tool.MarginMinRegular(4, gumi.MTRadio1(
		gumi.Material.Pallette.White,
		gumi.Material.Pallette.Red,
		func(self *gumi.MTRadio, active bool) {
			fmt.Printf("MTRadio %6s : %v\n", self.GetToMaterialColor(), active)
		},
	)),
	gumi.Tool.MarginMinRegular(4, gumi.MTRadio1(
		gumi.Material.Pallette.White,
		gumi.Material.Pallette.Green,
		func(self *gumi.MTRadio, active bool) {
			fmt.Printf("MTRadio %6s : %v\n", self.GetToMaterialColor(), active)
		},
	)),
	gumi.Tool.MarginMinRegular(4, gumi.MTRadio1(
		gumi.Material.Pallette.White,
		gumi.Material.Pallette.Blue,
		func(self *gumi.MTRadio, active bool) {
			fmt.Printf("MTRadio %6s : %v\n", self.GetToMaterialColor(), active)
		},
	)),
	gumi.Tool.MarginMinRegular(4, gumi.MTRadio1(
		gumi.Material.Pallette.White,
		gumi.Material.Pallette.Yellow,
		func(self *gumi.MTRadio, active bool) {
			fmt.Printf("MTRadio %6s : %v\n", self.GetToMaterialColor(), active)
		},
	)),
}
var VerticalProgress = []gumi.GUMI{
	gumi.Tool.MarginMinRegular(4, gumi.MTProgress2(gumi.Material.Pallette.White, gumi.Material.Pallette.White, gumi.AxisVertical)),
	gumi.Tool.MarginMinRegular(4, gumi.MTProgress2(gumi.Material.Pallette.White, gumi.Material.Pallette.Red, gumi.AxisVertical)),
	gumi.Tool.MarginMinRegular(4, gumi.MTProgress2(gumi.Material.Pallette.White, gumi.Material.Pallette.Green, gumi.AxisVertical)),
	gumi.Tool.MarginMinRegular(4, gumi.MTProgress2(gumi.Material.Pallette.White, gumi.Material.Pallette.Blue, gumi.AxisVertical)),
	gumi.Tool.MarginMinRegular(4, gumi.MTProgress2(gumi.Material.Pallette.White, gumi.Material.Pallette.Yellow, gumi.AxisVertical)),
}
var DropboxElems = []string{
	"Hello 0", "Hello 1", "Hello 2", "Hello 3", "Hello 4",
	"Hello 5", "Hello 6", "Hello 7", "Hello 8", "Hello 9",
	"Hello 10", "Hello 11", "Hello 12", "Hello 13", "Hello 14",
	"Hello 15", "Hello 16", "Hello 17", "Hello 18", "Hello 19",
}

func GoRuntimeInit() {
	runtime.LockOSThread()
	runtime.GOMAXPROCS(runtime.NumCPU())
}
func GLInit() {
	if err := gl.Init(); err != nil {
		panic(err)
	}
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.ClearColor(1.0, 1.0, 1.0, 1.0)
}
func SDL2Init() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		panic(err)
	}
}
func SDL2Window(w, h int) (*sdl.Window, sdl.GLContext) {
	var disp sdl.DisplayMode
	err := sdl.GetDesktopDisplayMode(0, &disp)
	if err != nil {
		panic(err)
	}
	var windW, windH int32
	if int32(w) > disp.W {
		windW = disp.W
	} else {
		windW = int32(w)
	}
	if int32(h) > disp.H {
		windH = disp.H
	} else {
		windH = int32(h)
	}
	wnd, err := sdl.CreateWindow("GUMI", (disp.W-windW)/2, (disp.H-windH)/2, windW, windH,
		sdl.WINDOW_ALLOW_HIGHDPI | sdl.WINDOW_BORDERLESS | sdl.WINDOW_OPENGL,
	)
	if err != nil {
		panic(err)
	}
	ctx, err := sdl.GLCreateContext(wnd)
	if err != nil {
		panic(err)
	}
	return wnd, ctx
}
func GUMIInit() {
	f, err := truetype.Parse(res.NanumSquareRoundR)
	if err != nil {
		panic(err)
	}
	gumi.DefaultStyleFont(f, 12)
}
func GLUMIInit() {
	err := glumi.Init()
	if err != nil {
		panic(err)
	}
}
