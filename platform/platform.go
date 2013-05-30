package platform

import "github.com/remogatto/egl"

var (
	Display   egl.Display
	Config    egl.Config
	Context   egl.Context
	Surface   egl.Surface
	NumConfig int32
	VisualId  int32
)
