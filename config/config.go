package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ServerAddress string
	PostgresURL   string
	KafkaBroker   string
	KafkaTopic    string
}

func LoadConfig() *Config {
	viper.SetConfigName("config")

	viper.AddConfigPath("config/")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Unable to decode into struct %v", err)
	}

	return &config
}
