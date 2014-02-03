// +build android

package android

import (
	"unsafe"
	"github.com/remogatto/egl"
	"github.com/remogatto/egl/platform"
)

// #include <android/native_window.h>
// #cgo LDFLAGS: -landroid
import "C"

var (
	DefaultContextAttributes = []int32{
		egl.CONTEXT_CLIENT_VERSION, 2,
		egl.NONE,
	}

	DefaultConfigAttributes = []int32{
		egl.RED_SIZE, 8,
		egl.GREEN_SIZE, 8,
		egl.BLUE_SIZE, 8,
		egl.DEPTH_SIZE, 8,
		egl.RENDERABLE_TYPE, egl.OPENGL_ES2_BIT,
		egl.SURFACE_TYPE, egl.WINDOW_BIT,
		egl.NONE,
	}
)

func getEGLDisp(disp egl.NativeDisplayType) egl.Display {
	if !egl.BindAPI(egl.OPENGL_ES_API) {
		panic("Error: eglBindAPI() failed")
	}

	egl_dpy := egl.GetDisplay(egl.DEFAULT_DISPLAY)
	if egl_dpy == egl.NO_DISPLAY {
		panic("Error: eglGetDisplay() failed\n")
	}

	var egl_major, egl_minor int32
	if !egl.Initialize(egl_dpy, &egl_major, &egl_minor) {
		panic("Error: eglInitialize() failed\n")
	}
	return egl_dpy
}

func EGLCreateWindowSurface(eglDisp egl.Display, config egl.Config, win egl.NativeWindowType) egl.Surface {
	eglSurf := egl.CreateWindowSurface(eglDisp, config, win, nil)
	if eglSurf == egl.NO_SURFACE {
		panic("Error: eglCreateWindowSurface failed\n")
	}
	return eglSurf
}

func getEGLNativeVisualId(eglDisp egl.Display, config egl.Config) int32 {
	var vid int32
	if !egl.GetConfigAttrib(eglDisp, config, egl.NATIVE_VISUAL_ID, &vid) {
		panic("Error: eglGetConfigAttrib() failed\n")
	}
	return vid
}

func chooseEGLConfig(eglDisp egl.Display, configAttr []int32) egl.Config {
	var config egl.Config
	var num_configs int32
	if !egl.ChooseConfig(eglDisp, configAttr, &config, 1, &num_configs) {
		panic("Error: couldn't get an EGL visual config\n")
	}

	return config
}

func Initialize(win unsafe.Pointer, configAttr, contextAttr []int32) *platform.EGLState {
	eglState := new(platform.EGLState)
	eglState.Display = getEGLDisp(egl.DEFAULT_DISPLAY)
	eglState.Config = chooseEGLConfig(eglState.Display, configAttr)
	eglState.VisualId = getEGLNativeVisualId(eglState.Display, eglState.Config)
	C.ANativeWindow_setBuffersGeometry((*[0]byte)(win), 0, 0, C.int32_t(eglState.VisualId))
	eglState.Surface = EGLCreateWindowSurface(eglState.Display, eglState.Config, egl.NativeWindowType(win))

	egl.BindAPI(egl.OPENGL_ES_API)
	eglState.Context = egl.CreateContext(eglState.Display, eglState.Config, egl.NO_CONTEXT, &contextAttr[0])

	eglState.SurfaceWidth = int(C.ANativeWindow_getWidth((*C.ANativeWindow)(win)))
	eglState.SurfaceHeight = int(C.ANativeWindow_getHeight((*C.ANativeWindow)(win)))

	return eglState
}
