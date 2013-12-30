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

import (
	"log"
	"strconv"
)

type Error struct {
	eglError int32
}

func NewError(eglError int32) *Error {
	return &Error{eglError}
}

func (e Error) Error() string {
	switch e.eglError {
	case NOT_INITIALIZED:
		return "Not Intialized"
	case BAD_ACCESS:
		return "Bad Access"
	case BAD_ALLOC:
		return "Bad Allocate"
	case BAD_ATTRIBUTE:
		return "Bad Attribute"
	case BAD_CONFIG:
		return "Bad Config"
	case BAD_CONTEXT:
		return "Bad Context"
	case BAD_DISPLAY:
		return "Bad Display"
	case BAD_MATCH:
		return "Bad Match"
	case BAD_PARAMETER:
		return "Bad Parameter"
	case BAD_SURFACE:
		return "Bad Surface"
	case BAD_CURRENT_SURFACE:
		return "Bad Current Surface"
	case BAD_NATIVE_PIXMAP:
		return "Bad Native Pixmap"
	case BAD_NATIVE_WINDOW:
		return "Bad Native Window"
	case SUCCESS:
		return "Success"
	default:
		return strconv.Itoa(int(e.eglError))
	}
}

func LogError(msg int32) {
	log.Println(NewError(msg).Error())
}
