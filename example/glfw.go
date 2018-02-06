package main

import (
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/iamGreedy/gumi"
	"github.com/iamGreedy/gumi/glumi"
	"github.com/iamGreedy/gumi/gutl"
	"log"
	"runtime"
)
func main() {
	var width, height = gutl.DefinedResolutions.Get("HD")
	var err error

	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	runtime.GOMAXPROCS(runtime.NumCPU())
	// Init GLFW
	if err = glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	defer glfw.Terminate()
	//
	var vidmod = glfw.GetPrimaryMonitor().GetVideoMode()
	GLFWHint()
	var windW, windH int
	if width > vidmod.Width {
		windW = vidmod.Width
	} else {
		windW = width
	}
	if height > vidmod.Height {
		windH = vidmod.Height
	} else {
		windH = height
	}
	window, err := glfw.CreateWindow(windW, windH, "Cube", nil, nil)
	if err != nil {
		panic(err)
	}
	window.SetPos((vidmod.Width-windW)/2, (vidmod.Height-windH)/2)
	window.MakeContextCurrent()
	// Init GL
	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version : ", version)

	// Init GLumi
	err = glumi.DefaultShader.Load()
	if err != nil {
		panic(err)
	}
	// GLumi Object allocate
	lumi := glumi.NewGLUMI()

	// window build
	window.SetKeyCallback(lumi.Event.DirectKey)
	window.SetCursorPosCallback(func(w *glfw.Window, xpos float64, ypos float64) {
		lumi.Event.Cursor(xpos/float64(windW)*float64(width), ypos/float64(windH)*float64(height))
	})
	window.SetMouseButtonCallback(lumi.Event.DirectMouseButton)
	window.SetCharCallback(lumi.Event.DirectRune)
	window.SetScrollCallback(lumi.Event.DirectScroll)
	//
	modal := gumi.ALModal0(
		gumi.LCenter0(
			gumi.NMinimum0(gumi.AxisAll,
				gumi.MTButton0("Modal, Hello!", func(self *gumi.MTButton) {
					p := self.Parent().Parent().Parent()
					p.(*gumi.ALModal).SetShow(!p.(*gumi.ALModal).GetShow())
				}),
			),
		),
	)
	scr := gumi.NewScreen(width, height)
	scr.Root(gumi.LinkingFrom(
		modal,
		gumi.NBackground0(),
		gumi.NDrawing1(
			gumi.Drawing.FPS(),
			gumi.Drawing.Ruler.Hint.Vertical(100),
			gumi.Drawing.Ruler.Hint.Horizontal(100),
		),
		gumi.NMargin0(gumi.RegularBlank(gumi.MinLength(20))),
		gumi.LVertical1(
			gumi.Tool.MarginMinRegular(4, gumi.MTButton1(gumi.Material.Pallette.Red, "Close", func(self *gumi.MTButton) {
				window.SetShouldClose(true)
			})),
			gumi.LHorizontal1(toggles...),
			gumi.LHorizontal1(radios...),
			gumi.Tool.MarginMinRegular(4, gumi.MTButton0("Modal", func(self *gumi.MTButton) {
				modal.SetShow(!modal.GetShow())
			})),
			gumi.Tool.MarginMinRegular(4, gumi.MTDropbox3(func(self *gumi.MTDropbox, selected string) {
				fmt.Printf("MTDropbox %6s : %s\n", self.GetMaterialColor(), selected)
			}, "Hello 0", "Hello 1", "Hello 2", "Hello 3", "Hello 4")),
			//gumi.ASpacer2(gumi.MinLength(50)),
			gumi.Tool.MarginMinRegular(4, gumi.MTButton0("Reset", func(self *gumi.MTButton) {
				for _, v := range progresses {
					v.Childrun()[0].(*gumi.MTProgress).Set(0)
				}
			})),
			gumi.Tool.MarginMinRegular(4, gumi.MTButton0("Activate", func(self *gumi.MTButton) {
				for i, v := range progresses {
					v.Childrun()[0].(*gumi.MTProgress).Set(float64(i+1) / 5)
				}
			})),
			gumi.ASpacer2(gumi.MinLength(12)),
			gumi.LVertical1(progresses...),
			gumi.ASpacer2(gumi.MinLength(12)),
			gumi.LinkingFrom(
				gumi.NMargin0(gumi.RegularBlank(gumi.MinLength(4))),
				gumi.MTEdit0(),
			),
			gumi.AText0("Hello, World!", gumi.Align_CENTER),
		),
	))
	scr.Init()
	// GLumi Screen Setup
	lumi.SetScreen(scr)
	err = lumi.Init(60)
	if err != nil {
		panic(err)
	}
	// Configure global settings
	glfw.SwapInterval(0)
	GLInit()
	lumi.Loop(func(lumi *glumi.GLUMIFullScreen) error {
		if window.ShouldClose(){
			return glumi.Stop
		}
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		return nil
	}, func(lumi *glumi.GLUMIFullScreen) error {
		window.SwapBuffers()
		glfw.PollEvents()
		return nil
	})
}

