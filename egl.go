/*
Copyright © 2013 mortdeus <mortdeus@gocos2d.org>

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
“Software”), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package egl

/*
#include <EGL/egl.h>
#include <EGL/eglplatform.h>
*/
import "C"
import (
	"unsafe"
)

func Initialize(
	disp Display, major, minor *int32) bool {
	return goBoolean(C.eglInitialize(
		C.EGLDisplay(unsafe.Pointer(disp)),
		(*C.EGLint)(major),
		(*C.EGLint)(minor)))
}
func Terminate(
	disp Display) bool {
	return goBoolean(C.eglTerminate(
		C.EGLDisplay(unsafe.Pointer(disp))))
}
func GetDisplay(
	displayID NativeDisplayType) Display {
	return Display(C.eglGetDisplay(
		C.EGLNativeDisplayType(unsafe.Pointer(displayID))))
}
func QueryString(
	disp Display, name int32) string {
	return C.GoString(C.eglQueryString(
		C.EGLDisplay(unsafe.Pointer(disp)),
		C.EGLint(name)))
}
func DestroySurface(
	disp Display, surface Surface) bool {
	return goBoolean(C.eglDestroySurface(
		C.EGLDisplay(unsafe.Pointer(disp)),
		C.EGLSurface(unsafe.Pointer(surface))))
}
func SwapInterval(
	disp Display, interval int32) bool {
	return goBoolean(C.eglSwapInterval(
		C.EGLDisplay(unsafe.Pointer(disp)),
		C.EGLint(interval)))
}

func DestroyContext(
	disp Display, ctx Context) bool {
	return goBoolean(C.eglDestroyContext(
		C.EGLDisplay(unsafe.Pointer(disp)),
		C.EGLContext(unsafe.Pointer(ctx))))
}
func GetCurrentSurface(readdraw int32) Surface {
	return Surface(C.eglGetCurrentSurface(
		C.EGLint(readdraw)))
}
func QuerySurface(
	disp Display, value *int32,
	attribute int32, surface Surface) bool {
	return goBoolean(C.eglQuerySurface(
		C.EGLDisplay(unsafe.Pointer(disp)),
		C.EGLSurface(unsafe.Pointer(surface)),
		C.EGLint(attribute),
		(*C.EGLint)(value)))
}
func GetConfigs(
	disp Display, configs *Config,
	configSize int32, numConfig *int32) bool {
	return goBoolean(C.eglGetConfigs(
		C.EGLDisplay(unsafe.Pointer(disp)),
		(*C.EGLConfig)(unsafe.Pointer(configs)),
		C.EGLint(configSize),
		(*C.EGLint)(unsafe.Pointer(numConfig))))
}

func GetConfigAttrib(
	disp Display, config Config,
	attribute int32, value *int32) bool {
	return goBoolean(C.eglGetConfigAttrib(
		C.EGLDisplay(unsafe.Pointer(disp)),
		C.EGLConfig(config),
		C.EGLint(attribute),
		(*C.EGLint)(unsafe.Pointer(value))))
}
func ChooseConfig(
	disp Display, attribList []int32, configs *Config,
	configSize int32, numConfig *int32) bool {
	return goBoolean(C.eglChooseConfig(
		C.EGLDisplay(unsafe.Pointer(disp)),
		(*C.EGLint)(&attribList[0]),
		(*C.EGLConfig)(unsafe.Pointer(configs)),
		C.EGLint(configSize),
		(*C.EGLint)(numConfig)))
}
func CreateContext(
	disp Display, config Config,
	shareContext Context, attribList *int32) Context {
	return Context(C.eglCreateContext(
		C.EGLDisplay(unsafe.Pointer(disp)),
		C.EGLConfig(unsafe.Pointer(config)),
		C.EGLContext(unsafe.Pointer(shareContext)),
		(*C.EGLint)(unsafe.Pointer(attribList))))
}

