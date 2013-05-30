package main

import (
	"flag"
	"fmt"
	"github.com/remogatto/egl"
	"github.com/remogatto/egl/platform"
	gl "github.com/remogatto/opengles2"
	"log"
)

var (
	Done                                   = make(chan bool, 1)
	verticesArrayBuffer, colorsArrayBuffer uint32
	attrPos, attrColor                     uint32
	currWidth, currHeight                  int

	vertices = [12]float32{
		-1.0, -1.0, 0.0, 1.0,
		1.0, -1.0, 0.0, 1.0,
		0.0, 1.0, 0.0, 1.0,
	}
	colors = [12]float32{
		1.0, 0.0, 0.0, 1.0,
		0.0, 1.0, 0.0, 1.0,
		0.0, 0.0, 1.0, 1.0,
	}
)

func check() {
	error := gl.GetError()
	if error != 0 {
		panic(fmt.Sprintf("An error occurred! Code: 0x%x", error))
	}
}

func initShaders() {
	program := Program(FragmentShader(fsh), VertexShader(vsh))
	gl.UseProgram(program)
	attrPos = uint32(gl.GetAttribLocation(program, "pos"))
	attrColor = uint32(gl.GetAttribLocation(program, "color"))
	gl.GenBuffers(1, &verticesArrayBuffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, verticesArrayBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, gl.SizeiPtr(len(vertices))*4, gl.Void(&vertices[0]), gl.STATIC_DRAW)
	gl.GenBuffers(1, &colorsArrayBuffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, colorsArrayBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, gl.SizeiPtr(len(colors))*4, gl.Void(&colors[0]), gl.STATIC_DRAW)
	gl.EnableVertexAttribArray(attrPos)
	gl.EnableVertexAttribArray(attrColor)
}

func draw(width, height int) {
	gl.Viewport(0, 0, gl.Sizei(width), gl.Sizei(height))
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.BindBuffer(gl.ARRAY_BUFFER, verticesArrayBuffer)
	gl.VertexAttribPointer(attrPos, 4, gl.FLOAT, false, 0, nil)
	gl.BindBuffer(gl.ARRAY_BUFFER, colorsArrayBuffer)
	gl.VertexAttribPointer(attrColor, 4, gl.FLOAT, false, 0, nil)
	gl.DrawArrays(gl.TRIANGLES, 0, 3)
	gl.Flush()
	gl.Finish()
}

func cleanup() {
	egl.DestroySurface(platform.Display, platform.Surface)
	egl.DestroyContext(platform.Display, platform.Context)
	egl.Terminate(platform.Display)
}

func reshape(width, height int) {
	currWidth, currHeight = width, height
	gl.Viewport(0, 0, gl.Sizei(width), gl.Sizei(height))
}

func printInfo() {
	log.Printf("GL_RENDERER   = %s\n", gl.GetString(gl.RENDERER))
	log.Printf("GL_VERSION    = %s\n", gl.GetString(gl.VERSION))
	log.Printf("GL_VENDOR     = %s\n", gl.GetString(gl.VENDOR))
	log.Printf("GL_EXTENSIONS = %s\n", gl.GetString(gl.EXTENSIONS))
}

func main() {
	info := flag.Bool("info", false, "display OpenGL renderer info")
	flag.Parse()
	initialize()
	if *info {
		printInfo()
	}
	defer cleanup()
	for {
		select {
		case <-Done:
			return
		default:
			draw(currWidth, currHeight)
			egl.SwapBuffers(platform.Display, platform.Surface)
		}
	}
}
