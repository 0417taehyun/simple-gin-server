package config

import (
	"github.com/spf13/viper"
)

var conf = Config{}

type Config struct {
	Application Application
	Database    Database
}

type Application struct {
	Address string
}

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func Init() {
	viper.AutomaticEnv()

	viper.SetDefault("APPLICATION_ADDRESS", "0.0.0.0:8000")

	conf = Config{
		Application: Application{
			Address: viper.GetString("APPLICATION_ADDRESS"),
		},
		Database: Database{
			Host:     viper.GetString("DATABASE_HOST"),
			Port:     viper.GetString("DATABASE_PORT"),
			User:     viper.GetString("DATABASE_USER"),
			Password: viper.GetString("DATABASE_PASSWORD"),
			Name:     viper.GetString("DATABASE_NAME"),
		},
	}
}

func GetConfig() Config {
	return conf
}
