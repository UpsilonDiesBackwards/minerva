package renderer

import (
	"fmt"
	"github.com/UpsilonDiesBackwards/minerva/internal/engine/window"
	"github.com/go-gl/gl/v4.2-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Renderer struct {
	ClearColor mgl32.Vec4
}

func New() *Renderer {
	if err := gl.Init(); err != nil {
		panic(err)
	} else {
		fmt.Println("Initialised OpenGL")
	}

	return &Renderer{ClearColor: mgl32.Vec4{0.0, 0.0, 0.0, 1.0}}
}

func (r *Renderer) SetClearColor(_r, g, b, a float32) {
	r.ClearColor = mgl32.Vec4{_r, g, b, a}
}

func (r *Renderer) Clear() {
	gl.ClearColor(r.ClearColor.X(), r.ClearColor.Y(), r.ClearColor.Z(), r.ClearColor.W())
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func (r *Renderer) Draw(entities []*Entity, win *window.Window) {
	var fov = float32(60.0)
	var projectionTransform = mgl32.Perspective(mgl32.DegToRad(fov),
		win.AspectRatio(), 0.1, 2000)

	defaultShader, err := NewShader("res/shaders/shader.vert", "res/shaders/shader.frag")
	if err != nil {
		panic(err)
	}

	for _, entity := range entities {
		entity.model.Draw(defaultShader, projectionTransform, ViewportTransform, entity.GetModelMatrix())
	}
}
