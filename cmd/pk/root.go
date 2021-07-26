package main

import (
	"fmt"
	"prostokat/configs"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Version: "0.1.0",
	Use:     "pk",
	Short:   "A polished GNU/Linux tilling utility",
	Long: `Prostokat is a minimal tilling utility built in Go.
Pk is Prostokat's CLI version.
Complete documentation is available at http://github.com/pedrobarco/prostokat.`,
}

var acfg configs.AppConfig
var pcfg configs.ProfileConfig

func init() {
	cobra.OnInitialize(initConfig)

	// set version template
	rootCmd.SetVersionTemplate(fmt.Sprintf("v%s\n", rootCmd.Version))

	// init
	rootCmd.AddCommand(initCmd)

	// start
	rootCmd.AddCommand(startCmd)

	// profiles
	// rootCmd.AddCommand(profilesCmd)

	// utilities
	// rootCmd.AddCommand(versionCmd)
	// rootCmd.AddCommand(infoCmd)
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.config/prostokat")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	profile := viper.GetString("profile")
	viper.SetConfigName(profile)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.config/prostokat/profiles")

	err = viper.MergeInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error merging config: %w \n", err))
	}

	if err := viper.Unmarshal(&acfg); err != nil {
		panic(fmt.Errorf("Fatal error loading config: %s \n", err))
	}

	if err := viper.Unmarshal(&pcfg); err != nil {
		panic(fmt.Errorf("Fatal error loading config: %s \n", err))
	}

}
