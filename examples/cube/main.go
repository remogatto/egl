package main

import (
	"github.com/remogatto/application"
	"github.com/remogatto/egl/examples/cube/cubelib"
	"image"
	"image/png"
	"log"
	"os"
	"runtime"
	"syscall"
	"time"
)

const FRAMES_PER_SECOND = 30

var (
	signal                sigterm
	currWidth, currHeight int
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
	cubelib.Initialize()

	// Create the 3D world
	world := cubelib.NewWorld()
	world.SetCamera(0.0, 0.0, 5.0)

	cube := cubelib.NewCube()
	cube.AttachTexture(loadImage("marmo.png"))

	world.Attach(cube)
	angle := float32(0.0)
	for {
		select {
		case <-l.pause:
			l.ticker.Stop()
			l.pause <- 0
		case <-l.terminate:
			cubelib.Cleanup()
			l.terminate <- 0
		case <-l.ticker.C:
			angle += 0.05
			cube.RotateY(angle)
			world.Draw()
			cubelib.Swap()
		}
	}
}

func loadImage(filename string) image.Image {
	// Open the file.
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Decode the image.
	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	return img
}

func main() {
	application.Register("render loop", newRenderLoop(FRAMES_PER_SECOND))
	application.InstallSignalHandler(&signal)
	exitCh := make(chan bool, 1)
	application.Run(exitCh)
	<-exitCh
}
