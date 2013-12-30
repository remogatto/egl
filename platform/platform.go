package platform

import "github.com/remogatto/egl"

type EGLState struct {
	Display egl.Display
	Config egl.Config
	Context egl.Context
	Surface egl.Surface
	NumConfig int32
	VisualId int32
	ContextAttributes []int32
	ConfigAttributes []int32
	SurfaceWidth, SurfaceHeight int
}

