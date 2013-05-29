package egl

/*
#cgo pkg-config: egl
#include <stdlib.h>
#include <EGL/egl.h>
#include <EGL/eglplatform.h>
*/
import "C"
import "unsafe"

type (
	Enum              uint32
	Config            uintptr
	Context           uintptr
	Display           uintptr
	Surface           uintptr
	ClientBuffer      uintptr
	NativeDisplayType unsafe.Pointer
	NativeWindowType  unsafe.Pointer
	NativePixmapType  unsafe.Pointer
)

func goBoolean(n C.EGLBoolean) bool {
	return n == 1
}
func eglBoolean(n bool) C.EGLBoolean {
	var b int
	if n == true {
		b = 1
	}
	return C.EGLBoolean(b)
}

/*
func ProcAdress(proc string) uintptr {

}
*/
