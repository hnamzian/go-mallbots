package config

import (
	"github.com/spf13/viper"

	"github.com/hnamzian/go-mallbots/internal/grpc"
	"github.com/hnamzian/go-mallbots/internal/http"
)

type AppConfig struct {
	LogLevel string          `mapstructure:"logLevel"`
	Http     http.HttpConfig `mapstructure:"http"`
	Grpc     grpc.GrpcConfig `mapstructure:"grpc"`
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
