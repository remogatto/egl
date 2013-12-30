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
		egl.DEPTH_SIZE, 8,
		egl.SURFACE_TYPE, egl.WINDOW_BIT,
		egl.NONE,
	}
	DefaultContextAttributes = []int32{
		egl.CONTEXT_CLIENT_VERSION, 2,
		egl.NONE,
	}
)

func Initialize(configAttr, contextAttr []int32) *platform.EGLState {
	var (
		config           egl.Config
		numConfig        int32
		visualId         int32
		width, height    int32
		dstRect, srcRect egl.VCRect
		nativeWindow     egl.EGLDispmanxWindow
	)
	display := egl.GetDisplay(egl.DEFAULT_DISPLAY)
	if ok := egl.Initialize(egl.Display, nil, nil); !ok {
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

	width, height = egl.GraphicsGetDisplaySize(0)

	dstRect.X = 0
	dstRect.Y = 0
	dstRect.Width = int32(width)
	dstRect.Height = int32(height)

	srcRect.X = 0
	srcRect.Y = 0
	srcRect.Width = int32(width << 16)
	srcRect.Height = int32(height << 16)

	dispman_display := egl.VCDispmanxDisplayOpen(0)
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

	nativeWindow.Element = dispman_element
	nativeWindow.Width = int(width)
	nativeWindow.Height = int(height)
	egl.VCDispmanxUpdateSubmitSync(dispman_update)

	surface := egl.CreateWindowSurface(
		display,
		config,
		egl.NativeWindowType(unsafe.Pointer(&nativeWindow)),
		nil)

	if surface == egl.NO_SURFACE {
		panic("Error in creating EGL surface")
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
