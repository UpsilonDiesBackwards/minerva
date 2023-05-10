package utils

import (
	"fmt"
	"github.com/go-gl/gl/v4.2-core/gl"
	"log"
	"strconv"
	"time"
	"unsafe"
)

func GetGLErrorVerbose() {
	if err := gl.GetError(); err != 0 {
		gl.Enable(gl.DEBUG_OUTPUT)
		gl.DebugMessageCallback(func(source uint32, gltype uint32, id uint32, severity uint32, length int32, message string, userParam unsafe.Pointer) {
			log.Printf("OpenGL Debug Message (source: 0x%X, type: 0x%X, id: %d, severity: 0x%X): %s\n", source, gltype, id, severity, message)
		}, gl.Ptr(nil))
	}
}

// FPS Counter
var startTime = time.Now()
var frameCount int
var FPS float64

var previousTime = time.Now()

func EnableFPSCounter(DeltaTime float64) {
	currentTime := time.Now()
	DeltaTime = currentTime.Sub(previousTime).Seconds()
	previousTime = currentTime

	// Handle FPS counter
	frameCount++
	if time.Since(startTime) >= time.Second {
		// Calculate the FPS
		FPS = float64(frameCount) / time.Since(startTime).Seconds()

		// Print the FPS
		str := strconv.FormatFloat(FPS, 'f', 3, 64)
		fmt.Printf("\rFPS %s", str)

		// Reset the frame count and start time
		frameCount = 0
		startTime = time.Now()
	}
}

func EnableWireFrameRendering() {
	// Enable line smoothing and blending
	gl.Enable(gl.BLEND)
	gl.Enable(gl.LINE_SMOOTH)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.Hint(gl.LINE_SMOOTH_HINT, gl.NICEST)

	// Set polygon mode to GL_LINE
	gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)
}
