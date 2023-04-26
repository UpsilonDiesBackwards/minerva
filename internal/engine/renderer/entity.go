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

func LoadModelFromFile(filePath string) (*Model, error) {
	var vertices []float32
	var indices []uint32
	var texcoords []float32
	var normals []float32

	fmt.Println("Attempting to read object file!")

	file, err := os.OpenFile(filePath, os.O_RDONLY, 0600)
	if err != nil {
		return nil, err
	} else {
		fmt.Println("Successfully opened object file")
	}
	defer file.Close()

	// Create a mutex to synchronize access to the slices
	var mutex sync.Mutex

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
				// split line into tokens
				tokens := strings.Fields(line)

				// convert tokens to floats
				x, _ := strconv.ParseFloat(tokens[1], 64)
				y, _ := strconv.ParseFloat(tokens[2], 64)
				z, _ := strconv.ParseFloat(tokens[3], 64)

				// create a new vertex and add to model vertex array
				mutex.Lock()
				vertices = append(vertices, float32(x), float32(y), float32(z))
				fmt.Println("Added vertices: ")
				mutex.Unlock()
			} else if strings.HasPrefix(line, "vt ") {
				tokens := strings.Fields(line)

				u, _ := strconv.ParseFloat(tokens[1], 64)
				v, _ := strconv.ParseFloat(tokens[2], 64)

				mutex.Lock()
				texcoords = append(texcoords, float32(u), float32(v))
				fmt.Println("Added vertex textures")
				mutex.Unlock()
			} else if strings.HasPrefix(line, "vn ") {
				tokens := strings.Fields(line)

				x, _ := strconv.ParseFloat(tokens[1], 64)
				y, _ := strconv.ParseFloat(tokens[2], 64)
				z, _ := strconv.ParseFloat(tokens[3], 64)

				mutex.Lock()
				normals = append(normals, float32(x), float32(y), float32(z))
				fmt.Println("Added vertex normals")
				mutex.Unlock()
			} else if strings.HasPrefix(line, "f ") {
				tokens := strings.Fields(line)

				for i := 1; i < len(tokens); i++ {
					vertexData := strings.Split(tokens[i], "/")

					vertexIndex, _ := strconv.Atoi(vertexData[0])
					//textureIndex, _ := strconv.Atoi(vertexData[1])
					//normalIndex, _ := strconv.Atoi(vertexData[2])

					// Subtract 1 from each index value to convert from 1-based to 0-based indexing
					mutex.Lock()
					indices = append(indices, uint32(vertexIndex-1))
					fmt.Println("Added vertex index data")
					mutex.Unlock()
				}
			}

			fmt.Println("AHHH")

			wg.Done()
			workerCh <- struct{}{}
		}(line)
	}

	fmt.Println("AHHH2")

	wg.Wait()
	fmt.Println("AHHH3")

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	fmt.Println("AHHH4")

	// Create the model
	// TODO: implement normals and texture normals
	fmt.Println("vert: ", vertices, "ind:", indices)

	model := NewModel(vertices, indices)
	fmt.Println("model vert: ", model.Vertices, "model ind:", model.Indices)

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
