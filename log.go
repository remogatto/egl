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
