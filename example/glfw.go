package main

import (
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/iamGreedy/gumi"
	"image"
	"log"
	"runtime"
	"strings"
)

const windowWidth = 800
const windowHeight = 600

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

func main() {
	scr := gumi.NewScreen(windowWidth, windowHeight)
	scr.Root(gumi.LinkingFrom(
		gumi.NBackground0(),
		gumi.NDrawing1(
			gumi.Drawing.Ruler.Hint.Vertical(100),
			gumi.Drawing.Ruler.Hint.Horizontal(100),
		),
		gumi.NMargin0(gumi.RegularBlank(gumi.MinLength(20))),
		gumi.NVertical1(
			toggles,
			radios,
			gumi.LinkingFrom(
				gumi.NMargin0(gumi.RegularBlank(gumi.MinLength(4))),
				gumi.MTButton1("Reset", func() {
					for _, v := range progresses.Childrun() {
						v.Childrun()[0].(*gumi.MTProgress).Set(0)
					}
				}),
			),
			gumi.LinkingFrom(
				gumi.NMargin0(gumi.RegularBlank(gumi.MinLength(4))),
				gumi.MTButton1("Activate", func() {
					for i, v := range progresses.Childrun() {
						v.Childrun()[0].(*gumi.MTProgress).Set(float64(i+1) / 5)
					}
				}),
			),
			progresses,
			gumi.AText0("Hello, World!", gumi.Align_CENTER),
		),
	))
	scr.Update(nil, nil)
	scr.Ready()

	//
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	window, err := glfw.CreateWindow(windowWidth, windowHeight, "Cube", nil, nil)
	if err != nil {
		panic(err)
	}
	window.SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		scr.Event(ConvertGLFWKey(key, action))
	})
	window.SetCursorPosCallback(func(w *glfw.Window, xpos float64, ypos float64) {
		scr.Event(ConvertGLFWCursor(xpos, ypos))
	})
	window.SetMouseButtonCallback(func(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {
		scr.Event(ConvertGLFWMouseButton(button, action))
	})
	window.SetCharCallback(func(w *glfw.Window, char rune) {

	})
	window.SetScrollCallback(func(w *glfw.Window, xoff float64, yoff float64) {

	})
	window.MakeContextCurrent()

	// Initialize Glow
	if err := gl.Init(); err != nil {
		panic(err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version", version)

	// Configure the vertex and fragment shaders
	program, err := newProgram(vertexShader, fragmentShader)
	if err != nil {
		panic(err)
	}

	gl.UseProgram(program)
	tex := newTexture()

	// Configure the vertex data
	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(screenVertice)*4*5, gl.Ptr(screenVertice), gl.STATIC_DRAW)

	vertAttrib := uint32(gl.GetAttribLocation(program, gl.Str("vert\x00")))
	gl.EnableVertexAttribArray(vertAttrib)
	gl.VertexAttribPointer(vertAttrib, 3, gl.FLOAT, false, 5*4, gl.PtrOffset(0))

	texCoordAttrib := uint32(gl.GetAttribLocation(program, gl.Str("vertTexCoord\x00")))
	gl.EnableVertexAttribArray(texCoordAttrib)
	gl.VertexAttribPointer(texCoordAttrib, 2, gl.FLOAT, false, 5*4, gl.PtrOffset(3*4))

	// Configure global settings
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.ClearColor(1.0, 1.0, 1.0, 1.0)

	previousTime := glfw.GetTime()
	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		// Update
		time := glfw.GetTime()
		elapsed := time - previousTime
		previousTime = time
		scr.Update(&gumi.Information{
			Dt: int64(elapsed * 1000),
		}, nil)
		scr.Ready()
		scr.Draw()
		setTexture(tex, scr.RGBA())
		// Render
		gl.UseProgram(program)

		gl.BindVertexArray(vao)
		gl.ActiveTexture(gl.TEXTURE0)
		gl.BindTexture(gl.TEXTURE_2D, tex)

		gl.DrawArrays(gl.TRIANGLES, 0, 6*2*3)

		// Maintenance
		window.SwapBuffers()
		glfw.PollEvents()
	}
}

