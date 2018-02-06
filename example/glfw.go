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
	"time"
)

const HISTORY = 16

type GLumiTest struct {
	updateSum   float64
	updateCount float64
	readySum    float64
	readyCount  float64
	drawSum     float64
	drawCount   float64
}

func (s *GLumiTest) Update(cb func()) {
	start := time.Now()
	cb()
	end := time.Now()
	//
	s.updateSum += end.Sub(start).Seconds()
	s.updateCount += 1
}
func (s *GLumiTest) Ready(cb func()) {
	start := time.Now()
	cb()
	end := time.Now()
	//
	s.readySum += end.Sub(start).Seconds()
	s.readyCount += 1
}
func (s *GLumiTest) Draw(cb func()) {
	start := time.Now()
	cb()
	end := time.Now()
	//
	s.drawSum += end.Sub(start).Seconds()
	s.drawCount += 1
}

func (s *GLumiTest) Print() {
	fmt.Printf(
		"Update : %10.3fms, Ready : %10.3fms, Draw : %10.3fms",
		s.updateSum/s.updateCount*1000,
		s.readySum/s.readyCount*1000,
		s.drawSum/s.drawCount*1000,
	)
}
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
	if width > vidmod.Width{
		windW = vidmod.Width
	}else {
		windW = width
	}
	if height > vidmod.Height{
		windH = vidmod.Height
	}else {
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
	glm := glumi.NewGLUMI(nil)

	// window build
	window.SetKeyCallback(glm.Event.DirectKey)
	window.SetCursorPosCallback(func(w *glfw.Window, xpos float64, ypos float64) {
		glm.Event.Cursor(xpos/float64(windW) * float64(width), ypos/float64(windH) * float64(height))
	})
	window.SetMouseButtonCallback(glm.Event.DirectMouseButton)
	window.SetCharCallback(glm.Event.DirectRune)
	window.SetScrollCallback(glm.Event.DirectScroll)
	//
	modal := gumi.ALModal0(
		gumi.LCenter0(

			gumi.MTButton1("Modal, Hello!", func(self *gumi.MTButton) {
				self.Parent().(*gumi.ALModal).SetShow(!self.Parent().(*gumi.ALModal).GetShow())
			}),
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
			gumi.Tool.MarginMinRegular(4, gumi.MTButton0(gumi.Red, "Close", func(self *gumi.MTButton) {
				window.SetShouldClose(true)
			})),
			gumi.LHorizontal1(toggles...),
			gumi.LHorizontal1(radios...),
			gumi.Tool.MarginMinRegular(4, gumi.MTButton1("Modal", func(self *gumi.MTButton) {
				modal.SetShow(!modal.GetShow())
			})),
			gumi.Tool.MarginMinRegular(4, gumi.MTDropbox3(func(self *gumi.MTDropbox, selected string) {
				fmt.Printf("MTDropbox %6s : %s\n", self.GetMaterialColor(), selected)
			}, "Hello 0", "Hello 1", "Hello 2", "Hello 3", "Hello 4")),
			//gumi.ASpacer2(gumi.MinLength(50)),
			gumi.Tool.MarginMinRegular(4, gumi.MTButton1("Reset", func(self *gumi.MTButton) {
				for _, v := range progresses {
					v.Childrun()[0].(*gumi.MTProgress).Set(0)
				}
			})),
			gumi.Tool.MarginMinRegular(4, gumi.MTButton1("Activate", func(self *gumi.MTButton) {
				for i, v := range progresses {
					v.Childrun()[0].(*gumi.MTProgress).Set(float64(i+1) / 5)
				}
			})),
			gumi.ASpacer2(gumi.MinLength(12)),
			gumi.LVertical1(progresses...),
			gumi.ASpacer2(gumi.MinLength(12)),
			gumi.LinkingFrom(
				gumi.NMargin0(gumi.RegularBlank(gumi.MinLength(4))),
				gumi.MTEdit3(),
			),
			gumi.AText0("Hello, World!", gumi.Align_CENTER),
		),
	))
	scr.Init()
	// GLumi Screen Setup
	glm.SetScreen(scr)
	err = glm.Init()
	if err != nil {
		panic(err)
	}
	glfw.SwapInterval(0)
	// Configure global settings
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.ClearColor(1.0, 1.0, 1.0, 1.0)
	previousTime := glfw.GetTime()
	//
	test := GLumiTest{}
	c := make(chan int)
	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		// Update
		t := glfw.GetTime()
		elapsed := t - previousTime
		previousTime = t
		//
		go func() {
			test.Update(func() {
				scr.Update(&gumi.Information{
					Dt: int64(elapsed * 1000),
				}, nil)
			})
			test.Ready(func() {
				scr.Ready()
			})
			test.Draw(func() {
				scr.Draw()
			})
			c <- 1
		}()

		<-c

		glm.Render.Upload()
		glm.Render.Draw()
		window.SwapBuffers()
		glfw.PollEvents()
	}
	test.Print()
}

