package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/slg"
)

var AppConfig *Config

type Config struct {
	ServerConfig   ServerConfig
	KafkaConfig    KafkaConfig
	DatabaseConfig DatabaseConfig
}

type ServerConfig struct {
	Port string `env:"SERVER_PORT"`
}

type KafkaConfig struct {
	Broker string `env:"KAFKA_BROKER"`
	Topic  string `env:"KAFKA_TOPIC"`
}

type DatabaseConfig struct {
}

func LoadConfig() error {
	config := &Config{}

	if err := env.Parse(config); err != nil {
		slg.Logger.Error("error loading config", "error", err)
		return err
	}
	AppConfig = config

	return nil
}
