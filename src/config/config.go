package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int
	}
	Database struct {
		Host     string
		Port     int
		User     string
		Password string
		Name     string
		SSLMode  string
	}
}

var Data Config

func LoadConfig() {
	viper.SetConfigName("env")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error read config: %v", err)
	}

	if err := viper.Unmarshal(&Data); err != nil {
		log.Fatalf("Error parse config: %v", err)
	}

	log.Println("Config loaded")
}