var progresses = gumi.NVertical1(
	gumi.LinkingFrom(
		gumi.NMargin0(gumi.RegularBlank(gumi.MinLength(4))),
		gumi.MTProgress0(
			gumi.White,
			gumi.White,
			0,
		),
	),
	gumi.LinkingFrom(
		gumi.NMargin0(gumi.RegularBlank(gumi.MinLength(4))),
		gumi.MTProgress0(
			gumi.White,
			gumi.Red,
			0,
		),
	),
	gumi.LinkingFrom(
		gumi.NMargin0(gumi.RegularBlank(gumi.MinLength(4))),
		gumi.MTProgress0(
			gumi.White,
			gumi.Blue,
			0,
		),
	),
	gumi.LinkingFrom(
		gumi.NMargin0(gumi.RegularBlank(gumi.MinLength(4))),
		gumi.MTProgress0(
			gumi.White,
			gumi.Green,
			0,
		),
	),
	gumi.LinkingFrom(
		gumi.NMargin0(gumi.RegularBlank(gumi.MinLength(4))),
		gumi.MTProgress0(
			gumi.White,
			gumi.Yellow,
			0,
		),
	),
)
var radios = gumi.NHorizontal1(
	gumi.LinkingFrom(
		gumi.NMargin0(gumi.RegularBlank(gumi.MinLength(4))),
		gumi.MTRadio0(
			gumi.White,
			gumi.White,
			func(active bool) {
				fmt.Printf("MTRadio %6s : %v\n", "White", active)
			},
		),
	),
	gumi.LinkingFrom(
		gumi.NMargin0(gumi.RegularBlank(gumi.MinLength(4))),
		gumi.MTRadio0(
			gumi.White,
			gumi.Red,
			func(active bool) {
				fmt.Printf("MTRadio %6s : %v\n", "Red", active)
			},
		),
	),
	gumi.LinkingFrom(
		gumi.NMargin0(gumi.RegularBlank(gumi.MinLength(4))),
		gumi.MTRadio0(
			gumi.White,
			gumi.Blue,
			func(active bool) {
				fmt.Printf("MTRadio %6s : %v\n", "Blue", active)
			},
		),
	),
	gumi.LinkingFrom(
		gumi.NMargin0(gumi.RegularBlank(gumi.MinLength(4))),
		gumi.MTRadio0(
			gumi.White,
			gumi.Green,
			func(active bool) {
				fmt.Printf("MTRadio %6s : %v\n", "Green", active)
			},
		),
	),
	gumi.LinkingFrom(
		gumi.NMargin0(gumi.RegularBlank(gumi.MinLength(4))),
		gumi.MTRadio0(
			gumi.White,
			gumi.Yellow,
			func(active bool) {
				fmt.Printf("MTRadio %6s : %v\n", "Yellow", active)
			},
		),
	),
)
var toggles = gumi.NHorizontal1(
	gumi.LinkingFrom(
		gumi.NMargin0(gumi.RegularBlank(gumi.MinLength(4))),
		gumi.MTToggle0(
			gumi.White,
			gumi.White,
			func(active bool) {
				fmt.Printf("MTToggle %6s : %v\n", "White", active)
			},
		),
	),
	gumi.LinkingFrom(
		gumi.NMargin0(gumi.RegularBlank(gumi.MinLength(4))),
		gumi.MTToggle0(
			gumi.White,
			gumi.Red,
			func(active bool) {
				fmt.Printf("MTToggle %6s : %v\n", "Red", active)
			},
		),
	),
	gumi.LinkingFrom(
		gumi.NMargin0(gumi.RegularBlank(gumi.MinLength(4))),
		gumi.MTToggle0(
			gumi.White,
			gumi.Blue,
			func(active bool) {
				fmt.Printf("MTToggle %6s : %v\n", "Blue", active)
			},
		),
	),
	gumi.LinkingFrom(
		gumi.NMargin0(gumi.RegularBlank(gumi.MinLength(4))),
		gumi.MTToggle0(
			gumi.White,
			gumi.Green,
			func(active bool) {
				fmt.Printf("MTToggle %6s : %v\n", "Green", active)
			},
		),
	),
	gumi.LinkingFrom(
		gumi.NMargin0(gumi.RegularBlank(gumi.MinLength(4))),
		gumi.MTToggle0(
			gumi.White,
			gumi.Yellow,
			func(active bool) {
				fmt.Printf("MTToggle %6s : %v\n", "Yellow", active)
			},
		),
	),
)

