package config

import (
	"errors"
	"os"

	"github.com/go-postgresql-crud/logger"
	"github.com/spf13/viper"
)

type config struct {
	DbHost, DbName, DbUser, DbPassword string
	DbPort                             int
}

// ConfigKeys ...
var ConfigKeys *config

// InitConfig ...
func InitConfig() error {

	viper.SetDefault("HOST", "localhost")
	viper.SetDefault("PORT", 5432)
	viper.SetDefault("USER", "")
	viper.SetDefault("PASSWORD", "")
	viper.SetDefault("DBNAME", "")

	if os.Getenv("ENVIRONMENT") == "DEV" {

		viper.SetConfigName("config")
		viper.SetConfigType("toml")
		viper.AddConfigPath("./")
		// windows path
		//viper.AddConfigPath("../.")
		err := viper.ReadInConfig()

		if err != nil {
			logger.Error(err, "Error when reading the config file. Default values will be used. Config file:")
		}

		logger.Info("initConfig", "file", viper.ConfigFileUsed())

	} else {
		viper.AutomaticEnv()
	}

	// validate the provider
	username := viper.GetString("USER")
	password := viper.GetString("PASSWORD")
	dbName := viper.GetString("DBNAME")

	if username == "" {
		return errors.New("no provider username found in config")
	}

	if password == "" {
		return errors.New("no provider password found in config")
	}

	if dbName == "" {
		return errors.New("no provider dbName found in config")
	}

	ConfigKeys = &config{
		DbHost:     viper.GetString("HOST"),
		DbPort:     viper.GetInt("PORT"),
		DbUser:     username,
		DbPassword: password,
		DbName:     dbName,
	}

	return nil
}
