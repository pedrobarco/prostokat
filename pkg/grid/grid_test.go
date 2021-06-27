package grid

import (
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

	for _, sample := range samples {
		g := Create(sample.cols, sample.rows)
		if g.Rows() != sample.resRows {
			t.Errorf("grid %v should have %d rows\n", g, sample.resRows)
		}
		if g.Cols() != sample.resCols {
			t.Errorf("grid %v should have %d cols\n", g, sample.resCols)
		}
		if g.MaxAreas() != sample.resAreas {
			t.Errorf("grid %v should have %d areas\n", g, sample.resAreas)
		}
	}
}
