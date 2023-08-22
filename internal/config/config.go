package config

import (
	"github.com/spf13/viper"
)

type AppConfig struct {
	LogLevel string `mapstructure:"logLevel"`
}

func InitConfig() *AppConfig {
	viper.SetConfigFile("config.yaml")

	viper.SetDefault("logLevel", "debug")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	var config AppConfig
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}

	return &config
}
