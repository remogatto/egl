// +build !raspberry

package cubelib

import (
	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/mousebind"
	"github.com/BurntSushi/xgbutil/xevent"
	"github.com/BurntSushi/xgbutil/xwindow"
	"github.com/remogatto/application"
	"log"
)

func newWindow(X *xgbutil.XUtil, reshaper Reshaper) *xwindow.Window {
	var (
		err error
		win *xwindow.Window
	)
	win, err = xwindow.Generate(X)
	if err != nil {
		log.Fatal(err)
	}
	win.Create(X.RootWin(), 0, 0, INITIAL_WINDOW_WIDTH, INITIAL_WINDOW_HEIGHT,
		xproto.CwBackPixel|xproto.CwEventMask,
		0, xproto.EventMaskButtonRelease)
	win.WMGracefulClose(
		func(w *xwindow.Window) {
			xevent.Detach(w.X, w.Id)
			mousebind.Detach(w.X, w.Id)
			// w.Destroy()
			xevent.Quit(X)
			application.Exit()
		})

	// In order to get ConfigureNotify events, we must listen to the window
	// using the 'StructureNotify' mask.
	win.Listen(xproto.EventMaskStructureNotify)

	win.Map()

	xevent.ConfigureNotifyFun(
		func(X *xgbutil.XUtil, ev xevent.ConfigureNotifyEvent) {
			reshaper.Resize(int(ev.Width), int(ev.Height))
		}).Connect(X, win.Id)

	// err = mousebind.ButtonReleaseFun(
	// 	func(X *xgbutil.XUtil, ev xevent.ButtonReleaseEvent) {
	// 		newWindow(X)
	// 	}).Connect(X, win.Id, "1", false, false)

	if err != nil {
		log.Fatal(err)
	}
	return win
}
