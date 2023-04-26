package window

import (
	"fmt"
	"github.com/go-gl/glfw/v3.2/glfw"
)

type Window struct {
	mWindow *glfw.Window
}

func New(w, h int, title string) (*Window, error) {
	if err := glfw.Init(); err != nil {
		return nil, err
	} else {
		fmt.Println("Initialised GLFW")
	}

	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	win, err := glfw.CreateWindow(w, h, title, nil, nil)
	if err != nil {
		return nil, err
	}

	win.MakeContextCurrent()

	return &Window{mWindow: win}, nil
}

func (w *Window) ShouldClose() bool {
	return w.mWindow.ShouldClose()
}

func (w *Window) SwapBuffers() {
	w.mWindow.SwapBuffers()
}

func (w *Window) PollEvents() {
	glfw.PollEvents()
}
