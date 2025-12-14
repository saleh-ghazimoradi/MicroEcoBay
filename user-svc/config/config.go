package config

import (
	"github.com/caarlos0/env/v11"
	"sync"
	"time"
)

var (
	instance *Config
	once     sync.Once
	initErr  error
)

type Config struct {
	Server      Server
	KafkaConfig KafkaConfig
	Database    Database
	JWT         JWT
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
	Broker []string `env:"KAFKA_BROKER"`
	Topic  string   `env:"KAFKA_TOPIC"`
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

type JWT struct {
	Secret string        `env:"JWT_SECRET"`
	Exp    time.Duration `env:"JWT_EXP"`
}

func GetConfig() (*Config, error) {
	once.Do(func() {
		instance = &Config{}
		initErr = env.Parse(instance)
		if initErr != nil {
			instance = nil
		}
	})
	return instance, initErr
}
