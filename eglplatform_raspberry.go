// +build raspberry

package egl

// The native window
type EGLDispmanxWindow struct {
	Element       DispmanxElementHandle
	Width, Height int
}
