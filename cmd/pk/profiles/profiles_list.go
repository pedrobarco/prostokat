package profiles

import (
	"fmt"
	"prostokat/configs"
	"strings"

	"github.com/spf13/cobra"
)

func NewCmdProfilesList(cfg *configs.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Lists existing named profiles",
		Run: func(cmd *cobra.Command, args []string) {
			parr := cfg.ListProfiles()
			profiles := strings.Join(parr, "\n")
			fmt.Println("= Profiles =")
			fmt.Println(profiles)
		},
	}

	return cmd
}
