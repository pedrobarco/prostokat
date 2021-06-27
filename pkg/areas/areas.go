package areas

import (
	"log"

	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/ewmh"
)

type Area struct {
	X, Y, W, H int
}

func ActiveWorkarea(xu *xgbutil.XUtil) *Area {
	wa, err := ewmh.WorkareaGet(xu)
	if err != nil {
		log.Fatal(err)
	}
	return &Area{
		X: wa[0].X,
		Y: wa[0].Y,
		W: int(wa[0].Width),
		H: int(wa[0].Height),
	}
}
