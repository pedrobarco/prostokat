package wm

import (
	"log"
	"prostokat/pkg/areas"
	"prostokat/pkg/frames"

	"github.com/BurntSushi/xgbutil/xwindow"
)

func tile(win *xwindow.Window, a *areas.Area) {
	f := frames.WindowFrames(win)

	dx := a.X - f.Left
	dy := a.Y
	w := a.W + (f.Left + f.Right)
	h := a.H + (f.Bot + f.Top)

	log.Printf("tile: window %d - (%d,%d) %dx%d\n", win.Id, dx, dy, w, h)
	win.MoveResize(dx, dy, w, h)
}
