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
				fmt.Println("Creating default config...")
				cfg.CreateDefaultConfig()
			} else if hasConfig && force {
				fmt.Println("Forcing config reset...")
				cfg.ResetConfig()
			}

			fmt.Println("All done: pk is ready!")
		},
	}

	cmd.Flags().BoolP("force", "f", true, "forces pk to reset configs")
	viper.BindPFlag("force", cmd.Flags().Lookup("force"))

	return cmd
}
