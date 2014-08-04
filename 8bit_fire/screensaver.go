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