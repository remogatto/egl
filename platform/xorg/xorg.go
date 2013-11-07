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

func Initialize(window egl.NativeWindowType, configAttr, contextAttr []int32) {
	platform.Display = egl.GetDisplay(egl.DEFAULT_DISPLAY)
	if ok := egl.Initialize(platform.Display, nil, nil); !ok {
		egl.LogError(egl.GetError())
	}
	if ok := egl.ChooseConfig(platform.Display, configAttr, &platform.Config, 1, &platform.NumConfig); !ok {
		egl.LogError(egl.GetError())
	}
	if ok := egl.GetConfigAttrib(platform.Display, platform.Config, egl.NATIVE_VISUAL_ID, &platform.VisualId); !ok {
		egl.LogError(egl.GetError())
	}
	egl.BindAPI(egl.OPENGL_ES_API)
	platform.Context = egl.CreateContext(platform.Display, platform.Config, egl.NO_CONTEXT, &contextAttr[0])
	platform.Surface = egl.CreateWindowSurface(platform.Display, platform.Config, window, nil)
	if ok := egl.MakeCurrent(platform.Display, platform.Surface, platform.Surface, platform.Context); !ok {
		egl.LogError(egl.GetError())
	}
	var val int32
	if ok := egl.QuerySurface(platform.Display, platform.Surface, egl.WIDTH, &val); !ok {
		egl.LogError(egl.GetError())
	}
	if ok := egl.QuerySurface(platform.Display, platform.Surface, egl.HEIGHT, &val); !ok {
		egl.LogError(egl.GetError())
	}
	if ok := egl.GetConfigAttrib(platform.Display, platform.Config, egl.SURFACE_TYPE, &val); !ok {
		egl.LogError(egl.GetError())
	}
	if ok := egl.GetConfigAttrib(platform.Display, platform.Config, egl.SURFACE_TYPE, &val); !ok {
		egl.LogError(egl.GetError())
	}
	if (val & egl.WINDOW_BIT) == 0 {
		panic("No WINDOW_BIT")
	}
}
