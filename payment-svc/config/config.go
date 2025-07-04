package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/saleh-ghazimoradi/MicroEcoBay/payment_service/slg"
	"time"
)

var AppConfig *Config

type Config struct {
	ServerConfig ServerConfig
	Database     Database
}

type Database struct {
	DatabaseHost     string        `env:"DATABASE_HOST"`
	DatabasePort     string        `env:"DATABASE_PORT"`
	DatabaseUser     string        `env:"DATABASE_USER"`
	DatabasePassword string        `env:"DATABASE_PASSWORD"`
	DatabaseName     string        `env:"DATABASE_NAME"`
	DatabaseSSLMode  string        `env:"DATABASE_SSLMODE"`
	MaxOpenConn      int           `env:"DB_MAX_OPEN_CONNECTIONS"`
	MaxIdleConn      int           `env:"DB_MAX_IDLE_CONNECTIONS"`
	MaxLifetime      time.Duration `env:"DB_MAX_LIFETIME"`
	MaxIdleTime      time.Duration `env:"DB_MAX_IDLE_TIME"`
	Timeout          time.Duration `env:"DB_TIMEOUT"`
}

type ServerConfig struct {
	Port string `env:"SERVER_PORT"`
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