//

func newProgram(vertexShaderSource, fragmentShaderSource string) (uint32, error) {
	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		return 0, err
	}

	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		return 0, err
	}

	program := gl.CreateProgram()

	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)

	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to link program: %v", log)
	}

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	return program, nil
}
func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}
func newTexture() uint32 {
	var texture uint32
	gl.GenTextures(1, &texture)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	return texture
}
func setTexture(texture uint32, img *image.RGBA) {
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(img.Rect.Size().X),
		int32(img.Rect.Size().Y),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(img.Pix))
}
func ConvertGLFWKey(key glfw.Key, action glfw.Action) gumi.Event {
	var convkey uint8
	switch key {
	case glfw.KeyUnknown:
		convkey = gumi.KEY_UNKNOWN
	case glfw.KeySpace:
		convkey = gumi.KEY_SPACE
	case glfw.KeyApostrophe:
		convkey = gumi.KEY_BACKQUOTE
	case glfw.KeyComma:
		convkey = gumi.KEY_COMMA
	case glfw.KeyMinus:
		convkey = gumi.KEY_MINUS
	case glfw.KeyPeriod:
		convkey = gumi.KEY_PERIOD
	case glfw.KeySlash:
		convkey = gumi.KEY_SLASH
	case glfw.Key0:
		convkey = gumi.KEY_0
	case glfw.Key1:
		convkey = gumi.KEY_1
	case glfw.Key2:
		convkey = gumi.KEY_2
	case glfw.Key3:
		convkey = gumi.KEY_3
	case glfw.Key4:
		convkey = gumi.KEY_4
	case glfw.Key5:
		convkey = gumi.KEY_5
	case glfw.Key6:
		convkey = gumi.KEY_6
	case glfw.Key7:
		convkey = gumi.KEY_7
	case glfw.Key8:
		convkey = gumi.KEY_8
	case glfw.Key9:
		convkey = gumi.KEY_9
	case glfw.KeySemicolon:
		convkey = gumi.KEY_SEMICOLON
	case glfw.KeyEqual:
		convkey = gumi.KEY_EQUAL
	case glfw.KeyA:
		convkey = gumi.KEY_A
	case glfw.KeyB:
		convkey = gumi.KEY_B
	case glfw.KeyC:
		convkey = gumi.KEY_C
	case glfw.KeyD:
		convkey = gumi.KEY_D
	case glfw.KeyE:
		convkey = gumi.KEY_E
	case glfw.KeyF:
		convkey = gumi.KEY_F
	case glfw.KeyG:
		convkey = gumi.KEY_G
	case glfw.KeyH:
		convkey = gumi.KEY_H
	case glfw.KeyI:
		convkey = gumi.KEY_I
	case glfw.KeyJ:
		convkey = gumi.KEY_J
	case glfw.KeyK:
		convkey = gumi.KEY_K
	case glfw.KeyL:
		convkey = gumi.KEY_L
	case glfw.KeyM:
		convkey = gumi.KEY_M
	case glfw.KeyN:
		convkey = gumi.KEY_N
	case glfw.KeyO:
		convkey = gumi.KEY_O
	case glfw.KeyP:
		convkey = gumi.KEY_P
	case glfw.KeyQ:
		convkey = gumi.KEY_Q
	case glfw.KeyR:
		convkey = gumi.KEY_R
	case glfw.KeyS:
		convkey = gumi.KEY_S
	case glfw.KeyT:
		convkey = gumi.KEY_T
	case glfw.KeyU:
		convkey = gumi.KEY_U
	case glfw.KeyV:
		convkey = gumi.KEY_V
	case glfw.KeyW:
		convkey = gumi.KEY_W
	case glfw.KeyX:
		convkey = gumi.KEY_X
	case glfw.KeyY:
		convkey = gumi.KEY_Y
	case glfw.KeyZ:
		convkey = gumi.KEY_Z
	case glfw.KeyLeftBracket:
		convkey = gumi.KEY_LEFTBRACKET
	case glfw.KeyBackslash:
		convkey = gumi.KEY_BACKSLASH
	case glfw.KeyRightBracket:
		convkey = gumi.KEY_RIGHTBRACKET
	case glfw.KeyGraveAccent:
		convkey = gumi.KEY_UNKNOWN
	case glfw.KeyWorld1:
		convkey = gumi.KEY_UNKNOWN
	case glfw.KeyWorld2:
		convkey = gumi.KEY_UNKNOWN
	case glfw.KeyEscape:
		convkey = gumi.KEY_ESC
	case glfw.KeyEnter:
		convkey = gumi.KEY_ENTER
	case glfw.KeyTab:
		convkey = gumi.KEY_TAB
	case glfw.KeyBackspace:
		convkey = gumi.KEY_BACKSPACE
	case glfw.KeyInsert:
		convkey = gumi.KEY_UNKNOWN
	case glfw.KeyDelete:
		convkey = gumi.KEY_UNKNOWN
	case glfw.KeyRight:
		convkey = gumi.KEY_RIGHT
	case glfw.KeyLeft:
		convkey = gumi.KEY_LEFT
	case glfw.KeyDown:
		convkey = gumi.KEY_DOWN
	case glfw.KeyUp:
		convkey = gumi.KEY_UP
	case glfw.KeyPageUp:
		convkey = gumi.KEY_UNKNOWN
	case glfw.KeyPageDown:
		convkey = gumi.KEY_UNKNOWN
	case glfw.KeyHome:
		convkey = gumi.KEY_UNKNOWN
	case glfw.KeyEnd:
		convkey = gumi.KEY_UNKNOWN
	case glfw.KeyCapsLock:
		convkey = gumi.KEY_UNKNOWN
	case glfw.KeyScrollLock:
		convkey = gumi.KEY_UNKNOWN
	case glfw.KeyNumLock:
		convkey = gumi.KEY_UNKNOWN
	case glfw.KeyPrintScreen:
		convkey = gumi.KEY_UNKNOWN
	case glfw.KeyPause:
		convkey = gumi.KEY_UNKNOWN
	case glfw.KeyF1:
		convkey = gumi.KEY_F1
	case glfw.KeyF2:
		convkey = gumi.KEY_F2
	case glfw.KeyF3:
		convkey = gumi.KEY_F3
	case glfw.KeyF4:
		convkey = gumi.KEY_F4
	case glfw.KeyF5:
		convkey = gumi.KEY_F5
	case glfw.KeyF6:
		convkey = gumi.KEY_F6
	case glfw.KeyF7:
		convkey = gumi.KEY_F7
	case glfw.KeyF8:
		convkey = gumi.KEY_F8
	case glfw.KeyF9:
		convkey = gumi.KEY_F9
	case glfw.KeyF10:
		convkey = gumi.KEY_F10
	case glfw.KeyF11:
		convkey = gumi.KEY_F11
	case glfw.KeyF12:
		convkey = gumi.KEY_F12
	case glfw.KeyF13:
		convkey = gumi.KEY_UNKNOWN
	case glfw.KeyF14:
		convkey = gumi.KEY_UNKNOWN
	case glfw.KeyF15:
		convkey = gumi.KEY_UNKNOWN
	case glfw.KeyF16:
		convkey = gumi.KEY_UNKNOWN
	case glfw.KeyF17:
		convkey = gumi.KEY_UNKNOWN
	case glfw.KeyF18:
		convkey = gumi.KEY_UNKNOWN
	case glfw.KeyF19:
		convkey = gumi.KEY_UNKNOWN
	case glfw.KeyF20:
		convkey = gumi.KEY_UNKNOWN
	case glfw.KeyF21:
		convkey = gumi.KEY_UNKNOWN
	case glfw.KeyF22:
		convkey = gumi.KEY_UNKNOWN
	case glfw.KeyF23:
		convkey = gumi.KEY_UNKNOWN
	case glfw.KeyF24:
		convkey = gumi.KEY_UNKNOWN
	case glfw.KeyF25:
		convkey = gumi.KEY_UNKNOWN
	case glfw.KeyKP0:
		convkey = gumi.KEY_0
	case glfw.KeyKP1:
		convkey = gumi.KEY_1
	case glfw.KeyKP2:
		convkey = gumi.KEY_2
	case glfw.KeyKP3:
		convkey = gumi.KEY_3
	case glfw.KeyKP4:
		convkey = gumi.KEY_4
	case glfw.KeyKP5:
		convkey = gumi.KEY_5
	case glfw.KeyKP6:
		convkey = gumi.KEY_6
	case glfw.KeyKP7:
		convkey = gumi.KEY_7
	case glfw.KeyKP8:
		convkey = gumi.KEY_8
	case glfw.KeyKP9:
		convkey = gumi.KEY_9
	case glfw.KeyKPDecimal:
		convkey = gumi.KEY_UNKNOWN
	case glfw.KeyKPDivide:
		convkey = gumi.KEY_SLASH
	case glfw.KeyKPMultiply:
		convkey = gumi.KEY_UNKNOWN
	case glfw.KeyKPSubtract:
		convkey = gumi.KEY_MINUS
	case glfw.KeyKPAdd:
		convkey = gumi.KEY_UNKNOWN
	case glfw.KeyKPEnter:
		convkey = gumi.KEY_ENTER
	case glfw.KeyKPEqual:
		convkey = gumi.KEY_EQUAL
	case glfw.KeyLeftShift:
		convkey = gumi.KEY_SHIFT
	case glfw.KeyLeftControl:
		convkey = gumi.KEY_CONTROL
	case glfw.KeyLeftAlt:
		convkey = gumi.KEY_ALT
	case glfw.KeyLeftSuper:
		convkey = gumi.KEY_UNKNOWN
	case glfw.KeyRightShift:
		convkey = gumi.KEY_SHIFT
	case glfw.KeyRightControl:
		convkey = gumi.KEY_CONTROL
	case glfw.KeyRightAlt:
		convkey = gumi.KEY_ALT
	case glfw.KeyRightSuper:
		convkey = gumi.KEY_UNKNOWN
	case glfw.KeyMenu:
		convkey = gumi.KEY_UNKNOWN
		//case glfw.KeyLast:
		//	convkey = gumi.KEY_UNKNOWN
	}
	//
	switch action {
	case glfw.Press:
		return gumi.EventKeyPress{Key: convkey}
	case glfw.Release:
		return gumi.EventKeyRelease{Key: convkey}
	default:
		fallthrough
	case glfw.Repeat:
		return nil
	}
}

