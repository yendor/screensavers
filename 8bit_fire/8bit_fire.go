package main

import (
    "github.com/go-gl/gl"
    glfw "github.com/go-gl/glfw3"
    "math"
)

func drawScene(window *glfw.Window) {
    gl.Clear(gl.COLOR_BUFFER_BIT)
    gl.LoadIdentity()

    // Read timer since glfw.Init()
    // time := glfw.GetTime()

    // width, height = window.GetFramebufferSize()
    // ratio := float64(width) / float64(height)

    // gl.Rotatef(float32(time)*50.0, -0.2, 0.0, 1.0)

    // gl.Translatef(left * multiplier, 0, 0)

    var x, y, width float32
    var i, j float32
    var c color

    x = -1
    y = -1
    width = 0.25

    for j = 0; j < 8; j++ {
        for i = 0; i < 8; i++ {
            c = getColor(x + width * i, y + width * j)

            drawSquare(x + width * i, y + width * j, width, c)
        }
    }
}

// getColor determines the colour of a square at a
// certain point x, y on our window and returns the color type
func getColor(x, y float32) color {
    white  := color{255,255,255}
    yellow := color{243, 199, 3}
    orange := color{240, 55, 30}
    red    := color{183, 16, 47}
    green  := color{0, 255, 0}

    absX := math.Abs(float64(x))
    absY := math.Abs(float64(x))

    maxAbs := math.Max(absX, absY)

    if maxAbs >= 0.75 {
        return red
    } else if maxAbs > 0.5 {
        return orange
    } else if maxAbs > 0.0 {
        return yellow
    } else if maxAbs == 0 {
        return white
    } else {
        return green
    }
}