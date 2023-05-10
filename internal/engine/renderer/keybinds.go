package renderer

import (
	"fmt"
	"github.com/UpsilonDiesBackwards/minerva/internal/engine/input"
	"github.com/UpsilonDiesBackwards/minerva/internal/engine/viewport"
	"github.com/UpsilonDiesBackwards/minerva/internal/engine/window"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

var ViewportTransform mgl32.Mat4
var u = &input.UserInput{}

func InputRunner(win *window.Window, deltaTime float64) error {
	c := &viewport.CameraViewport

	adjustedViewportSpeed := deltaTime * float64(c.Speed)

	if input.ActionState[input.VIEWPORT_FORWARD] {
		// Move the camera forward.
		c.Position = c.Position.Add(c.Front.Mul(adjustedViewportSpeed))
	}
	if input.ActionState[input.VIEWPORT_BACKWARDS] {
		// Move the camera backward.
		c.Position = c.Position.Sub(c.Front.Mul(adjustedViewportSpeed))
	}
	if input.ActionState[input.VIEWPORT_LEFT] {
		// Move the camera to the left.
		c.Position = c.Position.Sub(c.Front.Cross(c.Up).Mul(adjustedViewportSpeed))
	}
	if input.ActionState[input.VIEWPORT_RIGHT] {
		// Move the camera to the right.
		c.Position = c.Position.Add(c.Front.Cross(c.Up).Mul(adjustedViewportSpeed))
	}
	if input.ActionState[input.VIEWPORT_RAISE] {
		// Move the camera to the right.
		c.Position = c.Position.Add(c.Up.Mul(adjustedViewportSpeed))
	}
	if input.ActionState[input.VIEWPORT_LOWER] {
		// Move the camera to the right.
		c.Position = c.Position.Sub(c.Up.Mul(adjustedViewportSpeed))
	}

	if input.ActionState[input.INPUT_TEST] {
		fmt.Println("Input test!")
	}

	if input.ActionState[input.QUIT_PROGRAM] {
		fmt.Println("Exiting!")
		glfw.Terminate()
	}

	// Cursor transform
	ViewportTransform = c.GetTransform()
	c.UpdateDirection(u)
	u.CheckpointCursorChange()

	input.InputManager(win, u)
	return nil
}
