package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"prostokat/configs"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize and configure prostokat",
	Long:  "Launches an interactive Getting Started workflow for the prostokat CLI",
	Run: func(cmd *cobra.Command, args []string) {
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
		fmt.Printf("created app config at %s\n", appConfigPath)

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
		fmt.Printf("created default profile at %s\n", profileConfigPath)
	},
}
