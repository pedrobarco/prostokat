package main

import (
	"log"
	"prostokat/pkg/areas"
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

	g := grid.Create(3, 1)
	l1 := areas.Area{X: 0, Y: 0, W: 1, H: 1}
	l2 := areas.Area{X: 1, Y: 0, W: 1, H: 1}
	l3 := areas.Area{X: 2, Y: 0, W: 1, H: 1}
	g.CreateLayout([]areas.Area{l1, l2, l3})
	if err != nil {
		log.Fatalf("wm setup failed: %s\n", err)
	}

	wm := wm.WM{
		X:         X,
		Grid:      g,
		Mousebind: "Shift-2",
	}

	wm.Init()
}
