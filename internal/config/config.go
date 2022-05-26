package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Port       string
	DBUser     string
	DBPassword string
	DBName     string
}

func New() *Config {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	return &Config{
		Port:       viper.GetString("PORT"),
		DBUser:     viper.GetString("DBUSER"),
		DBPassword: viper.GetString("DBPASSWORD"),
		DBName:     viper.GetString("DBNAME"),
	}
}
