package main

import (
	"github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"
	"math"
	// "log"
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


	y = -1
	x = -1
	width = 0.25

	for j = x; j < 1; j += width {
		for i = y; i < 1; i += width {
			c = getColor(j, i+1)
			drawSquare(j, i, width, c)
		}
	}
	x = -1
	c = getColor(x, y+1)
	drawSquare(x, y, width, c)

	// x = -0.75
	// c = getColor(x, y+1)
	// drawSquare(x, y, width, c)

	// x = -0.5
	// c = getColor(x, y+1)
	// drawSquare(x, y, width, c)

	// x = -0.25
	// c = getColor(x, y+1)
	// drawSquare(x, y, width, c)

	// x = 0
	// c = getColor(x, y+1)
	// drawSquare(x, y, width, c)

	// x = 0.25
	// c = getColor(x, y+1)
	// drawSquare(x, y, width, c)

	// x = 0.5
	// c = getColor(x, y+1)
	// drawSquare(x, y, width, c)

	// x = 0.75
	// c = getColor(x, y+1)
	// drawSquare(x, y, width, c)

	// x = 1
	// c = getColor(x, y+1)
	// drawSquare(x, y, width, c)


	// y = -0.75
	// x = -0.75
	// c = getColor(x, y+1)
	// drawSquare(x, y, width, c)

	// x = -0.5
	// c = getColor(x, y+1)
	// drawSquare(x, y, width, c)

	// x = -0.25
	// c = getColor(x, y+1)
	// drawSquare(x, y, width, c)

	// x = 0
	// c = getColor(x, y+1)
	// drawSquare(x, y, width, c)

	// x = 0.25
	// c = getColor(x, y+1)
	// drawSquare(x, y, width, c)

	// x = 0.5
	// c = getColor(x, y+1)
	// drawSquare(x, y, width, c)

	// x = 0.75
	// c = getColor(x, y+1)
	// drawSquare(x, y, width, c)

	// x = 1
	// c = getColor(x, y+1)
	// drawSquare(x, y, width, c)
}

// getColor determines the colour of a square at a
// certain point x, y on our window and returns the color type
func getColor(x, y float32) color {
	white := color{255, 255, 255}
	yellow := color{243, 199, 3}
	orange := color{240, 55, 30}
	red := color{183, 16, 47}
	green := color{0, 255, 0}

	black := color{0,0,0}

	colors := make([]color, 5)

	colors[0] = white
	colors[1] = yellow
	colors[2] = orange
	colors[3] = red
	colors[4] = green

	distance := getDistance(0.0, 0.0, float64(x), float64(y))

	colorIndex := int(math.Abs(distance*4))
	if colorIndex > 4 {
		return black
	}
	// log.Printf("%v -> %v\n", (x+y), colorIndex)
	// log.Print(colorIndex)

	// colorIndex := int(math.Mod(maxAbs, 5))
	return colors[colorIndex]
}

func getDistance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2 -x1, 2) + math.Pow(y2 - y1, 2))
}
