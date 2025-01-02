package config

import "github.com/ARAFarm-Lab/go-sdk/constants"

type Config struct {
	config AppConfig
	env    constants.ENV
}

type AppConfig struct {
	DB DBConfig `yaml:"db"`
}

type DBConfig struct {
	DSN string `yaml:"dsn"`
}
