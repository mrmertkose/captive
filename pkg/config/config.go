package config

import (
	"captive/config"
	"github.com/gofiber/fiber/v2/log"
	"github.com/spf13/viper"
)

var configurations config.Config

func Get() config.Config {
	return configurations
}

func Set() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading the configs")
	}

	err := viper.Unmarshal(&configurations)
	if err != nil {
		log.Fatal("unable to decode into struct")
	}
}
