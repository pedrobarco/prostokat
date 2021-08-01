package profiles

import (
	"fmt"
	"prostokat/configs"

	"github.com/spf13/cobra"
)

func NewCmdProfilesCreate(cfg *configs.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Creates a new named profile",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			profile := args[0]

			err := cfg.CreateProfile(profile)
			if err != nil {
				fmt.Printf("Could not create profile %s: %s\n", profile, err)
				return
			}

			fmt.Printf("Profile %s was successfully created \n", profile)
		},
	}

	return cmd
}
