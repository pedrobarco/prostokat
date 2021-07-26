package main

import (
	"fmt"
	"prostokat/pkg/areas"
	"prostokat/pkg/grid"
	"prostokat/pkg/wm"

	"github.com/BurntSushi/xgbutil"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start prostokat window manager",
	Long:  "Starts running prostokat window manager.",
	Run: func(cmd *cobra.Command, args []string) {
		x, err := xgbutil.NewConn()
		if err != nil {
			panic(fmt.Errorf("Fatal error creating X conn: %s \n", err))
		}

		g := grid.Create(viper.GetInt("cols"), viper.GetInt("rows"))

		var l grid.Layout
		for _, a := range pcfg.Layouts {
			l = append(l, &areas.Area{X: a.Posx, Y: a.Posy, W: a.Width, H: a.Height})
		}

		err = g.CreateLayout(l)
		if err != nil {
			panic(fmt.Errorf("Fatal error parsing layouts: %s \n", err))
		}

		wm := wm.WM{
			X:         x,
			Grid:      g,
			Mousebind: acfg.Mousebind,
		}

		wm.Init()
	},
}

func init() {
	startCmd.Flags().IntP("cols", "c", pcfg.Grid.Cols, "Columns for tilling grid")
	viper.BindPFlag("cols", startCmd.Flags().Lookup("cols"))

	startCmd.Flags().IntP("rows", "r", pcfg.Grid.Rows, "Rows for tilling grid")
	viper.BindPFlag("rows", startCmd.Flags().Lookup("rows"))
}
