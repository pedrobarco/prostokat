package grid

import (
	"prostokat/pkg/areas"
	"testing"
)

func TestGrid(t *testing.T) {
	samples := []struct {
		rows     int
		cols     int
		resRows  int
		resCols  int
		resAreas int
	}{
		{rows: 1, cols: 0, resRows: 1, resCols: 1, resAreas: 1},
		{rows: 0, cols: 1, resRows: 1, resCols: 1, resAreas: 1},
		{rows: 1, cols: 1, resRows: 1, resCols: 1, resAreas: 1},
		{rows: 1, cols: 3, resRows: 1, resCols: 3, resAreas: 3},
		{rows: 2, cols: 4, resRows: 2, resCols: 4, resAreas: 8},
	}

	for _, s := range samples {
		g := Create(s.cols, s.rows)
		if g.Cols() != s.resCols {
			t.Errorf("grid %v should have %d cols\n", g, s.resCols)
		}
		if g.Rows() != s.resRows {
			t.Errorf("grid %v should have %d rows\n", g, s.resRows)
		}
		if g.MaxAreas() != s.resAreas {
			t.Errorf("grid %v should have %d areas\n", g, s.resAreas)
		}
	}
}

func TestLayout(t *testing.T) {
	samples := []struct {
		grid    *Grid
		layout  []areas.Area
		success bool
	}{
		{grid: Create(1, 1), layout: []areas.Area{{X: 0, Y: 0, W: 1, H: 1}}, success: true},
		{grid: Create(1, 1), layout: []areas.Area{{X: 0, Y: 0, W: 2, H: 1}}, success: false},
		{grid: Create(2, 1), layout: []areas.Area{{X: 0, Y: 0, W: 2, H: 1}}, success: true},
		{grid: Create(2, 1), layout: []areas.Area{{X: 1, Y: 0, W: 2, H: 1}}, success: false},
		{grid: Create(4, 1), layout: []areas.Area{{X: 0, Y: 0, W: 1, H: 1}, {X: 1, Y: 0, W: 2, H: 1}, {X: 3, Y: 0, W: 1, H: 1}}, success: true},
	}

	for _, s := range samples {
		err := s.grid.CreateLayout(s.layout)
		if err != nil && s.success {
			t.Errorf("layout %v for grid %v should have not returned an error: %s\n", s.layout, s.grid, err)
		} else if err == nil && !s.success {
			t.Errorf("layout %v for grid %v should have returned an error: %s\n", s.layout, s.grid, err)
		}
	}
}
