package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"

	"github.com/hnamzian/go-mallbots/internal/web"
	"github.com/hnamzian/go-mallbots/internal/grpc"
	"github.com/hnamzian/go-mallbots/internal/pg"
)

type AppConfig struct {
	LogLevel string          `mapstructure:"logLevel"`
	Web      web.WebConfig   `mapstructure:"web"`
	Grpc     grpc.GrpcConfig `mapstructure:"grpc"`
	PG       pg.PGConfig     `mapstructure:"pg"`
}

func InitConfig() (*AppConfig, error) {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		return nil, fmt.Errorf("CONFIG_PATH is not set")
	}
	viper.SetConfigFile(configPath)

	viper.SetDefault("logLevel", "debug")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %s", err)
	}

	var config AppConfig
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config file: %s", err)
	}

	return &config, nil
}
