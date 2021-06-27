package winutil

import (
	"fmt"
	"log"

	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/ewmh"
	"github.com/BurntSushi/xgbutil/xwindow"
)

func ListWindowNames(xu *xgbutil.XUtil) {
	cl, err := ewmh.ClientListGet(xu)
	if err != nil {
		log.Fatal(err)
	}

	for _, id := range cl {
		name, err := ewmh.WmNameGet(xu, id)
		if err != nil {
			log.Printf("Could not get window %d: %s",
				id, err)
			continue
		}

		fmt.Printf("%s %d\n", name, id)
	}
}

func WindowDetails(w *xwindow.Window) {
	fmt.Println()
	fmt.Printf("Window %d\n", w.Id)
	geo, err := w.DecorGeometry()
	if err != nil {
		log.Fatalf("could not get geometry: %s", err)
	}
	fmt.Printf("(%d, %d) %dx%d\n", geo.X(), geo.Y(), geo.Width(), geo.Height())
	fmt.Println()
}
