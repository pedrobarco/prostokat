package grid

import (
	"fmt"
	"prostokat/pkg/areas"

	"github.com/BurntSushi/xgbutil"
)

type Grid struct {
	cols, rows int
}

func sanitize(x int) int {
	if x < 2 {
		return 1
	}
	return x
}

func Create(cols, rows int) *Grid {
	return &Grid{sanitize(cols), sanitize(rows)}
}

func (g *Grid) Cols() int {
	return g.cols
}

func (g *Grid) Rows() int {
	return g.rows
}
func (g *Grid) MaxAreas() int {
	return g.cols * g.rows
}

func (g *Grid) Areas(xu *xgbutil.XUtil) []areas.Area {
	wma := areas.ActiveWorkarea(xu)

	var (
		n  = g.Cols()
		m  = g.Rows()
		dx = 0
		dy = 0
		dw = wma.W / n
		dh = wma.H / m
	)

	a := []areas.Area{}
	for i := 0; i < n; i++ {
		x := dx + (dw * i)
		for j := 0; j < m; j++ {
			y := dy + (dh * j)
			// in a simple column layout Y is always the same
			ta := areas.Area{X: x, Y: y, W: dw, H: dh}
			a = append(a, ta)
		}
	}
	return a
}

func (g *Grid) ClosestArea(xu *xgbutil.XUtil, x, y int) (*areas.Area, error) {
	wma := areas.ActiveWorkarea(xu)

	var (
		n  = g.Cols()
		m  = g.Rows()
		dx = 0
		dy = 0
		dw = wma.W / n
		dh = wma.H / m
	)

	for i := 0; i < n; i++ {
		xmin := dx + (dw * i)
		xmax := xmin + dw
		for j := 0; j < m; j++ {
			ymin := dy + (dh * j)
			ymax := ymin + dh
			// Bounding box
			if xmin <= x && x <= xmax && ymin <= y && y <= ymax {
				return &areas.Area{X: xmin, Y: ymin, W: dw, H: dh}, nil
			}
		}
	}

	return nil, fmt.Errorf("grid: no area containing (%d, %d)\n", x, y)
}
