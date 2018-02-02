package glumi

import (
	"runtime"
	"path/filepath"
	"os"
	"io/ioutil"
	"strings"
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
)

const DefaultVertexShader = "DefaultShader.vs.glsl"
const DefaultFlagmentShader = "DefaultShader.fs.glsl"

var DefaultShader _DefaultShader

type _DefaultShader struct{
	Source struct{Vertex, Fragment string}
	Compiled struct{Vertex, Fragment  uint32}
}

func (s *_DefaultShader) Load() error {
	_, b, _, _ := runtime.Caller(0)
	basepath   := filepath.Dir(b)
	// vertex shader source
	fvert, err := os.Open(filepath.Join(basepath, DefaultVertexShader))

	if err != nil {
		return (err)
	}
	vertSRC, err := ioutil.ReadAll(fvert)
	if err != nil {
		return (err)
	}
	// flagment shader source
	ffrag, err := os.Open(filepath.Join(basepath, DefaultFlagmentShader))
	if err != nil {
		return (err)
	}
	fragSRC, err := ioutil.ReadAll(ffrag)
	if err != nil {
		return (err)
	}
	//DefaultShader
	DefaultShader.Source.Vertex = string(vertSRC)
	DefaultShader.Source.Fragment = string(fragSRC)
	DefaultShader.Compiled.Vertex, err = compileShader(DefaultShader.Source.Vertex, gl.VERTEX_SHADER)
	if err != nil {
		return (err)
	}
	DefaultShader.Compiled.Fragment, err = compileShader(DefaultShader.Source.Fragment, gl.FRAGMENT_SHADER)
	if err != nil {
		return (err)
	}
	return nil
}
func (s *_DefaultShader) Unload() {
	gl.DeleteShader(s.Compiled.Fragment)
	gl.DeleteShader(s.Compiled.Vertex)
	//
	s.Source.Vertex = ""
	s.Source.Fragment = ""
	s.Compiled.Vertex = 0
	s.Compiled.Fragment = 0
}
func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source + "\x00")
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