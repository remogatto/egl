package platform

import (
	"github.com/remogatto/egl"
	"unsafe"
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
	Display   egl.Display
	Config    egl.Config
	Context   egl.Context
	Surface   egl.Surface
	NumConfig int32
	VisualId  int32
)

func Initialize(window egl.NativeWindowType, configAttr, contextAttr []int32) {
	Display = egl.GetDisplay(egl.NativeDisplayType(unsafe.Pointer(nil)))
	if ok := egl.Initialize(Display, nil, nil); !ok {
		egl.LogError(egl.GetError())
	}
	if ok := egl.ChooseConfig(Display, configAttr, &Config, 1, &NumConfig); !ok {
		egl.LogError(egl.GetError())
	}
	if ok := egl.GetConfigAttrib(Display, Config, egl.NATIVE_VISUAL_ID, &VisualId); !ok {
		egl.LogError(egl.GetError())
	}
	egl.BindAPI(egl.OPENGL_ES_API)
	Context = egl.CreateContext(Display, Config, egl.NO_CONTEXT, &contextAttr[0])
	Surface = egl.CreateWindowSurface(Display, Config, window, nil)
	if ok := egl.MakeCurrent(Display, Surface, Surface, Context); !ok {
		egl.LogError(egl.GetError())
	}
	var val int32
	if ok := egl.QuerySurface(Display, &val, egl.WIDTH, Surface); !ok {
		egl.LogError(egl.GetError())
	}
	if ok := egl.QuerySurface(Display, &val, egl.HEIGHT, Surface); !ok {
		egl.LogError(egl.GetError())
	}
	if ok := egl.GetConfigAttrib(Display, Config, egl.SURFACE_TYPE, &val); !ok {
		egl.LogError(egl.GetError())
	}
	if ok := egl.GetConfigAttrib(Display, Config, egl.SURFACE_TYPE, &val); !ok {
		egl.LogError(egl.GetError())
	}
	if (val & egl.WINDOW_BIT) == 0 { panic("No WINDOW_BIT") }
}
