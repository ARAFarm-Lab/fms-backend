package main

import (
	"context"
	"fms-backend/internal/infrastructure/config"
	"fmt"
	"os"
	"strings"

	sdkConfig "github.com/ARAFarm-Lab/go-sdk/config"
	"github.com/ARAFarm-Lab/go-sdk/constants"
	"github.com/ARAFarm-Lab/go-sdk/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	envStr := os.Getenv("ENV")
	if envStr == "" {
		envStr = "development"
	}

	env := constants.ENV(strings.ToLower(envStr))

	var appConfig config.AppConfig
	err := sdkConfig.InitializeConfig(env, &appConfig)
	if err != nil {
		log.Fatal(context.Background(), nil, err, "Failed to initialize config")
	}

	log.NewLogger(env, "./log/test")

	db, err := initializeDB(appConfig.DB)
	if err != nil {
		log.Fatal(context.Background(), nil, err, "Failed to connect to the database")
	}

	fmt.Println(db)
}

func initializeDB(config config.DBConfig) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.DSN), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
