package configs

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	App         *AppConfig
	Profile     *ProfileConfig
	appFile     *ConfigFile
	profileFile *ConfigFile
}

/*
func (cfg *Config) LoadProfile(profile string) error {
    profileExists := cfg.FilesConfig.hasProfile(profile)
    if !profileExists {
        return fmt.Errorf("error loading profile: profile does not exist")
    }
}
*/

func Create(afile *ConfigFile, pfile *ConfigFile) *Config {
	return &Config{
		appFile:     afile,
		profileFile: pfile,
	}
}

func (c *Config) SetAppConfig(cfg *AppConfig) {
	c.App = cfg
}

func (c *Config) SetProfileConfig(cfg *ProfileConfig) {
	c.Profile = cfg
}

func (c *Config) HasConfig() bool {
	return c.appFile.hasConfigFile()
}

func (c *Config) ResetConfig() {
	c.appFile.deleteConfigFolder()
	c.CreateDefaultConfig()
}

func (c *Config) CreateDefaultConfig() {
	// create app config folder
	c.appFile.createConfigFolder()
	// create default app config
	appConfig := GetDefaultAppConfig()
	bs, err := yaml.Marshal(appConfig)
	if err != nil {
		log.Fatalf("could not marshal default app config: %s \n", err)
	}
	// save app config to file
	c.appFile.createDefaultConfigFile(bs)

	// create profile config folder
	c.profileFile.createConfigFolder()
	// create default profile config
	profileConfig := GetDefaultProfileConfig()
	bs, err = yaml.Marshal(profileConfig)
	if err != nil {
		log.Fatalf("could not marshal default profile config: %s \n", err)
	}
	// save profile config to file
	c.profileFile.createDefaultConfigFile(bs)
	fmt.Println("All done: pk is ready!")
}
