package renderer

import (
	"github.com/go-gl/gl/v4.2-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Model struct {
	Vertices  []float32
	Indices   []uint32
	Normals   []float32
	TexCoords []float32

	TexIndices    []uint32
	NormalIndices []uint32

	vao          uint32
	ubo          uint32
	texBuffer    uint32
	normalBuffer uint32
}

type PerspectiveBlock struct {
	project *mgl32.Mat4
	camera  *mgl32.Mat4
	world   *mgl32.Mat4
}

func NewModel(vertices []float32, indices []uint32, normals []float32, normalIndices, texIndices []uint32, texCoords []float32) *Model {
	model := &Model{
		Vertices:      vertices,
		Indices:       indices,
		Normals:       normals,
		TexIndices:    texIndices,
		TexCoords:     texCoords,
		NormalIndices: normalIndices,
	}
	model.setupVAO()
	return model
}

func (m *Model) setupVAO() {
	var vbo, ebo uint32

	// Create VAO, VBO, EBO, and UBO
	gl.GenVertexArrays(1, &m.vao)
	gl.BindVertexArray(m.vao)

	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(m.Vertices)*4, gl.Ptr(m.Vertices), gl.STATIC_DRAW)

	gl.GenBuffers(1, &ebo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(m.Indices)*4, gl.Ptr(m.Indices), gl.STATIC_DRAW)

	gl.GenBuffers(1, &m.ubo)
	gl.BindBuffer(gl.UNIFORM_BUFFER, m.ubo)
	gl.BufferData(gl.UNIFORM_BUFFER, 3*16*4, nil, gl.DYNAMIC_DRAW)
	gl.BindBufferBase(gl.UNIFORM_BUFFER, 1, m.ubo)

	//// Create texture buffer and load texture data using texIndices
	//gl.GenBuffers(1, &m.texBuffer)
	//gl.BindBuffer(gl.ARRAY_BUFFER, m.texBuffer)
	//gl.BufferData(gl.ARRAY_BUFFER, len(m.TexCoords)*4, gl.Ptr(m.TexCoords[0]), gl.STATIC_DRAW)
	//
	//// Create normal buffer and load normal data using normalIndices
	//gl.GenBuffers(1, &m.normalBuffer)
	//gl.BindBuffer(gl.ARRAY_BUFFER, m.normalBuffer)
	//gl.BufferData(gl.ARRAY_BUFFER, len(m.Normals)*4, gl.Ptr(m.Normals), gl.STATIC_DRAW)

	// Define vertex attributes
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(0)

	// Define index attributes
	gl.VertexAttribPointer(1, 2, gl.FLOAT, false, 0, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(1)

	//// Define texture attributes
	//gl.VertexAttribPointer(2, 2, gl.FLOAT, false, 0, gl.PtrOffset(0))
	//gl.EnableVertexAttribArray(2)
	//
	//// Define normals attributes
	//gl.VertexAttribPointer(3, 3, gl.FLOAT, false, 0, gl.PtrOffset(0))
	//gl.EnableVertexAttribArray(3)

	// Unbind VAO, VBO and EBO
	gl.BindVertexArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)
	gl.BindBuffer(gl.UNIFORM_BUFFER, 0)
}

func (m *Model) Draw(shader *Shader, projectionMatrix, cameraMatrix, modelMatrix mgl32.Mat4) {
	// Bind VAO and set uniforms
	gl.BindVertexArray(m.vao)

	gl.BindBuffer(gl.UNIFORM_BUFFER, m.ubo)
	gl.BufferSubData(gl.UNIFORM_BUFFER, 0, 16*4, gl.Ptr(&projectionMatrix[0]))
	gl.BufferSubData(gl.UNIFORM_BUFFER, 16*4, 16*4, gl.Ptr(&cameraMatrix[0]))
	gl.BufferSubData(gl.UNIFORM_BUFFER, 32*4, 16*4, gl.Ptr(&modelMatrix[0]))
	gl.BindBufferBase(gl.UNIFORM_BUFFER, 1, m.ubo)

	shader.Use()

	// Draw model
	gl.DrawElements(gl.TRIANGLES, int32(len(m.Indices)), gl.UNSIGNED_INT, gl.PtrOffset(0))

	// Unbind VAO and texture
	gl.BindVertexArray(0)
	gl.BindBuffer(gl.UNIFORM_BUFFER, 0)
	gl.BindTexture(gl.TEXTURE_2D, 0)
}
