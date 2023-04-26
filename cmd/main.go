package main

import (
	"github.com/UpsilonDiesBackwards/minerva/internal/engine/renderer"
	"github.com/UpsilonDiesBackwards/minerva/internal/engine/scene"
	"github.com/UpsilonDiesBackwards/minerva/internal/engine/window"
	"github.com/go-gl/mathgl/mgl32"
	"log"
	"runtime"
)

const (
	width  int    = 1024
	height int    = 768
	title  string = "Minerva"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	rend := renderer.New()

	win, err := window.New(width, height, title)
	if err != nil {
		log.Fatalln("Failed to create Minerva window: ", err)
	}

	rend.SetClearColor(0.2, 0.3, 0.3, 1.0)

	dScene := scene.NewScene()

	model, err := renderer.LoadModelFromFile("res/models/cube.obj")
	if err != nil {
		panic(err)
	}

	entity := renderer.NewEntity(model, mgl32.Vec3{0, 0, 1}, mgl32.Vec3{0, 0, 0}, mgl32.Quat{})

	dScene.AddEntity(entity)

	for !win.ShouldClose() {
		rend.Clear()

		dScene.Draw(rend)

		win.SwapBuffers()
		win.PollEvents()
	}

}