func ConvertGLFWMouseButton(button glfw.MouseButton, action glfw.Action) gumi.Event {
	var convkey uint8
	switch button {
	case glfw.MouseButton1:
		convkey = gumi.KEY_MOUSE1
	case glfw.MouseButton2:
		convkey = gumi.KEY_MOUSE2
	case glfw.MouseButton3:
		convkey = gumi.KEY_MOUSE3
	default:
		convkey = gumi.KEY_UNKNOWN
	}
	switch action {
	case glfw.Press:
		return gumi.EventKeyPress{Key: convkey}
	case glfw.Release:
		return gumi.EventKeyRelease{Key: convkey}
	default:
		fallthrough
	case glfw.Repeat:
		return nil
	}
}
func ConvertGLFWCursor(xpos, ypos float64) gumi.Event {
	return gumi.EventCursor{
		X: uint16(xpos),
		Y: uint16(ypos),
	}
}
func ConvertGLFWRune(r rune) gumi.Event {
	return gumi.EventRune{
		Rune: r,
	}
}
func ConvertGLFWScroll(xoff, yoff float64) gumi.Event {
	return gumi.EventScroll{
		X: uint16(xoff),
		Y: uint16(yoff),
	}
}

var vertexShader = `
#version 410
in vec3 vert;
in vec2 vertTexCoord;
out vec2 fragTexCoord;
void main() {
    fragTexCoord = vertTexCoord;
    gl_Position = vec4(vert, 1);
}
` + "\x00"

var fragmentShader = `
#version 410
uniform sampler2D tex;
in vec2 fragTexCoord;
out vec4 outputColor;
void main() {
    outputColor = texture(tex, fragTexCoord);
	//outputColor = vec4(0,0,0,1);
}
` + "\x00"

var screenVertice = [][5]float32{
	{-1, -1, 0, 0, 1},
	{+1, -1, 0, 1, 1},
	{-1, +1, 0, 0, 0},
	//
	{+1, -1, 0, 1, 1},
	{+1, +1, 0, 1, 0},
	{-1, +1, 0, 0, 0},
}
