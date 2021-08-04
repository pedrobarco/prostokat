package configs

type ProfileConfig struct {
	Grid    GridConfig     `yaml:"grid"`
	Layouts []LayoutConfig `yaml:"layouts"`
}

type GridConfig struct {
	Cols int `yaml:"cols"`
	Rows int `yaml:"rows"`
}

type LayoutConfig struct {
	Posx   int `yaml:"posx"`
	Posy   int `yaml:"posy"`
	Width  int `yaml:"width"`
	Height int `yaml:"height"`
}

func GetDefaultProfileConfig() *ProfileConfig {
	return &ProfileConfig{
		Grid: GridConfig{
			Cols: 3,
			Rows: 1,
		},
		Layouts: []LayoutConfig{
			{
				Posx:   0,
				Posy:   0,
				Width:  1,
				Height: 1,
			},
			{
				Posx:   1,
				Posy:   0,
				Width:  1,
				Height: 1,
			},
			{
				Posx:   2,
				Posy:   0,
				Width:  1,
				Height: 1,
			},
		},
	}
}
