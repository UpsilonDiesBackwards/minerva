package viewport

import (
	"github.com/UpsilonDiesBackwards/minerva/internal/engine/input"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/go-gl/mathgl/mgl64"
	"math"
)

// Define what a struct is
type Camera struct {
	Position mgl64.Vec3
	Front    mgl64.Vec3

	Up      mgl64.Vec3
	Right   mgl64.Vec3
	WorldUp mgl64.Vec3

	Yaw   float64
	Pitch float64

	Speed       float64
	Sensitivity float64
}

// Create a new Camera
var CameraViewport = Camera{
	Position: mgl64.Vec3{0, 0, 0},
	Up:       mgl64.Vec3{0, 1, 0},
	WorldUp:  mgl64.Vec3{0, 1, 0},
	Yaw:      -90,
	Pitch:    0,
	Speed:    12,

	Sensitivity: 0.075,
}

func (c *Camera) UpdateDirection(u *input.UserInput) {
	cursor := u.CursorChange()

	dx := -c.Sensitivity * cursor[0]
	dy := c.Sensitivity * cursor[1]

	c.Pitch += dy
	if c.Pitch > 89 {
		c.Pitch = 89
	} else if c.Pitch < -89 {
		c.Pitch = -89
	}
	c.Yaw = math.Mod(c.Yaw+dx, 360)
	c.UpdateVec()
}

func (c *Camera) UpdateVec() {
	// x, y, z
	c.Front[0] = math.Cos(mgl64.DegToRad(c.Pitch)) * math.Cos(mgl64.DegToRad(c.Yaw))
	c.Front[1] = math.Sin(mgl64.DegToRad(c.Pitch))
	c.Front[2] = math.Cos(mgl64.DegToRad(c.Pitch)) * math.Sin(mgl64.DegToRad(c.Yaw))
	c.Front = c.Front.Normalize()

	// Gram-Schmidt process to figure out right and up vectors
	c.Right = c.WorldUp.Cross(c.Front).Normalize()
	c.Up = c.Right.Cross(c.Front).Normalize()
}

func (camera *Camera) GetTransform() mgl32.Mat4 {
	cameraTarget := camera.Position.Add(camera.Front)

	return mgl32.LookAt(
		float32(camera.Position.X()), float32(camera.Position.Y()), float32(camera.Position.Z()),
		float32(cameraTarget.X()), float32(cameraTarget.Y()), float32(cameraTarget.Z()),
		float32(camera.Up.X()), float32(camera.Up.Y()), float32(camera.Up.Z()),
	)
}
