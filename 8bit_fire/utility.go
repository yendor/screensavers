package main

import (
    "github.com/go-gl/gl"
)

type color struct {
    red, green, blue float32
}

func drawSquare(x float32, y float32, width float32, c color) {
    gl.Begin(gl.QUADS)

    gl.Color3f((c.red+1)/256, (c.green+1)/256, (c.blue+1)/256)

    // bottom left
    gl.Vertex3f(x, y, 0.0)

    // top left
    gl.Vertex3f(x, y + width, 0.0)

    // top right
    gl.Vertex3f(x + width, y + width, 0.0)

    // bottom right
    gl.Vertex3f(x + width, y, 0.0)

    gl.End()
}