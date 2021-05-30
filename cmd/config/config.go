package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	ListenAddress string `json:"listenAddress",yaml:"listenAddress"`
}

func Load() (*Config, error) {
	viper.SetConfigFile("./cmd/config/config.yaml")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
