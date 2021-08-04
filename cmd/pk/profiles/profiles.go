package profiles

import (
	"prostokat/configs"

	"github.com/spf13/cobra"
)

func NewCmdProfiles(cfg *configs.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "profiles",
		Short: "Manage prostokat profiles",
	}

	cmd.AddCommand(NewCmdProfilesActivate(cfg))
	cmd.AddCommand(NewCmdProfilesCreate(cfg))
	cmd.AddCommand(NewCmdProfilesDelete(cfg))
	cmd.AddCommand(NewCmdProfilesEdit(cfg))
	cmd.AddCommand(NewCmdProfilesDescribe(cfg))
	cmd.AddCommand(NewCmdProfilesList(cfg))

	return cmd
}
