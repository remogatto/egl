package xorg

import (
	"github.com/remogatto/egl"
	"github.com/remogatto/egl/platform"
)

var (
	DefaultConfigAttributes = []int32{
		egl.RED_SIZE, 8,
		egl.GREEN_SIZE, 8,
		egl.BLUE_SIZE, 8,
		egl.ALPHA_SIZE, 8,
		egl.DEPTH_SIZE, 8,
		egl.RENDERABLE_TYPE,
		egl.OPENGL_ES2_BIT,
		egl.NONE,
	}
	DefaultContextAttributes = []int32{
		egl.CONTEXT_CLIENT_VERSION, 2,
		egl.NONE,
	}
)

func Initialize(window egl.NativeWindowType, configAttr, contextAttr []int32) *platform.EGLState {
	var (
		config        egl.Config
		numConfig     int32
		visualId      int32
		width, height int32
	)
	display := egl.GetDisplay(egl.DEFAULT_DISPLAY)
	if ok := egl.Initialize(display, nil, nil); !ok {
		egl.LogError(egl.GetError())
	}
	if ok := egl.ChooseConfig(display, configAttr, &config, 1, &numConfig); !ok {
		egl.LogError(egl.GetError())
	}
	if ok := egl.GetConfigAttrib(display, config, egl.NATIVE_VISUAL_ID, &visualId); !ok {
		egl.LogError(egl.GetError())
	}
	egl.BindAPI(egl.OPENGL_ES_API)
	context := egl.CreateContext(display, config, egl.NO_CONTEXT, &contextAttr[0])
	surface := egl.CreateWindowSurface(display, config, window, nil)

	var val int32
	if ok := egl.QuerySurface(display, surface, egl.WIDTH, &width); !ok {
		egl.LogError(egl.GetError())
	}
	if ok := egl.QuerySurface(display, surface, egl.HEIGHT, &height); !ok {
		egl.LogError(egl.GetError())
	}
	if ok := egl.GetConfigAttrib(display, config, egl.SURFACE_TYPE, &val); !ok {
		egl.LogError(egl.GetError())
	}
	if ok := egl.GetConfigAttrib(display, config, egl.SURFACE_TYPE, &val); !ok {
		egl.LogError(egl.GetError())
	}
	if (val & egl.WINDOW_BIT) == 0 {
		panic("No WINDOW_BIT")
	}
	return &platform.EGLState{
		Display:           display,
		Config:            config,
		Context:           context,
		Surface:           surface,
		NumConfig:         numConfig,
		VisualId:          visualId,
		ContextAttributes: contextAttr,
		ConfigAttributes:  configAttr,
		SurfaceWidth:      int(width),
		SurfaceHeight:     int(height),
	}
}
