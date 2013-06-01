package main

import (
	"flag"
	"fmt"
	"github.com/remogatto/application"
	"github.com/remogatto/egl"
	"github.com/remogatto/egl/platform"
	gl "github.com/remogatto/opengles2"
	"image/png"
	"log"
	"os"
	"runtime"
	"syscall"
	"time"
)

const FRAMES_PER_SECOND = 24

var (
	signal                          sigterm
	verticesArrayBuffer             uint32
	textureBuffer                   uint32
	unifTexture, attrPos, attrTexIn uint32
	currWidth, currHeight           int

	vertices = [24]float32{
		-1.0, -1.0, 0.0, 1.0, 0.0, 1.0,
		1.0, -1.0, 0.0, 1.0, 1.0, 1.0,
		1.0, 1.0, 0.0, 1.0, 1.0, 0.0,
		-1.0, 1.0, 0.0, 1.0, 0.0, 0.0,
	}
)

// sigterm is a type for handling a SIGTERM signal.
type sigterm int

func (h *sigterm) HandleSignal(s os.Signal) {
	switch ss := s.(type) {
	case syscall.Signal:
		switch ss {
		case syscall.SIGTERM, syscall.SIGINT:
			application.Exit()
		}
	}
}

// emulatorLoop sends a cmdRenderFrame command to the rendering backend
// (displayLoop) each 1/50 second.
type renderLoop struct {
	ticker           *time.Ticker
	pause, terminate chan int
}

// newRenderLoop returns a new renderLoop instance. It takes the
// number of frame-per-second as argument.
func newRenderLoop(fps int) *renderLoop {
	renderLoop := &renderLoop{
		ticker:    time.NewTicker(time.Duration(1e9 / fps)),
		pause:     make(chan int),
		terminate: make(chan int),
	}
	return renderLoop
}

// Pause returns the pause channel of the loop.
// If a value is sent to this channel, the loop will be paused.
func (l *renderLoop) Pause() chan int {
	return l.pause
}

// Terminate returns the terminate channel of the loop.
// If a value is sent to this channel, the loop will be terminated.
func (l *renderLoop) Terminate() chan int {
	return l.terminate
}

// Run runs renderLoop.
// The loop renders a frame and swaps the buffer for each tick
// received.
func (l *renderLoop) Run() {
	runtime.LockOSThread()
	initialize()
	gl.Viewport(0, 0, INITIAL_WINDOW_WIDTH, INITIAL_WINDOW_HEIGHT)
	gl.ClearColor(0.0, 0.0, 0.0, 1.0)
	initShaders()
	for {
		select {
		case <-l.pause:
			l.ticker.Stop()
			l.pause <- 0
		case <-l.terminate:
			cleanup()
			l.terminate <- 0
		case <-l.ticker.C:
			draw(currWidth, currHeight)
			egl.SwapBuffers(platform.Display, platform.Surface)
		}
	}
}

func check() {
	error := gl.GetError()
	if error != 0 {
		panic(fmt.Sprintf("An error occurred! Code: 0x%x", error))
	}
}

func loadImage() ([]byte, int, int) {
	// Open the file.
	file, err := os.Open("gopher.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Decode the image.
	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	bounds := img.Bounds()
	width, height := bounds.Size().X, bounds.Size().Y
	buffer := make([]byte, width*height*4)
	index := 0
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			buffer[index] = byte(r)
			buffer[index+1] = byte(g)
			buffer[index+2] = byte(b)
			buffer[index+3] = byte(a)
			index += 4
		}
	}
	return buffer, width, height
}

func initShaders() {
	program := Program(FragmentShader(fsh), VertexShader(vsh))
	gl.UseProgram(program)

	attrPos = uint32(gl.GetAttribLocation(program, "pos"))
	attrTexIn = uint32(gl.GetAttribLocation(program, "texIn"))
	unifTexture = gl.GetUniformLocation(program, "texture")
	gl.EnableVertexAttribArray(attrPos)
	gl.EnableVertexAttribArray(attrTexIn)

	// Upload vertices data
	gl.GenBuffers(1, &verticesArrayBuffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, verticesArrayBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, gl.SizeiPtr(len(vertices))*4, gl.Void(&vertices[0]), gl.STATIC_DRAW)

	// Upload texture data
	imageBuffer, width, height := loadImage()
	gl.GenTextures(1, &textureBuffer)
	gl.BindTexture(gl.TEXTURE_2D, textureBuffer)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, gl.Sizei(width), gl.Sizei(height), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Void(&imageBuffer[0]))
}

func draw(width, height int) {
	gl.Viewport(0, 0, gl.Sizei(width), gl.Sizei(height))
	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.BindBuffer(gl.ARRAY_BUFFER, verticesArrayBuffer)
	gl.VertexAttribPointer(attrPos, 4, gl.FLOAT, false, 6*4, nil)

	// bind texture - FIX size of vertex

	gl.VertexAttribPointer(attrTexIn, 2, gl.FLOAT, false, 6*4, gl.Void(uintptr(4*4)))

	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, textureBuffer)
	gl.Uniform1i(int32(unifTexture), 0)

	gl.DrawArrays(gl.TRIANGLE_FAN, 0, 4)
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
	if *info {
		printInfo()
	}
	application.Register("render loop", newRenderLoop(FRAMES_PER_SECOND))
	application.InstallSignalHandler(&signal)
	exitCh := make(chan bool, 1)
	application.Run(exitCh)
	<-exitCh
}
