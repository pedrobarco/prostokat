package frames

import (
	"log"

	"github.com/BurntSushi/xgbutil/xprop"
	"github.com/BurntSushi/xgbutil/xwindow"
)

type Frame struct {
	Left, Right, Top, Bot int
}

func WindowFrames(win *xwindow.Window) *Frame {
	prop, err := xprop.GetProperty(win.X, win.Id, "_GTK_FRAME_EXTENTS")
	if err != nil {
		prop, err = xprop.GetProperty(win.X, win.Id, "_NET_FRAME_EXTENTS")
		if err != nil {
			log.Fatal(err)
		}
	}
	return &Frame{
		int(prop.Value[0]),
		int(prop.Value[4]),
		int(prop.Value[8]),
		int(prop.Value[12]),
	}
}
