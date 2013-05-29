package main

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/mousebind"
	"github.com/BurntSushi/xgbutil/xevent"
	"github.com/mortdeus/mathgl"
	"github.com/remogatto/egl"
	platform "github.com/remogatto/egl/platforms/xorg"
	gl "github.com/remogatto/opengles2"
	"log"
	"math"
	"time"
)

const (
	INITIAL_WINDOW_WIDTH = 640
	INITIAL_WINDOW_HEIGHT = 480
)

var (
	X                                      *xgbutil.XUtil
	Done                                   = make(chan bool, 1)
	verticesArrayBuffer, colorsArrayBuffer uint32
	redraw                                 = true
	attrPos, attrColor                     uint32
	viewRotX                               float32
	viewRotY                               float32
	uMatrix                                int32

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

func initialize() {
	X, err := xgbutil.NewConn()
	if err != nil {
		log.Fatal(err)
	}
	mousebind.Initialize(X)
	xWindow := newWindow(X, INITIAL_WINDOW_WIDTH, INITIAL_WINDOW_HEIGHT)
	go xevent.Main(X)
	platform.Initialize(
		egl.NativeWindowType(uintptr(xWindow.Id)),
		platform.DefaultConfigAttributes,
		platform.DefaultContextAttributes)
	gl.Viewport(0, 0, INITIAL_WINDOW_WIDTH, INITIAL_WINDOW_HEIGHT)
	gl.ClearColor(0.5, 0.5, 0.5, 1.0)
	initShaders()
}

func initShaders() {
	program := Program(FragmentShader(fsh), VertexShader(vsh))
	gl.UseProgram(program)
	attrPos = uint32(gl.GetAttribLocation(program, "pos"))
	attrColor = uint32(gl.GetAttribLocation(program, "color"))
	uMatrix = int32(gl.GetUniformLocation(program, "modelviewProjection"))
	gl.GenBuffers(1, &verticesArrayBuffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, verticesArrayBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, gl.SizeiPtr(len(vertices))*4, gl.Void(&vertices[0]), gl.STATIC_DRAW)
	gl.GenBuffers(1, &colorsArrayBuffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, colorsArrayBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, gl.SizeiPtr(len(colors))*4, gl.Void(&colors[0]), gl.STATIC_DRAW)
	gl.EnableVertexAttribArray(attrPos)
	gl.EnableVertexAttribArray(attrColor)
}

func update() {
	time.Sleep(time.Millisecond * 10)
}

func draw() {
	var mat, rot, scale mathgl.Mat4

	viewRotX += 0.1
	makeZRotMatrix(float32(viewRotX), &rot)
	makeScaleMatrix(0.5, 0.5, 0.5, &scale)
	rot.Multiply(&scale)
	mat = rot
	gl.UniformMatrix4fv(uMatrix, 1, false, (*float32)(&mat[0]))

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
	gl.Viewport(0, 0, gl.Sizei(width), gl.Sizei(height))
}

func makeZRotMatrix(angle float32, m *mathgl.Mat4) {
	c := float32(math.Cos(float64(angle) * math.Pi / 180.0))
	s := float32(math.Sin(float64(angle) * math.Pi / 180.0))
	m.Identity()
	m[0] = c
	m[1] = s
	m[4] = -s
	m[5] = c
}

func makeScaleMatrix(xs, ys, zs float32, m *mathgl.Mat4) {
	m[0] = xs
	m[5] = ys
	m[10] = zs
	m[15] = 1.0
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
			draw()
			egl.SwapBuffers(platform.Display, platform.Surface)
		}
	}
}
