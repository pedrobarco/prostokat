package main

import (
	"log"
	"prostokat/pkg/grid"
	"prostokat/pkg/wm"

	"github.com/BurntSushi/xgbutil"
)

func main() {
	/*
		Workflow:
		- connect to X
		- initalize event handler
		- on mouse button press:
			- get mouse pos
			- get active window
			- use mouse pos to get closest area
			- move and tile window
		- on window drag:
			- listen for window movement
			- calculate window position
			- calculate snapping areas
			- if in snapping area
			- move and tile window
	*/

	X, err := xgbutil.NewConn()
	if err != nil {
		log.Fatal(err)
	}

	wm := wm.WM{
		X:         X,
		Grid:      grid.Create(3, 1),
		Layout:    []*grid.Grid{},
		Mousebind: "Shift-2",
	}

	wm.Init()
}
