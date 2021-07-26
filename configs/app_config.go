package configs

type AppConfig struct {
	Mousebind string `yaml:"mousebind"`
	Profile   string `yaml:"profile"`
}

func GetDefaultAppConfig() *AppConfig {
	return &AppConfig{
		Mousebind: "Shift-2",
		Profile:   "default",
	}
}
