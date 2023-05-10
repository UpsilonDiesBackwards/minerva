package main

import (
	"github.com/UpsilonDiesBackwards/minerva/internal/engine/audio"
	"github.com/UpsilonDiesBackwards/minerva/internal/engine/renderer"
	"github.com/UpsilonDiesBackwards/minerva/internal/engine/scene"
	"github.com/UpsilonDiesBackwards/minerva/internal/engine/utils"
	"github.com/UpsilonDiesBackwards/minerva/internal/engine/window"
	"github.com/go-gl/mathgl/mgl32"
	"log"
	"runtime"
	"time"
)

var title = "Minerva"

func init() {
	runtime.LockOSThread()
}

func main() {
	rend := renderer.New()
	rend.SetClearColor(0.2, 0.3, 0.3, 1.0)

	win, err := window.New(title)
	if err != nil {
		log.Fatalln("Failed to create Minerva window: ", err)
	}

	var DeltaTime float64

	dScene := scene.NewScene()

	model, err := renderer.LoadModelFromFile("res/models/cube.obj")
	if err != nil {
		panic(err)
	}
	model2, err := renderer.LoadModelFromFile("res/models/cube.obj")
	if err != nil {
		panic(err)
	}
	entity := renderer.NewEntity(model, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{1, 1, 1}, mgl32.Quat{})
	entity2 := renderer.NewEntity(model2, mgl32.Vec3{-2, -2, 1}, mgl32.Vec3{1, 1, 1}, mgl32.Quat{})

	dScene.AddEntity(entity)
	dScene.AddEntity(entity2)

	audio.LoadAudio()

	var previousTime = time.Now()
	for !win.ShouldClose() {
		rend.Clear()

		DeltaTime = CalculateDeltaTime(previousTime)
		previousTime = time.Now()

		renderer.InputRunner(win, DeltaTime)
		dScene.Draw(rend, win)

		utils.EnableFPSCounter(DeltaTime)
		//utils.EnableWireFrameRendering()
		utils.GetGLErrorVerbose()

		win.SwapBuffers()
		win.PollEvents()

	}
}

func CalculateDeltaTime(previousTime time.Time) float64 {
	currentTime := time.Now()
	deltaTime := currentTime.Sub(previousTime).Seconds()
	return deltaTime
}
