package renderer

import (
	"bufio"
	"fmt"
	"github.com/go-gl/mathgl/mgl32"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

type Entity struct {
	position, scale mgl32.Vec3
	rotation        mgl32.Quat
	model           *Model
}

func NewEntity(model *Model, position, scale mgl32.Vec3, rotation mgl32.Quat) *Entity {
	return &Entity{
		position: position,
		scale:    scale,
		rotation: rotation,
		model:    model,
	}
}

var loaderMutex sync.Mutex

func LoadModelFromFile(filePath string) (*Model, error) {
	loaderMutex.Lock()

	var vertices []float32
	var indices []uint32
	var texIndices []uint32
	var normalIndices []uint32
	var texcoords []float32
	var normals []float32

	// Create a mutex for each slice
	var (
		verticesMutex  sync.Mutex
		indicesMutex   sync.Mutex
		texcoordsMutex sync.Mutex
		normalsMutex   sync.Mutex
	)

	file, err := os.OpenFile(filePath, os.O_RDONLY, 0600)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("Could not close object file: %s\n", err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	lineCh := make(chan string)

	// Go routine for reading the file
	go func() {
		for scanner.Scan() {
			lineCh <- scanner.Text()
		}
		close(lineCh)
	}()

	// create worker pool to process lines
	numWorkers := runtime.NumCPU()
	workerCh := make(chan struct{}, numWorkers)
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		workerCh <- struct{}{}
	}

	for line := range lineCh {
		<-workerCh
		wg.Add(1)

		go func(line string) {
			if len(line) < 2 {
				return
			}

			// Process lines
			if strings.HasPrefix(line, "v ") {
				verticesMutex.Lock()

				// split line into tokens
				tokens := strings.Fields(line)

				// convert tokens to floats
				x, _ := strconv.ParseFloat(tokens[1], 32)
				y, _ := strconv.ParseFloat(tokens[2], 32)
				z, _ := strconv.ParseFloat(tokens[3], 32)

				// create a new vertex and add to model vertex array
				vertices = append(TriangulateVertices(vertices), float32(x), float32(-y), float32(z))
				verticesMutex.Unlock()
			} else if strings.HasPrefix(line, "vt ") {
				texcoordsMutex.Lock()

				tokens := strings.Fields(line)

				_u, _ := strconv.ParseFloat(tokens[1], 32)
				v, _ := strconv.ParseFloat(tokens[2], 32)

				texcoords = append(texcoords, float32(_u), float32(v))
				texcoordsMutex.Unlock()
			} else if strings.HasPrefix(line, "vn ") {
				normalsMutex.Lock()

				tokens := strings.Fields(line)

				x, _ := strconv.ParseFloat(tokens[1], 32)
				y, _ := strconv.ParseFloat(tokens[2], 32)
				z, _ := strconv.ParseFloat(tokens[3], 32)

				normals = append(normals, float32(x), float32(y), float32(z))
				normalsMutex.Unlock()
			} else if strings.HasPrefix(line, "f ") {
				tokens := strings.Fields(line)

				for i := 1; i < len(tokens); i++ {
					indicesMutex.Lock()

					vertexData := strings.Split(tokens[i], "/")

					// Extract vertex index, texture coordinate index, and normal vector index
					vertexIndex, _ := strconv.Atoi(vertexData[0])
					textureIndex, _ := strconv.Atoi(vertexData[1])
					normalIndex, _ := strconv.Atoi(vertexData[2])

					// Subtract 1 from each index value to convert from 1-based to 0-based indexing
					indices = append(indices, uint32(vertexIndex-1))
					texIndices = append(texIndices, uint32(textureIndex-1))
					normalIndices = append(normalIndices, uint32(normalIndex-1))
					indicesMutex.Unlock()
				}
			}

			wg.Done()
			workerCh <- struct{}{}
		}(line)
	}

	wg.Wait()

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	model := NewModel(vertices, indices, normals, normalIndices, texIndices, texcoords)

	defer loaderMutex.Unlock()

	return model, nil
}

func (e *Entity) GetModelMatrix() mgl32.Mat4 {
	return mgl32.Translate3D(e.position.X(), e.position.Y(), e.position.Z()).
		Mul4(e.rotation.Mat4()).
		Mul4(mgl32.Scale3D(e.scale.X(), e.scale.Y(), e.scale.Z()))
}

func (e *Entity) SetPosition(x, y, z float32) {
	e.position = mgl32.Vec3{x, y, z}
}

func (e *Entity) SetScale(x, y, z float32) {
	e.scale = mgl32.Vec3{x, y, z}
}

func (e *Entity) SetRotation(angle float32, axis mgl32.Vec3) {
	e.rotation = mgl32.QuatRotate(angle, axis)
}

func (e *Entity) Rotate(angle float32, axis mgl32.Vec3) {
	e.rotation = e.rotation.Mul(mgl32.QuatRotate(angle, axis))
}

func (e *Entity) SetModel(model *Model) {
	e.model = model
}

func (e *Entity) Update(deltaTime float32) {
	// Change this code
	e.position = e.position.Add(mgl32.Vec3{0.0, 0.0, 1.0}.Mul(deltaTime))
	e.rotation = e.rotation.Mul(mgl32.QuatRotate(deltaTime, mgl32.Vec3{0.0, 1.0, 0.0}))
	e.scale = e.scale.Add(mgl32.Vec3{0.1, 0.1, 0.1}.Mul(deltaTime))
}

func TriangulateVertices(vertices []float32) []float32 {
	// Triangulate *.obj mesh. I have no idea if it works or not, but it gets angry if I don't use it. :O
	triangulatedVertices := make([]float32, 0)
	for i := 0; i < len(vertices); i += 3 {
		triangulatedVertices = append(triangulatedVertices, vertices[i], vertices[i+1], vertices[i+2])
	}
	return triangulatedVertices
}
