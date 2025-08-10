package config

import (
	"github.com/spf13/viper"
	"log"
)

var AppConfig Config

type Config struct {
	ServerPort     string `mapstructure:"PORT"`
	DBHost         string `mapstructure:"POSTGRES_HOST"`
	DBUserName     string `mapstructure:"POSTGRES_USER"`
	DBUserPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBName         string `mapstructure:"POSTGRES_DB"`
	DBPort         string `mapstructure:"POSTGRES_PORT"`
}

func LoadConfig() (err error) {
	viper.SetConfigFile("app.env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Println("No app.env file found or failed to load.")
		return err
	}

	err = viper.Unmarshal(&AppConfig)
	return err
}
