// +build !raspberry

package cubelib

import (
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/mousebind"
	"github.com/BurntSushi/xgbutil/xevent"
	"github.com/remogatto/egl"
	"github.com/remogatto/egl/platform/xorg"
	gl "github.com/remogatto/opengles2"
	"log"
)

const (
	INITIAL_WINDOW_WIDTH  = 640
	INITIAL_WINDOW_HEIGHT = 480
)

var (
	X        *xgbutil.XUtil
	reshaper XorgReshaper
)

type XorgReshaper struct {
	width, height int
}

func (r XorgReshaper) Resize(width, height int) {
	gl.Viewport(0, 0, gl.Sizei(width), gl.Sizei(height))
	r.width, r.height = width, height
}
func (r XorgReshaper) Width() int  { return r.width }
func (r XorgReshaper) Height() int { return r.height }

func Initialize() {
	X, err := xgbutil.NewConn()
	if err != nil {
		log.Fatal(err)
	}
	mousebind.Initialize(X)
	xWindow := newWindow(X, reshaper)
	go xevent.Main(X)
	xorg.Initialize(
		egl.NativeWindowType(uintptr(xWindow.Id)),
		xorg.DefaultConfigAttributes,
		xorg.DefaultContextAttributes)

	reshaper.Resize(INITIAL_WINDOW_WIDTH, INITIAL_WINDOW_HEIGHT)
}
