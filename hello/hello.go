package main

import (
	"fmt"
	"github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"
)

var (
	width, height    int
	centerX, centerY float32
)

func errorCallback(err glfw.ErrorCode, desc string) {
	fmt.Printf("%v: %v\n", err, desc)
}

func main() {
	glfw.SetErrorCallback(errorCallback)

	if !glfw.Init() {
		panic("Can't init glfw!")
	}
	defer glfw.Terminate()

	monitor, err := glfw.GetPrimaryMonitor()
	if err != nil {
		panic(err)
	}

	monitorMode, err := monitor.GetVideoMode()
	if err != nil {
		panic(err)
	}
	width = monitorMode.Width
	height = monitorMode.Height

	window, err := glfw.CreateWindow(monitorMode.Width, monitorMode.Height, "Testing", nil, nil)
	if err != nil {
		panic(err)
	}

	window.SetFramebufferSizeCallback(framebufferSizeCallback)

	window.MakeContextCurrent()

	for !window.ShouldClose() {
		//Do OpenGL stuff
		drawScene(window)

		window.SwapBuffers()
		glfw.PollEvents()
	}
}

func framebufferSizeCallback(window *glfw.Window, width, height int) {
	ratio := float64(width) / float64(height)

	updateDimensions(window)

	gl.Viewport(0, 0, width, height)
	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(-ratio, ratio, -1.0, 1.0, 1.0, -1.0)

	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
}

func updateDimensions(window *glfw.Window) {
	width, height = window.GetFramebufferSize()
	centerX = float32(width) / 2
	centerY = float32(height) / 2
}

func drawScene(window *glfw.Window) {
	var mouseX, mouseY float64

	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.LoadIdentity()

	// Read timer since glfw.Init()
	time := glfw.GetTime()

	mouseX, mouseY = window.GetCursorPosition()
	mouseRatioX, mouseRatioY := float32(mouseX)/float32(width), float32(mouseY)/float32(height)
	gl.Translatef(mouseRatioX-(centerX/float32(width)), -mouseRatioY+(centerY/float32(height)), 0)
	gl.Rotatef(float32(time)*50.0, -0.2, 0.0, 1.0)
	/* gl.Translatef(-0.9, 0, 0)*/

	gl.Begin(gl.TRIANGLES)
	gl.Color3f(1.0, 0.0, 0.0)
	gl.Vertex3f(-0.6, -0.4, 0.0)
	gl.Color3f(0.0, 1.0, 0.0)
	gl.Vertex3f(0.6, -0.4, 0.0)
	gl.Color3f(0.0, 0.0, 1.0)
	gl.Vertex3f(0.0, 0.6, 0.0)
	gl.End()
}
