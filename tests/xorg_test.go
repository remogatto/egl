// +build !raspberry

package test

import (
	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/mousebind"
	"github.com/BurntSushi/xgbutil/xevent"
	"github.com/BurntSushi/xgbutil/xwindow"
	"github.com/remogatto/egl"
	"log"
)

const (
	WINDOW_WIDTH  = 640
	WINDOW_HEIGHT = 480
)

var (
	testWin    egl.NativeWindowType
	X          *xgbutil.XUtil
	configAttr = []int32{
		egl.RED_SIZE, 8,
		egl.GREEN_SIZE, 8,
		egl.BLUE_SIZE, 8,
		egl.ALPHA_SIZE, 8,
		egl.DEPTH_SIZE, 8,
		egl.RENDERABLE_TYPE,
		egl.OPENGL_ES2_BIT,
		egl.NONE,
	}
	contextAttr = []int32{
		egl.CONTEXT_CLIENT_VERSION, 2,
		egl.NONE,
	}
)

func newWindow(X *xgbutil.XUtil, width, height int) *xwindow.Window {
	var (
		err error
		win *xwindow.Window
	)
	win, err = xwindow.Generate(X)
	if err != nil {
		log.Fatal(err)
	}
	win.Create(X.RootWin(), 0, 0, width, height,
		xproto.CwBackPixel|xproto.CwEventMask,
		0, xproto.EventMaskButtonRelease)
	win.WMGracefulClose(
		func(w *xwindow.Window) {
			xevent.Detach(w.X, w.Id)
			mousebind.Detach(w.X, w.Id)
			w.Destroy()
			xevent.Quit(X)
		})

	win.Map()

	if err != nil {
		log.Fatal(err)
	}
	return win
}

func openWin(X *xgbutil.XUtil) *xwindow.Window {
	mousebind.Initialize(X)
	win := newWindow(X, WINDOW_WIDTH, WINDOW_HEIGHT)
	go xevent.Main(X)
	return win
}

func initPlatform() {
	var err error
	X, err = xgbutil.NewConn()
	if err != nil {
		log.Fatal(err)
	}
	testWin = egl.NativeWindowType(uintptr(openWin(X).Id))
}
