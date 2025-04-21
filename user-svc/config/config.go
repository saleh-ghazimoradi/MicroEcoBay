package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/slg"
	"time"
)

var AppConfig *Config

type Config struct {
	Server      Server
	KafkaConfig KafkaConfig
	Database    Database
}

type Server struct {
	BodyLimit    int           `env:"BODY_LIMIT"`    // 1024 * 1024
	WriteTimeout time.Duration `env:"WRITE_TIMEOUT"` // 10s
	ReadTimeout  time.Duration `env:"READ_TIMEOUT"`  // 5s
	IdleTimeout  time.Duration `env:"IDLE_TIMEOUT"`  // 30s
	RateLimit    int           `env:"RATE_LIMIT"`    // 100
	RateLimitExp time.Duration `env:"RATE_EXP"`      // 60s
	Port         string        `env:"PORT"`          // 3000
	Timeout      time.Duration `env:"TIMEOUT"`       // 30s
}

type KafkaConfig struct {
	Broker string `env:"KAFKA_BROKER"`
	Topic  string `env:"KAFKA_TOPIC"`
}

type Database struct {
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
