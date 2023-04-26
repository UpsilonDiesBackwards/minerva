package renderer

import (
	"fmt"
	"github.com/go-gl/gl/v4.2-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Renderer struct {
	clearColor mgl32.Vec4
}

func New() *Renderer {
	if err := gl.Init(); err != nil {
		panic(err)
	} else {
		fmt.Println("Initialised OpenGL")
	}

	return &Renderer{clearColor: mgl32.Vec4{0.0, 0.0, 0.0, 1.0}}
}

func (r *Renderer) SetClearColor(_r, g, b, a float32) {
	r.clearColor = mgl32.Vec4{_r, g, b, a}
}

func (r *Renderer) Clear() {
	gl.ClearColor(r.clearColor.X(), r.clearColor.Y(), r.clearColor.Z(), r.clearColor.W())
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func (r *Renderer) Draw(entities []*Entity) {
	defaultShader, err := NewShader("res/shaders/shader.vert", "res/shaders/shader.frag")
	if err != nil {
		panic(err)
	}

	for _, entity := range entities {
		entity.model.Draw(defaultShader, mgl32.Mat4{1.0, 1.0, 1.0, 1.0})
	}
}