var progresses = []gumi.GUMI{
	gumi.Tool.MarginMinRegular(4, gumi.MTProgress0(
		gumi.Material.Pallette.White,
		gumi.Material.Pallette.White,
	)),
	gumi.Tool.MarginMinRegular(4, gumi.MTProgress0(
		gumi.Material.Pallette.White,
		gumi.Material.Pallette.Red,
	)),
	gumi.Tool.MarginMinRegular(4, gumi.MTProgress0(
		gumi.Material.Pallette.White,
		gumi.Material.Pallette.Green,
	)),
	gumi.Tool.MarginMinRegular(4, gumi.MTProgress0(
		gumi.Material.Pallette.White,
		gumi.Material.Pallette.Blue,
	)),
	gumi.Tool.MarginMinRegular(4, gumi.MTProgress0(
		gumi.Material.Pallette.White,
		gumi.Material.Pallette.Yellow,
	)),
}
var radios = []gumi.GUMI{
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
var toggles = []gumi.GUMI{
	gumi.Tool.MarginMinRegular(4, gumi.MTToggle1(
		gumi.Material.Pallette.White,
		gumi.Material.Pallette.White,
		func(self *gumi.MTToggle, active bool) {
			fmt.Printf("MTToggle %6s : %v\n", self.GetToMaterialColor(), active)
		},
	)),
	gumi.Tool.MarginMinRegular(4, gumi.MTToggle1(
		gumi.Material.Pallette.White,
		gumi.Material.Pallette.Red,
		func(self *gumi.MTToggle, active bool) {
			fmt.Printf("MTToggle %6s : %v\n", self.GetToMaterialColor(), active)
		},
	)),
	gumi.Tool.MarginMinRegular(4, gumi.MTToggle1(
		gumi.Material.Pallette.White,
		gumi.Material.Pallette.Green,
		func(self *gumi.MTToggle, active bool) {
			fmt.Printf("MTToggle %6s : %v\n", self.GetToMaterialColor(), active)
		},
	)),
	gumi.Tool.MarginMinRegular(4, gumi.MTToggle1(
		gumi.Material.Pallette.White,
		gumi.Material.Pallette.Blue,
		func(self *gumi.MTToggle, active bool) {
			fmt.Printf("MTToggle %6s : %v\n", self.GetToMaterialColor(), active)
		},
	)),
	gumi.Tool.MarginMinRegular(4, gumi.MTToggle1(
		gumi.Material.Pallette.White,
		gumi.Material.Pallette.Yellow,
		func(self *gumi.MTToggle, active bool) {
			fmt.Printf("MTToggle %6s : %v\n", self.GetToMaterialColor(), active)
		},
	)),
}

func GLInit()  {
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.ClearColor(1.0, 1.0, 1.0, 1.0)
}
func GLFWHint() {

	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.Decorated, glfw.False)

}
