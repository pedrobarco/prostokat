package profiles

import (
	"fmt"
	"prostokat/configs"

	"github.com/spf13/cobra"
)

var (
	describeProfileCmd = &cobra.Command{
		Use:   "describe",
		Short: "Describes a named profile by listing its properties",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("describe profile")
		},
	}

	listProfilesCmd = &cobra.Command{
		Use:   "list",
		Short: "Lists existing named profiles",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("list profiles")
		},
	}
)

func NewCmdProfiles(cfg *configs.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "profiles",
		Short: "Manage prostokat profiles",
	}

	cmd.AddCommand(NewCmdProfilesActivate(cfg))
	cmd.AddCommand(NewCmdProfilesCreate(cfg))
	cmd.AddCommand(NewCmdProfilesDelete(cfg))
	// cmd.AddCommand(describeProfileCmd)
	// cmd.AddCommand(listProfilesCmd)
	return cmd
}
