package renderer

import (
	"fmt"
	"github.com/go-gl/gl/v4.2-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"os"
	"strings"
)

type Shader struct {
	Program uint32
}

func NewShader(vFile, fFile string) (*Shader, error) {
	vSource, err := os.ReadFile(vFile)
	if err != nil {
		return nil, err
	}

	fSource, err := os.ReadFile(fFile)
	if err != nil {
		return nil, err
	}

	// Compile vertex shader
	vShader := gl.CreateShader(gl.VERTEX_SHADER)
	cVertSource, free := gl.Strs(string(vSource) + "\x00")
	gl.ShaderSource(vShader, 1, cVertSource, nil)
	gl.CompileShader(vShader)
	free()
	var status int32
	gl.GetShaderiv(vShader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(vShader, gl.INFO_LOG_LENGTH, &logLength)
		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(vShader, logLength, nil, gl.Str(log))
		return nil, fmt.Errorf("failed to compile Vertex Shader: %v", log)
	}

	fShader := gl.CreateShader(gl.FRAGMENT_SHADER)
	cFragSource, free := gl.Strs(string(fSource) + "\x00")
	gl.ShaderSource(fShader, 1, cFragSource, nil)
	gl.CompileShader(fShader)
	free()
	gl.GetShaderiv(fShader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(fShader, gl.INFO_LOG_LENGTH, &logLength)
		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(fShader, logLength, nil, gl.Str(log))
		return nil, fmt.Errorf("failed to compile Fragment Shader: %v", log)
	}

	// Link shader program
	program := gl.CreateProgram()
	gl.AttachShader(program, vShader)
	gl.AttachShader(program, fShader)
	gl.LinkProgram(program)
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)
		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))
		return nil, fmt.Errorf("failed to link shader program: %v", log)
	}

	// Delete shader objects
	gl.DeleteShader(vShader)
	gl.DeleteShader(fShader)

	return &Shader{Program: program}, nil
}

func (s *Shader) Use() {
	gl.UseProgram(s.Program)
}

func (s *Shader) SetInt(name string, value int32) {
	gl.Uniform1i(gl.GetUniformLocation(s.Program, gl.Str(name+"\x00")), value)
}

func (s *Shader) SetFloat(name string, value float32) {
	gl.Uniform1f(gl.GetUniformLocation(s.Program, gl.Str(name+"\x00")), value)
}

func (s *Shader) SetVec3(name string, value mgl32.Vec3) {
	gl.Uniform3fv(gl.GetUniformLocation(s.Program, gl.Str(name+"\x00")), 1, &value[0])
}

func (s *Shader) SetVec4(name string, value mgl32.Vec4) {
	gl.Uniform4fv(gl.GetUniformLocation(s.Program, gl.Str(name+"\x00")), 1, &value[0])
}

func (s *Shader) SetMat4(name string, value mgl32.Mat4) {
	gl.UniformMatrix4fv(gl.GetUniformLocation(s.Program, gl.Str(name+"\x00")), 1, false, &value[0])
}

func (s *Shader) Delete() {
	gl.DeleteProgram(s.Program)
}
