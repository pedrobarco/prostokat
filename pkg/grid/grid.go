package grid

import (
	"fmt"
	"prostokat/pkg/areas"
	"strings"

	"github.com/BurntSushi/xgbutil"
)

type Layout = []*areas.Area

type Grid struct {
	cols, rows int
	layout     Layout
}

func sanitize(x int) int {
	if x < 2 {
		return 1
	}
	return x
}

func Create(cols, rows int) *Grid {
	return &Grid{sanitize(cols), sanitize(rows), nil}
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

func (g *Grid) CreateLayout(arr []*areas.Area) error {
	var (
		maxCols = g.cols
		maxRows = g.rows
	)

	for _, a := range arr {
		var (
			maxColsLeft = maxCols - a.X
			maxRowsLeft = maxRows - a.Y
		)

		if a.W*a.H == 0 {
			return fmt.Errorf("invalid layout: area %v must be > 0\n", a)
		}
		if a.W > maxColsLeft {
			return fmt.Errorf("invalid layout: grid has %d cols\n", maxCols)
		}
		if a.H > maxRowsLeft {
			return fmt.Errorf("invalid layout: grid has %d rows\n", maxRows)
		}
	}

	g.layout = arr
	return nil
}

func (g *Grid) String() string {
	b := strings.Builder{}
	fmt.Fprintf(&b, "Grid: %dx%d\n", g.cols, g.rows)
	for i, l := range g.layout {
		fmt.Fprintf(&b, "Layout %d: (%d, %d) %dx%d\n", i, l.X, l.Y, l.W, l.H)
	}
	fmt.Fprintln(&b, g.printLayout())
	return b.String()
}

func (g *Grid) printLayout() string {
	b := strings.Builder{}
	for i := 0; i < g.rows; i++ {
		rstr := ""
		for j := 0; j < g.cols; j++ {
			s := " "
			for _, l := range g.layout {
				if l.X == j && l.Y == i {
					s += "|"
				}
			}
			rstr += s
		}
		rstr += " |"
		rstr = strings.Trim(rstr, " ")
		fmt.Fprint(&b, rstr)
	}
	return b.String()
}

func (g *Grid) Layout(xu *xgbutil.XUtil) Layout {
	wma := areas.ActiveWorkarea(xu)

	var (
		n  = g.Cols()
		m  = g.Rows()
		dx = 0
		dy = 0
		dw = wma.W / n
		dh = wma.H / m
	)

	a := []*areas.Area{}
	for i := 0; i < n; i++ {
		x := dx + (dw * i)
		for j := 0; j < m; j++ {
			y := dy + (dh * j)
			ta := areas.Area{X: x, Y: y, W: dw, H: dh}
			a = append(a, &ta)
		}
	}
	return a
}

func (g *Grid) ClosestArea(xu *xgbutil.XUtil, x, y int) (*areas.Area, error) {
	wma := areas.ActiveWorkarea(xu)

	var (
		n  = g.Cols()
		m  = g.Rows()
		dw = wma.W / n
		dh = wma.H / m
	)

	for _, l := range g.layout {
		w := dw * l.W
		h := dh * l.H
		xmin := dw * l.X
		xmax := xmin + w
		ymin := dh * l.Y
		ymax := ymin + h
		if xmin <= x && x <= xmax && ymin <= y && y <= ymax {
			return &areas.Area{X: xmin, Y: ymin, W: w, H: h}, nil
		}
	}

	return nil, fmt.Errorf("grid: no area containing (%d, %d)\n", x, y)
}
