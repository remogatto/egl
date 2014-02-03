// +build !raspberry, !android

package egl

/*
#cgo pkg-config: egl
#include <EGL/egl.h>
#include <EGL/eglplatform.h>
*/
import "C"
