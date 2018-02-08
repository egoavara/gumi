package glumi

import (
	"github.com/iamGreedy/gumi"
	"time"
)

type GLUMIFullScreen struct {
	screen *gumi.Screen
	Event  Handler
	Render GLRender
	fps FPS
	//
	updateCount uint64
	readyCount uint64
	drawCount uint64
}

func NewGLUMI() *GLUMIFullScreen {
	temp := &GLUMIFullScreen{}
	temp.Event = Handler{glumi: temp, keymap:make(map[gumi.GUMIKey]gumi.EventKind)}
	temp.Render = GLRender{glumi:temp}
	return temp
}
func (s *GLUMIFullScreen) Init(fps int) error {
	err := s.Render.init()
	if err != nil {
		return err
	}
	s.screen.Init()
	s.fps.SetInterval((1000 * time.Millisecond) / time.Duration(fps) / 8)
	return nil
}
func (s *GLUMIFullScreen) SetScreen(screen *gumi.Screen) {
	s.screen = screen
}
func (s *GLUMIFullScreen) GetScreen() *gumi.Screen {
	return s.screen
}

func (s *GLUMIFullScreen) Loop(fnBefore, fnAfter func(lumi *GLUMIFullScreen) error) (err error) {
	s.fps.Start()
	defer s.fps.Stop()
	var prev, curr time.Time

	prev = s.fps.Wait()
	for loopcount:=0;true;loopcount++{
		curr = s.fps.Wait()
		err = fnBefore(s)
		if err != nil{
			break
		}
		// GUMI
		s.screen.Update(&gumi.Information{
			Dt: int64(curr.Sub(prev).Seconds() * 1000),
		}, nil)
		if loopcount % 8 == 0{
			s.screen.Ready()
		}
		s.screen.Draw()
		// GLFW
		s.Render.Upload()
		s.Render.Draw()
		err = fnAfter(s)
		if err != nil{
			break
		}
		prev = curr
	}
	if err == Stop{
		return nil
	}
	return err
}