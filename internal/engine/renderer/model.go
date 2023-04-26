package renderer

import (
	"github.com/go-gl/gl/v4.2-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Model struct {
	Vertices  []float32
	Indices   []uint32
	Texcoords []float32
	Normals   []float32

	vao uint32
}

func NewModel(vertices []float32, indices []uint32) *Model {
	model := &Model{
		Vertices: vertices,
		Indices:  indices,
	}

	model.setupVAO()

	return model
}

func (m *Model) setupVAO() {
	var vbo, ebo uint32

	// Create VAO, VBO and EBO
	gl.GenVertexArrays(1, &m.vao)
	gl.BindVertexArray(m.vao)

	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(m.Vertices)*4, gl.Ptr(m.Vertices), gl.STATIC_DRAW)

	gl.GenBuffers(1, &ebo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(m.Indices)*4, gl.Ptr(m.Indices), gl.STATIC_DRAW)

	// Define vertex attributes
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 5*4, nil)
	gl.EnableVertexAttribArray(0)

	gl.VertexAttribPointer(1, 2, gl.FLOAT, false, 5*4, gl.PtrOffset(3*4))
	gl.EnableVertexAttribArray(1)

	// Unbind VAO, VBO and EBO
	gl.BindVertexArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)
}

func (m *Model) Draw(shader *Shader, modelMatrix mgl32.Mat4) {
	// Bind VAO and set uniforms
	gl.BindVertexArray(m.vao)
	shader.SetMat4("model", modelMatrix)

	shader.Use()

	// Draw model
	gl.DrawElements(gl.TRIANGLES, int32(len(m.Indices)), gl.UNSIGNED_INT, nil)

	// Unbind VAO and texture
	gl.BindVertexArray(0)
	gl.BindTexture(gl.TEXTURE_2D, 0)
}
