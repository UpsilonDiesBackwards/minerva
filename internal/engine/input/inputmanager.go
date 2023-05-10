package input

import (
	"github.com/UpsilonDiesBackwards/minerva/internal/engine/window"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl64"
)

type KeyAction int

// UserInput Types of user input
type UserInput struct {
	InitialAction bool // Keyboard

	cursor         mgl64.Vec2 // Mouse
	cursorChange   mgl64.Vec2 //
	cursorLast     mgl64.Vec2 //
	bufferedChange mgl64.Vec2 //
}

// Key Actions
const (
	NO_ACTION = iota

	VIEWPORT_FORWARD
	VIEWPORT_BACKWARDS
	VIEWPORT_LEFT
	VIEWPORT_RIGHT
	VIEWPORT_RAISE
	VIEWPORT_LOWER

	INPUT_TEST
	QUIT_PROGRAM
)

var ActionState = make(map[KeyAction]bool)

// Using the map we just created, map key codes to key action
var keyToActionMap = map[glfw.Key]KeyAction{
	glfw.KeyW:         VIEWPORT_FORWARD,
	glfw.KeyS:         VIEWPORT_BACKWARDS,
	glfw.KeyA:         VIEWPORT_LEFT,
	glfw.KeyD:         VIEWPORT_RIGHT,
	glfw.KeySpace:     VIEWPORT_RAISE,
	glfw.KeyLeftShift: VIEWPORT_LOWER,

	glfw.KeyTab:    INPUT_TEST,
	glfw.KeyEscape: QUIT_PROGRAM,
}

// InputManager Create an InputManager
func InputManager(aW *window.Window, uI *UserInput) {
	aW.SetKeyCallback(KeyCallBack)
	aW.SetCursorPosCallback(uI.MouseCallBack)
}

// KeyCallBack Create Keyboard call back.
func KeyCallBack(aW *glfw.Window, key glfw.Key, scancode int, action glfw.Action, modifier glfw.ModifierKey) {
	// Get corresponding action for key code
	a, ok := keyToActionMap[key]
	if !ok {
		return
	} // Key code not in action map, return

	// Update action state on key event
	switch action {
	case glfw.Press: // Key was pressed
		ActionState[a] = true
	case glfw.Release: // Key was released
		ActionState[a] = false
	}
}

// Cursor Create Cursor
func (cInput UserInput) Cursor() mgl64.Vec2 { return cInput.cursor }

// CursorChange Cursor Change
func (cInput UserInput) CursorChange() mgl64.Vec2 { return cInput.cursorChange }

// CheckpointCursorChange Checkpoint cursor change
func (cInput *UserInput) CheckpointCursorChange() {
	cInput.cursorChange[0] = cInput.bufferedChange[0]
	cInput.cursorChange[1] = cInput.bufferedChange[1]

	cInput.cursor[0] = cInput.cursor[0]
	cInput.cursor[1] = cInput.cursor[1]

	cInput.bufferedChange[0] = 0
	cInput.bufferedChange[1] = 0
}

// MouseCallBack Create Mouse Callback - Will be used for viewport camera transform update
func (cInput *UserInput) MouseCallBack(aW *glfw.Window, xpos, ypos float64) {
	if cInput.InitialAction {
		cInput.cursorLast[0] = xpos
		cInput.cursorLast[1] = ypos
		cInput.InitialAction = false
	}

	cInput.bufferedChange[0] += xpos - cInput.cursorLast[0]
	cInput.bufferedChange[1] += ypos - cInput.cursorLast[1]

	cInput.cursorLast[0] = xpos
	cInput.cursorLast[1] = ypos
}
