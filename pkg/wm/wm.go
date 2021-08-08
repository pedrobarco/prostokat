package wm

import (
	"fmt"
	"log"
	"prostokat/pkg/grid"
	"strings"

	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/ewmh"
	"github.com/BurntSushi/xgbutil/mousebind"
	"github.com/BurntSushi/xgbutil/xevent"
	"github.com/BurntSushi/xgbutil/xwindow"
)

type WM struct {
	X         *xgbutil.XUtil
	Grid      *grid.Grid
	Mousebind string
}

func (wm *WM) config() string {
	b := strings.Builder{}
	btnStr := strings.Split(wm.Mousebind, "-")

	fmt.Fprintln(&b, "= Prostokat WM =")
	fmt.Fprint(&b, wm.Grid.String())
	fmt.Fprintf(&b, "Mousebind: %s-Mouse%s\n", btnStr[0], btnStr[1])

	return b.String()
}

func (wm *WM) Init() {
	fmt.Print(wm.config())

	mousebind.Initialize(wm.X)

	// Tilling with mousebind
	cb1 := mousebind.ButtonPressFun(
		func(xu *xgbutil.XUtil, ev xevent.ButtonPressEvent) {
			mx, my := int(ev.RootX), int(ev.RootY)
			log.Printf("event: buttonpress at (%d, %d)", mx, my)
			wm.tileByMPos(mx, my)
		})
	err := cb1.Connect(wm.X, wm.X.RootWin(), wm.Mousebind, false, true)
	if err != nil {
		log.Fatal(err)
	}

	// Snapping with focus
	// TODO: attach listener to all windows
	// TODO: listen for FocusInEvent
	// TODO: window is dragging when ev.mode = 2

	log.Printf("listening for events...")
	xevent.Main(wm.X)
}

func (wm *WM) tileByMPos(mx, my int) {
	xwin, err := ewmh.ActiveWindowGet(wm.X)
	if err != nil {
		log.Println("tileByMPos: no active window")
		return
	}

	types, err := ewmh.WmWindowTypeGet(wm.X, xwin)
	if err != nil {
		types = []string{}
	}
	for _, t := range types {
		if t == "_NET_WM_WINDOW_TYPE_DESKTOP" {
			log.Println("tileByMPos: no active window")
			return
		}
	}

	a, err := wm.Grid.ClosestArea(wm.X, mx, my)
	if err != nil {
		log.Fatalf("tileByMPos: could not get closest area: %s\n", err)
	}

	w := xwindow.New(wm.X, xwin)
	tile(w, a)
}
