package main

import (
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/golang/freetype/truetype"
	"github.com/iamGreedy/gumi"
	"github.com/iamGreedy/gumi/glumi"
	"github.com/veandco/go-sdl2/sdl"
	"runtime"
	"github.com/iamGreedy/gumi/example/sdl2example"
	"github.com/iamGreedy/gumi/example/sdl2example/asset"
	"github.com/iamGreedy/gumi/gcore"
)

func main() {
	var width, height = gumi.DefinedResolutions.Get("VGA")
	var err error

	// init go:runtime
	GoRuntimeInit()
	// Init SDL2
	SDL2Init()
	wnd := SDL2Window(width, height)
	//suf, err := wnd.GetSurface()
	//fmt.Println(wnd.GetSize())
	//fmt.Println(suf.W, suf.H)
	//wnd.SetBordered(false)
	ctx, err := sdl.GLCreateContext(wnd)
	gcore.Assert(err)
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
	scr := gumi.NewScreen(width, height)
	scr.Root(sdl2example.HelloImage)

	// GLumi Screen Setup
	lumi.SetScreen(scr)
	err = lumi.Init(0)
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
			//if v, ok := event.(*sdl.MouseMotionEvent); ok{
			//}

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
	gcore.Assert(sdl.Init(sdl.INIT_EVERYTHING))
	sdl.GLSetAttribute(sdl.GL_CONTEXT_MAJOR_VERSION, 4)
	sdl.GLSetAttribute(sdl.GL_CONTEXT_MINOR_VERSION, 1)
	sdl.GLSetAttribute(sdl.GL_CONTEXT_FORWARD_COMPATIBLE_FLAG, 1)
	sdl.GLSetAttribute(sdl.GL_CONTEXT_PROFILE_MASK, sdl.GL_CONTEXT_PROFILE_CORE)
	sdl.GLSetAttribute(sdl.GL_DOUBLEBUFFER, 1)
}
func SDL2Window(w, h int) *sdl.Window {
	var disp sdl.DisplayMode
	gcore.Assert(sdl.GetDesktopDisplayMode(0, &disp))
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
	wnd, err := sdl.CreateWindow("GUMI", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, windW, windH,
		sdl.WINDOW_OPENGL|sdl.WINDOW_ALWAYS_ON_TOP,
	)
	gcore.Assert(err)
	return wnd
}
func GUMIInit() {
	f, err := truetype.Parse(asset.MustAsset("NanumSquareRoundR.ttf"))
	if err != nil {
		panic(err)
	}
	gumi.ModifyDefaultStyle(f, 12)
}
func GLUMIInit() {
	err := glumi.Init()
	if err != nil {
		panic(err)
	}
}
