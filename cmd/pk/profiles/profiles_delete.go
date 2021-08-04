package profiles

import (
	"fmt"
	"prostokat/configs"

	"github.com/spf13/cobra"
)

func NewCmdProfilesDelete(cfg *configs.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Deletes a named profile",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			profile := args[0]

			err := cfg.DeleteProfile(profile)
			if err != nil {
				fmt.Printf("Could not delete profile %s: %s\n", profile, err)
				return
			}

			fmt.Printf("Profile %s was successfully deleted \n", profile)
		},
	}

	return cmd
}
