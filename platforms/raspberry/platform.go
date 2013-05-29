package platform

import (
	"github.com/remogatto/egl"
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
	Display                     egl.Display
	Config                      egl.Config
	Context                     egl.Context
	Surface                     egl.Surface
	NumConfig                   int32
	VisualId                         int32
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
	Display = egl.GetDisplay(0)
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
	println(Display, Context)
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

	Surface = egl.CreateWindowSurface(
		Display,
		Config,
		egl.NativeWindowType(unsafe.Pointer(&NativeWindow)),
		nil)
	assert(Surface != egl.NO_SURFACE)

	// connect the context to the surface
	result := egl.MakeCurrent(Display, Surface, Surface, Context)
	assert(result)

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
}

