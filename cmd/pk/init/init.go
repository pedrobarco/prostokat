package init

import (
	"fmt"
	"prostokat/configs"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewCmdInit(cfg *configs.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize or reinitialize pk",
		Long: `This command wipes pk config and profiles and creates a new config and default profile. 
Ideal for a fresh start!`,
		Run: func(cmd *cobra.Command, args []string) {
			hasConfig := cfg.HasConfig()

			force := false
			if viper.IsSet("force") {
				force = viper.GetBool("force")
			}

			if !hasConfig {
				cfg.CreateDefaultConfig()
				return
			}

			if hasConfig && force {
				fmt.Println("Forcing config reset...")
				cfg.ResetConfig()
				return
			}

			fmt.Println("Config folder already exists: use -f/--force flag to reinitialize pk")
		},
	}

	cmd.Flags().BoolP("force", "f", true, "forces to recreate configs")
	viper.BindPFlag("force", cmd.Flags().Lookup("force"))

	return cmd
}
