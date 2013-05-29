package egl

import "log"
import "strconv"

func LogError(msg int32) {
	log.SetPrefix("[EGL] Error: ")
	var s string
	switch msg {
	case NOT_INITIALIZED:
		s = "Not Intialized"
	case BAD_ACCESS:
		s = "Bad Access"
	case BAD_ALLOC:
		s = "Bad Allocate"
	case BAD_ATTRIBUTE:
		s = "Bad Attribute"
	case BAD_CONFIG:
		s = "Bad Config"
	case BAD_CONTEXT:
		s = "Bad Context"
	case BAD_DISPLAY:
		s = "Bad Display"
	case BAD_MATCH:
		s = "Bad Match"
	case BAD_PARAMETER:
		s = "Bad Parameter"
	case BAD_SURFACE:
		s = "Bad Surface"
	case BAD_CURRENT_SURFACE:
		s = "Bad Current Surface"
	case BAD_NATIVE_PIXMAP:
		s = "Bad Native Pixmap"
	case BAD_NATIVE_WINDOW:
		s = "Bad Native Window"
	case SUCCESS:
		s = "Success"
	default:
		s = strconv.Itoa(int(msg))
	}
	log.Println(s)
	panic("panicked!")
}
