package glumi

import (
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"image"
	"strings"
)

type GLRender struct {
	glumi *GLUMI
	//
	program    uint32
	frameimage uint32
	vao, vbo   uint32
}

func (s *GLRender) init() error {
	err := s.newProgram()
	if err != nil {
		return err
	}
	s.newImage()
	s.newVAO()
	s.newVBO()
	s.attrib()
	return nil
}
func (s *GLRender) newProgram() error {
	program := gl.CreateProgram()
	gl.AttachShader(program, DefaultShader.Compiled.Vertex)
	gl.AttachShader(program, DefaultShader.Compiled.Fragment)
	gl.LinkProgram(program)

	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))

		return fmt.Errorf("failed to link program: %v", log)
	}
	s.program = program
	return nil
}
func (s *GLRender) newImage() {
	gl.GenTextures(1, &s.frameimage)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, s.frameimage)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
}
func (s *GLRender) setImage(img *image.RGBA) {
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, s.frameimage)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(img.Bounds().Size().X),
		int32(img.Bounds().Size().Y),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(img.Pix))
}
func (s *GLRender) newVAO() {
	gl.GenVertexArrays(1, &s.vao)
	gl.BindVertexArray(s.vao)
}
func (s *GLRender) newVBO() {
	gl.GenBuffers(1, &s.vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, s.vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(ScreenModeling)*4*5, gl.Ptr(ScreenModeling), gl.STATIC_DRAW)
}
func (s *GLRender) attrib() {
	vertAttrib := uint32(gl.GetAttribLocation(s.program, gl.Str("vert\x00")))
	gl.EnableVertexAttribArray(vertAttrib)
	gl.VertexAttribPointer(vertAttrib, 3, gl.FLOAT, false, 5*4, gl.PtrOffset(0))

	texCoordAttrib := uint32(gl.GetAttribLocation(s.program, gl.Str("vertTexCoord\x00")))
	gl.EnableVertexAttribArray(texCoordAttrib)
	gl.VertexAttribPointer(texCoordAttrib, 2, gl.FLOAT, false, 5*4, gl.PtrOffset(3*4))
}
func (s *GLRender) Upload() {
	s.setImage(s.glumi.screen.RGBA())
}
func (s *GLRender) Draw() {
	gl.UseProgram(s.program)
	gl.BindVertexArray(s.vao)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, s.frameimage)
	gl.DrawArrays(gl.TRIANGLES, 0, 6*2*3)
}
