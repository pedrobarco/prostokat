package profiles

import (
	"fmt"
	"prostokat/configs"

	"github.com/spf13/cobra"
)

func NewCmdProfilesActivate(cfg *configs.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "activate",
		Short: "Activates an existing named profile",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			profile := args[0]
			err := cfg.LoadProfile(profile)
			if err != nil {
				fmt.Printf("Could not load profile %s: %s \n", profile, err)
				return
			}
			fmt.Printf("Profile %s is now active \n", profile)
		},
	}

	return cmd
}
