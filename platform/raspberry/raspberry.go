package raspberry

import (
	"github.com/remogatto/egl"
	"github.com/remogatto/egl/platform"
	"log"
	"runtime/debug"
	"unsafe"
)

var (
	DefaultConfigAttributes = []int32{
		egl.RED_SIZE, 8,
		egl.GREEN_SIZE, 8,
		egl.BLUE_SIZE, 8,
		egl.ALPHA_SIZE, 8,
		egl.SURFACE_TYPE, egl.WINDOW_BIT,
		egl.NONE,
	}
	DefaultContextAttributes = []int32{
		egl.CONTEXT_CLIENT_VERSION, 2,
		egl.NONE,
	}
	dstRect, srcRect            egl.VCRect
	NativeWindow                egl.EGLDispmanxWindow
	screen_width, screen_height uint32
)

func assert(cond bool) {
	if !cond {
		debug.PrintStack()
		panic("Assertion failed!")
	}
}

func Initialize(configAttr, contextAttr []int32) {
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
	screen_width, screen_height = egl.GraphicsGetDisplaySize(0)
	log.Printf("Display size W: %d H: %d\n", screen_width, screen_height)

	dstRect.X = 0
	dstRect.Y = 0
	dstRect.Width = int32(screen_width)
	dstRect.Height = int32(screen_height)

	srcRect.X = 0
	srcRect.Y = 0
	srcRect.Width = int32(screen_width << 16)
	srcRect.Height = int32(screen_height << 16)

	dispman_display := egl.VCDispmanxDisplayOpen(0 /* LCD */)
	dispman_update := egl.VCDispmanxUpdateStart(0)

	dispman_element := egl.VCDispmanxElementAdd(
		dispman_update,
		dispman_display,
		0, /*layer */
		&dstRect,
		0, /*src */
		&srcRect,
		egl.DISPMANX_PROTECTION_NONE,
		nil, /*alpha */
		nil, /*clamp */
		0 /*transform */)

	NativeWindow.Element = dispman_element
	NativeWindow.Width = int(screen_width)
	NativeWindow.Height = int(screen_height)
	egl.VCDispmanxUpdateSubmitSync(dispman_update)

	platform.Surface = egl.CreateWindowSurface(
		platform.Display,
		platform.Config,
		egl.NativeWindowType(unsafe.Pointer(&NativeWindow)),
		nil)
	assert(platform.Surface != egl.NO_SURFACE)

	// connect the context to the surface
	result := egl.MakeCurrent(platform.Display, platform.Surface, platform.Surface, platform.Context)
	assert(result)

	var val int32
	if ok := egl.QuerySurface(platform.Display, &val, egl.WIDTH, platform.Surface); !ok {
		egl.LogError(egl.GetError())
	}

	if ok := egl.QuerySurface(platform.Display, &val, egl.HEIGHT, platform.Surface); !ok {
		egl.LogError(egl.GetError())
	}
	if ok := egl.GetConfigAttrib(platform.Display, platform.Config, egl.SURFACE_TYPE, &val); !ok {
		egl.LogError(egl.GetError())
	}
}

