package main

import (
	"fmt"
	"prostokat/pkg/areas"
	"prostokat/pkg/grid"
	"strconv"
	"strings"
)

type GridFlag struct {
	grid *grid.Grid
}

func (f *GridFlag) String() string {
	return fmt.Sprint(f.grid)
}

func (f *GridFlag) Set(value string) error {
	if f.grid != nil {
		return fmt.Errorf("the grid flag is already set")
	}

	splitter := strings.Split(value, "x")
	if len(splitter) != 2 {
		return fmt.Errorf("invalid grid format: should be \"<N>x<M>\"")
	}

	var (
		n, _ = strconv.Atoi(splitter[0])
		m, _ = strconv.Atoi(splitter[1])
	)
	f.grid = grid.Create(n, m)
	return nil
}

type LayoutFlag struct {
	areas []*areas.Area
}

func (f *LayoutFlag) String() string {
	b := strings.Builder{}
	for _, l := range f.areas {
		fmt.Fprintf(&b, "%d,%d %dx%d", l.X, l.Y, l.W, l.H)
	}
	return b.String()
}

func (f *LayoutFlag) Set(value string) error {
	splitter := strings.Split(value, ",")
	if len(splitter) != 2 {
		return fmt.Errorf("invalid area format: area should be \"<X>,<Y> <W>x<H>\"")
	}

	pos := strings.Split(splitter[0], "x")
	if len(pos) != 2 {
		return fmt.Errorf("invalid position format: use the following format \"<X>,<Y>\"")
	}

	size := strings.Split(splitter[1], "x")
	if len(size) != 2 {
		return fmt.Errorf("invalid size format: use the following format \"<W>x<H>\"")
	}

	var (
		x, _ = strconv.Atoi(pos[0])
		y, _ = strconv.Atoi(pos[1])
		w, _ = strconv.Atoi(size[0])
		h, _ = strconv.Atoi(size[1])
	)
	f.areas = append(f.areas, &areas.Area{X: x, Y: y, W: w, H: h})
	return nil
}