var progresses = []gumi.GUMI{
	gumi.Tool.MarginMinRegular(4, gumi.MTProgress0(
		gumi.White,
		gumi.White,
		0,
	)),
	gumi.Tool.MarginMinRegular(4, gumi.MTProgress0(
		gumi.White,
		gumi.Red,
		0,
	)),
	gumi.Tool.MarginMinRegular(4, gumi.MTProgress0(
		gumi.White,
		gumi.Green,
		0,
	)),
	gumi.Tool.MarginMinRegular(4, gumi.MTProgress0(
		gumi.White,
		gumi.Blue,
		0,
	)),
	gumi.Tool.MarginMinRegular(4, gumi.MTProgress0(
		gumi.White,
		gumi.Yellow,
		0,
	)),
}
var radios = []gumi.GUMI{
	gumi.Tool.MarginMinRegular(4, gumi.MTRadio0(
		gumi.White,
		gumi.White,
		func(self *gumi.MTRadio, active bool) {
			fmt.Printf("MTRadio %6s : %v\n", self.GetToMaterialColor(), active)
		},
	)),
	gumi.Tool.MarginMinRegular(4, gumi.MTRadio0(
		gumi.White,
		gumi.Red,
		func(self *gumi.MTRadio, active bool) {
			fmt.Printf("MTRadio %6s : %v\n", self.GetToMaterialColor(), active)
		},
	)),
	gumi.Tool.MarginMinRegular(4, gumi.MTRadio0(
		gumi.White,
		gumi.Green,
		func(self *gumi.MTRadio, active bool) {
			fmt.Printf("MTRadio %6s : %v\n", self.GetToMaterialColor(), active)
		},
	)),
	gumi.Tool.MarginMinRegular(4, gumi.MTRadio0(
		gumi.White,
		gumi.Blue,
		func(self *gumi.MTRadio, active bool) {
			fmt.Printf("MTRadio %6s : %v\n", self.GetToMaterialColor(), active)
		},
	)),
	gumi.Tool.MarginMinRegular(4, gumi.MTRadio0(
		gumi.White,
		gumi.Yellow,
		func(self *gumi.MTRadio, active bool) {
			fmt.Printf("MTRadio %6s : %v\n", self.GetToMaterialColor(), active)
		},
	)),
}
var toggles = []gumi.GUMI{
	gumi.Tool.MarginMinRegular(4, gumi.MTToggle0(
		gumi.White,
		gumi.White,
		func(self *gumi.MTToggle, active bool) {
			fmt.Printf("MTToggle %6s : %v\n", self.GetToMaterialColor(), active)
		},
	)),
	gumi.Tool.MarginMinRegular(4, gumi.MTToggle0(
		gumi.White,
		gumi.Red,
		func(self *gumi.MTToggle, active bool) {
			fmt.Printf("MTToggle %6s : %v\n", self.GetToMaterialColor(), active)
		},
	)),
	gumi.Tool.MarginMinRegular(4, gumi.MTToggle0(
		gumi.White,
		gumi.Green,
		func(self *gumi.MTToggle, active bool) {
			fmt.Printf("MTToggle %6s : %v\n", self.GetToMaterialColor(), active)
		},
	)),
	gumi.Tool.MarginMinRegular(4, gumi.MTToggle0(
		gumi.White,
		gumi.Blue,
		func(self *gumi.MTToggle, active bool) {
			fmt.Printf("MTToggle %6s : %v\n", self.GetToMaterialColor(), active)
		},
	)),
	gumi.Tool.MarginMinRegular(4, gumi.MTToggle0(
		gumi.White,
		gumi.Yellow,
		func(self *gumi.MTToggle, active bool) {
			fmt.Printf("MTToggle %6s : %v\n", self.GetToMaterialColor(), active)
		},
	)),
}

func GLFWHint() {

	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.Decorated, glfw.False)
}
