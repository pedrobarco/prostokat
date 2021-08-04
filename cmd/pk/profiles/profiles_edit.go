package profiles

import (
	"fmt"
	"os/exec"
	"prostokat/configs"

	"github.com/spf13/cobra"
)

func NewCmdProfilesEdit(cfg *configs.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "edit",
		Short: "Opens a named profile in your favourite editor",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			profile := args[0]
			if hasProfile := cfg.HasProfile(profile); !hasProfile {
				fmt.Printf("Could not edit profile %s: profile does not exist \n", profile)
				return
			}

			file := cfg.GetProfileFile(profile)
			editCmd := exec.Command("xdg-open", file)
			err := editCmd.Run()
			if err != nil {
				panic(fmt.Errorf("Could not edit profile: %s \n", err))
			}
		},
	}

	return cmd
}
