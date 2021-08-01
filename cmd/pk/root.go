package main

import (
	"fmt"
	"log"
	"os"
	"path"

	cmdInit "prostokat/cmd/pk/init"
	"prostokat/cmd/pk/start"
	"prostokat/configs"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Version: "0.1.2",
	Use:     "pk",
	Short:   "A polished GNU/Linux tilling utility",
	Long: `Prostokat is a minimal tilling utility built in Go and pk is its CLI version.
Complete documentation is available at http://github.com/pedrobarco/prostokat.`,
}

var afile *configs.ConfigFile
var pfile *configs.ConfigFile
var cfg *configs.Config

func init() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("could not get user home directory: %s \n", err)
	}

	afile = &configs.ConfigFile{
		Name: "config",
		Type: "yaml",
		Path: path.Join(homedir, ".config/prostokat"),
	}

	pfile = &configs.ConfigFile{
		Name: "default",
		Type: afile.Type,
		Path: path.Join(afile.Path, "profiles"),
	}

	cfg = configs.Create(afile, pfile)

	cobra.OnInitialize(loadAppConfig)

	// set version template
	rootCmd.SetVersionTemplate(fmt.Sprintf("v%s\n", rootCmd.Version))

	// init
	rootCmd.AddCommand(cmdInit.NewCmdInit(cfg))

	// start
	rootCmd.AddCommand(start.NewCmdStart(cfg))

	// profiles
	// rootCmd.AddCommand(NewCmdProfiles())

	// utilities
	// rootCmd.AddCommand(versionCmd)
	// rootCmd.AddCommand(infoCmd)
}

func loadAppConfig() {
	viper.SetConfigName(afile.Name)
	viper.SetConfigType(afile.Type)
	viper.AddConfigPath(afile.Path)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Initalizing config...")
			cfg.CreateDefaultConfig()
			viper.ReadInConfig()
			fmt.Println()
		} else {
			panic(fmt.Errorf("Fatal error reading config: %w \n", err))
		}
	}

	var acfg *configs.AppConfig
	if err := viper.Unmarshal(&acfg); err != nil {
		panic(fmt.Errorf("Fatal error loading config: %s \n", err))
	}

	cfg.SetAppConfig(acfg)
	loadProfileConfig(acfg.Profile)
}

func loadProfileConfig(profile string) {
	viper.SetConfigName(profile)
	viper.SetConfigType(pfile.Type)
	viper.AddConfigPath(pfile.Path)

	err := viper.MergeInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error merging config: %w \n", err))
	}

	var pcfg *configs.ProfileConfig
	if err := viper.Unmarshal(&pcfg); err != nil {
		panic(fmt.Errorf("Fatal error loading config: %s \n", err))
	}

	cfg.SetProfileConfig(pcfg)
}
