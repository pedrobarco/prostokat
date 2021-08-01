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
	c.appFile.saveConfig(bs)

	// create profile config folder
	c.profileFile.createConfigFolder()
	// create default profile config
	profileConfig := GetDefaultProfileConfig()
	bs, err = yaml.Marshal(profileConfig)
	if err != nil {
		log.Fatalf("could not marshal default profile config: %s \n", err)
	}
	// save profile config to file
	c.profileFile.saveConfig(bs)
}

func (cfg *Config) HasProfile(profile string) bool {
	return cfg.profileFile.hasConfigFileByName(profile)
}

func (cfg *Config) LoadProfile(profile string) error {
	// check if profile exists
	if profileExists := cfg.HasProfile(profile); !profileExists {
		return fmt.Errorf("error loading profile: profile does not exist")
	}
	// set new profile
	cfg.App.Profile = profile
	// convert to yaml
	bs, err := yaml.Marshal(cfg.App)
	if err != nil {
		return fmt.Errorf("error marshalling profile config: %s \n", err)
	}
	// save to file
	cfg.appFile.saveConfig(bs)
	return nil
}

func (cfg *Config) CreateProfile(profile string) error {
	// check if profile exists
	if profileExists := cfg.HasProfile(profile); profileExists {
		return fmt.Errorf("error creating profile: profile already exists")
	}
	// convert to yaml
	profileConfig := GetDefaultProfileConfig()
	bs, err := yaml.Marshal(profileConfig)
	if err != nil {
		return fmt.Errorf("error marshalling profile config: %s \n", err)
	}
	// create new profile config
	cfg.profileFile.saveConfigToFile(profile, bs)
	return nil
}

func (cfg *Config) DeleteProfile(profile string) error {
	// check if profile is active
	if profileIsActive := profile == cfg.App.Profile; profileIsActive {
		return fmt.Errorf("error deleting profile: cannot delete an active profile")
	}
	// check if profile exists
	if profileExists := cfg.HasProfile(profile); !profileExists {
		return fmt.Errorf("error deleting profile: profile does not exists")
	}
	// delete profile config
	cfg.profileFile.deleteConfig(profile)
	return nil
}
