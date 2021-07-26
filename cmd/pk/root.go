package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"prostokat/configs"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

var rootCmd = &cobra.Command{
	Version: "0.1.1",
	Use:     "pk",
	Short:   "A polished GNU/Linux tilling utility",
	Long: `Prostokat is a minimal tilling utility built in Go.
Pk is Prostokat's CLI version.
Complete documentation is available at http://github.com/pedrobarco/prostokat.`,
}

func init() {
	cobra.OnInitialize(loadConfig)

	// set version template
	rootCmd.SetVersionTemplate(fmt.Sprintf("v%s\n", rootCmd.Version))

	// start
	rootCmd.AddCommand(startCmd)

	// profiles
	// rootCmd.AddCommand(profilesCmd)

	// utilities
	// rootCmd.AddCommand(versionCmd)
	// rootCmd.AddCommand(infoCmd)
}

var acfg configs.AppConfig
var pcfg configs.ProfileConfig

func loadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.config/prostokat")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			initConfig()
			viper.ReadInConfig()
		} else {
			panic(fmt.Errorf("Fatal error reading config: %w \n", err))
		}
	}

	profile := viper.GetString("profile")
	viper.SetConfigName(profile)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.config/prostokat/profiles")

	err := viper.MergeInConfig()
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

func initConfig() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("could not get user home directory: %s \n", err)
	}

	var (
		appConfigDir  = homedir + "/.config/prostokat"
		appConfigName = "config"
		appConfigType = "yaml"
		appConfigFile = appConfigName + "." + appConfigType
		appConfigPath = appConfigDir + "/" + appConfigFile

		profileConfigDir  = appConfigDir + "/profiles"
		profileConfigFile = "default.yaml"
		profileConfigPath = profileConfigDir + "/" + profileConfigFile

		dirMode  = fs.FileMode(0775)
		fileMode = fs.FileMode(0600)
	)

	// delete old configs
	err = os.RemoveAll(appConfigDir)
	if err != nil {
		log.Fatalf("could not delete config folder: %s \n", err)
	}

	// create app config folder
	err = os.Mkdir(appConfigDir, dirMode)
	if err != nil {
		log.Fatalf("could not create config folder: %s \n", err)
	}

	// create profile config folder
	err = os.Mkdir(profileConfigDir, dirMode)
	if err != nil {
		log.Fatalf("could not create profile folder: %s \n", err)
	}

	// create default app config
	appConfig := configs.GetDefaultAppConfig()
	bs, err := yaml.Marshal(appConfig)
	if err != nil {
		log.Fatalf("could not marshal default app config: %s \n", err)
	}
	err = os.WriteFile(appConfigPath, bs, fileMode)
	if err != nil {
		log.Fatalf("could not create default app config: %s \n", err)
	}
	// fmt.Printf("created app config at %s\n", appConfigPath)

	// create default profile config
	profileConfig := configs.GetDefaultProfileConfig()
	bs, err = yaml.Marshal(profileConfig)
	if err != nil {
		log.Fatalf("could not marshal default profile config: %s \n", err)
	}
	err = os.WriteFile(profileConfigPath, bs, fileMode)
	if err != nil {
		log.Fatalf("could not create default profile config: %s \n", err)
	}
	// fmt.Printf("created default profile at %s\n", profileConfigPath)
}
