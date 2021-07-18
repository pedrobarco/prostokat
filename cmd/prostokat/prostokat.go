package main

import (
	"flag"
	"fmt"
	"log"
	"prostokat/pkg/wm"

	"github.com/BurntSushi/xgbutil"
)

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
func main() {
	var garg GridFlag
	var larg LayoutFlag

	flag.Var(&garg, "grid", "grid to tile, e.g. \"3x1\" (size[N x M] cols x rows)")
	flag.Var(&larg, "layout", "layout in grid, e.g. \"0x0,1x1\" (pos[X x Y],size[W x H])")

	var flagAlias = map[string]string{
		"grid":   "g",
		"layout": "l",
	}

	for from, to := range flagAlias {
		flagSet := flag.Lookup(from)
		flag.Var(flagSet.Value, to, fmt.Sprintf("alias to %s", flagSet.Name))
	}

	flag.Parse()

	X, err := xgbutil.NewConn()
	if err != nil {
		log.Fatal(err)
	}

	g := garg.grid
	err = g.CreateLayout(larg.areas)
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
