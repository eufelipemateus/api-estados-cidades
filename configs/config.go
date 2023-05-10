package config

import (
	"github.com/spf13/viper"
)

var cfg *config

type config struct {
	API aPiConfig
}

type aPiConfig struct {
	Port       string
	Name       string
	GenerateDB bool
}

func init() {
	viper.SetDefault("api.port", "9000")
}

func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()

	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	cfg = new(config)
	cfg.API = aPiConfig{
		Port: viper.GetString("api.port"),
		Name: viper.GetString("api.name"),
	}

	return nil
}

func GetServerPort() string {
	return cfg.API.Port
}

func GetAppName() string {
	return cfg.API.Name
}