func CreateWindowSurface(
	disp Display, config Config,
	win NativeWindowType, attribList *int32) Surface {
	return Surface(C.eglCreateWindowSurface(
		C.EGLDisplay(unsafe.Pointer(disp)),
		C.EGLConfig(unsafe.Pointer(config)),
		C.EGLNativeWindowType(uintptr(win)),
		(*C.EGLint)(attribList)))
}
func CreatePbufferSurface(
	disp Display, config Config, attribList *int32) Surface {
	return Surface(C.eglCreatePbufferSurface(
		C.EGLDisplay(unsafe.Pointer(disp)),
		C.EGLConfig(unsafe.Pointer(config)),
		(*C.EGLint)(attribList)))
}
func CreatePixmapSurface(
	disp Display, config Config,
	pixmap NativePixmapType, attribList *int32) Surface {
	return Surface(C.eglCreatePixmapSurface(
		C.EGLDisplay(unsafe.Pointer(disp)),
		C.EGLConfig(config),
		C.EGLNativePixmapType(uintptr(pixmap)),
		(*C.EGLint)(attribList)))
}

func CreatePbufferFromClientBuffer(
	disp Display, buftype Enum, config Config,
	buffer ClientBuffer, attribList *int32) Surface {
	return Surface(C.eglCreatePbufferFromClientBuffer(
		C.EGLDisplay(unsafe.Pointer(disp)),
		C.EGLenum(buftype),
		C.EGLClientBuffer(buffer),
		C.EGLConfig(unsafe.Pointer(config)),
		(*C.EGLint)(attribList)))
}
func SurfaceAttrib(
	disp Display, surface Surface,
	attribute int32, value int32) bool {
	return goBoolean(C.eglSurfaceAttrib(
		C.EGLDisplay(unsafe.Pointer(disp)),
		C.EGLSurface(unsafe.Pointer(surface)),
		C.EGLint(attribute),
		C.EGLint(value)))
}
func BindTexImage(
	disp Display, surface Surface, buffer int32) bool {
	return goBoolean(C.eglBindTexImage(
		C.EGLDisplay(unsafe.Pointer(disp)),
		C.EGLSurface(unsafe.Pointer(surface)),
		C.EGLint(buffer)))
}
func ReleaseTexImage(
	disp Display, surface Surface, buffer int32) bool {
	return goBoolean(C.eglReleaseTexImage(
		C.EGLDisplay(unsafe.Pointer(disp)),
		C.EGLSurface(unsafe.Pointer(surface)),
		C.EGLint(buffer)))
}
func MakeCurrent(
	disp Display, draw Surface,
	read Surface, ctx Context) bool {
	return goBoolean(C.eglMakeCurrent(
		C.EGLDisplay(unsafe.Pointer(disp)),
		C.EGLSurface(unsafe.Pointer(draw)),
		C.EGLSurface(unsafe.Pointer(read)),
		C.EGLContext(unsafe.Pointer(ctx))))
}
func QueryContext(
	disp Display, ctx Context,
	attribute int32, value *int32) bool {
	return goBoolean(C.eglQueryContext(
		C.EGLDisplay(unsafe.Pointer(disp)),
		C.EGLContext(unsafe.Pointer(ctx)),
		C.EGLint(attribute),
		(*C.EGLint)(value)))
}
func CopyBuffers(
	disp Display, surface Surface,
	target NativePixmapType) bool {
	return goBoolean(C.eglCopyBuffers(
		C.EGLDisplay(unsafe.Pointer(disp)),
		C.EGLSurface(unsafe.Pointer(surface)),
		C.EGLNativePixmapType(uintptr(target))))
}
func SwapBuffers(
	disp Display, surface Surface) bool {
	return goBoolean(C.eglSwapBuffers(
		C.EGLDisplay(unsafe.Pointer(disp)),
		C.EGLSurface(unsafe.Pointer(surface))))
}
func BindAPI(api Enum) bool {
	return goBoolean(C.eglBindAPI(
		C.EGLenum(api)))
}
func WaitNative(engine int32) bool {
	return goBoolean(C.eglWaitNative(
		C.EGLint(engine)))
}
func QueryAPI() Enum {
	return Enum(C.eglQueryAPI())
}
func WaitClient() bool {
	return goBoolean(C.eglWaitClient())
}
func WaitGL() bool {
	return goBoolean(C.eglWaitGL())
}
func ReleaseThread() bool {
	return goBoolean(C.eglReleaseThread())
}
func GetCurrentDisplay() Display {
	return Display(C.eglGetCurrentDisplay())
}
func GetError() int32 {
	return int32(C.eglGetError())
}
