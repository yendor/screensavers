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

type color struct {
    red, green, blue float32
}


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
    width = monitorMode.Width;
    height = monitorMode.Height;

    width = 1024
    height = 1024

    window, err := glfw.CreateWindow(width, height, "Testing", nil, nil)
    if err != nil {
        panic(err)
    }


    window.SetFramebufferSizeCallback(framebufferSizeCallback)

    window.MakeContextCurrent()

    for !window.ShouldClose() {
        //Do OpenGL stuff
        drawScene(window);
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
    gl.Clear(gl.COLOR_BUFFER_BIT)
    gl.LoadIdentity()

    // Read timer since glfw.Init()
    time := glfw.GetTime()

    // width, height = window.GetFramebufferSize()
    // ratio := float64(width) / float64(height)

    // gl.Rotatef(float32(time)*50.0, -0.2, 0.0, 1.0)
    // multiplier := float32(1.0)
    // left := float32(time)
    // if left >= float32(ratio) {
    //     multiplier = - 0.1
    // }

    // if left <= float32(-ratio) {
    //     multiplier = 0.1
    // }

    // gl.Translatef(left * multiplier, 0, 0)

    var left, top, width float32
    var i, j float32


    left = -1
    top = -1
    width = 0.5


    for j = 0; j < 4; j++ {
        for i = 0; i < 4; i++ {
            c := color{1.0/(i+1) * float32(time), 1.0/(j+1) * float32(time), 0.5 * float32(time)}
            drawSquare(left + width * i, top + width * j, width, c)
        }
    }
}

func drawSquare(left float32, top float32, width float32, c color) {
    gl.Begin(gl.QUADS)

    gl.Color3f(c.red, c.green, c.blue)

    // bottom left
    gl.Vertex3f(left, top, 0.0)

    // top left
    gl.Vertex3f(left, top + width, 0.0)

    // top right
    gl.Vertex3f(left + width, top + width, 0.0)

    // bottom right
    gl.Vertex3f(left + width, top, 0.0)

    gl.End()
}