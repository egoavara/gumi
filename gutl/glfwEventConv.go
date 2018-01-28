package gutl

import (
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/iamGreedy/gumi"
)

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
		convkey = gumi.KEY_AL
	case glfw.KeyLeftSuper:
	case glfw.KeyRightShift:
	case glfw.KeyRightControl:
	case glfw.KeyRightAlt:
	case glfw.KeyRightSuper:
	case glfw.KeyMenu:
	case glfw.KeyLast:
	}
}
