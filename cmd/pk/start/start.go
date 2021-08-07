package start

import (
	"fmt"
	"prostokat/configs"
	"prostokat/pkg/areas"
	"prostokat/pkg/grid"
	"prostokat/pkg/wm"

	"github.com/BurntSushi/xgbutil"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewCmdStart(cfg *configs.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start prostokat window manager",
		Long:  "Starts running prostokat window manager.",
		Run: func(cmd *cobra.Command, args []string) {
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
			x, err := xgbutil.NewConn()
			if err != nil {
				panic(fmt.Errorf("Fatal error creating X conn: %s \n", err))
			}

			/*
			   if viper.IsSet("profile") {
			       profile := viper.GetString("profile")
			       fmt.Println(profile)
			       // TODO: implement LoadProfile logic in configs
			       // cfg.LoadProfile(profile)
			   }
			*/

			var l grid.Layout
			for _, a := range cfg.Profile.Layouts {
				l = append(l, &areas.Area{X: a.Posx, Y: a.Posy, W: a.Width, H: a.Height})
			}

			g := grid.Create(cfg.Profile.Grid.Cols, cfg.Profile.Grid.Rows)
			err = g.CreateLayout(l)
			if err != nil {
				panic(fmt.Errorf("Fatal error parsing layouts: %s \n", err))
			}

			wm := wm.WM{
				X:         x,
				Grid:      g,
				Mousebind: cfg.App.Mousebind,
			}

			wm.Init()
		},
	}

	cmd.Flags().StringP("profile", "p", "default", "Profile to be used by pk")
	err := viper.BindPFlag("profile", cmd.Flags().Lookup("profile"))
	if err != nil {
		panic(fmt.Errorf("Unexpected error while parsing flags: %s \n", err))
	}

	return cmd
}
