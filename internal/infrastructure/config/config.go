package config

import "github.com/ARAFarm-Lab/go-sdk/constants"

func (config *Config) GetConfig() AppConfig {
	return config.config
}

func (config *Config) IsDevelopment() bool {
	return config.env == constants.DEVELOPMENT
}
