package glumi

import "github.com/iamGreedy/gumi"

type GLUMI struct {
	screen *gumi.Screen
	Event  GLFWEvent
	Render GLRender
}

func NewGLUMI(screen *gumi.Screen) *GLUMI {
	temp := &GLUMI{}
	temp.screen = screen
	temp.Event = GLFWEvent{glumi: temp}
	temp.Render = GLRender{glumi:temp}
	return temp
}
func (s *GLUMI) Init() error {
	err := s.Render.init()
	if err != nil {
		return err
	}
	return nil
}
func (s *GLUMI) SetScreen(screen *gumi.Screen) {
	s.screen = screen
}
func (s *GLUMI) GetScreen() *gumi.Screen {
	return s.screen
}

