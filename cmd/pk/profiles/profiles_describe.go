package profiles

import (
	"fmt"
	"prostokat/configs"

	"github.com/spf13/cobra"
)

func NewCmdProfilesDescribe(cfg *configs.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "describe",
		Short: "Describes a named profile by listing its properties",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			profile := args[0]

			pcfg, err := cfg.GetProfile(profile)
			if err != nil {
				fmt.Printf("Could not create profile %s: %s\n", profile, err)
				return
			}

			fmt.Printf("= %s = \n", profile)
			fmt.Println(string(pcfg))
		},
	}

	return cmd
}
