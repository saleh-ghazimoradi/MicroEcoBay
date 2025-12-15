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
	Postgresql  Postgresql
	JWT         JWT
}

type Server struct {
	Host         string        `env:"SERVER_HOST"`
	Port         string        `env:"SERVER_PORT"`
	IdleTimeout  time.Duration `env:"SERVER_IDLE_TIMEOUT"`
	ReadTimeout  time.Duration `env:"SERVER_READ_TIMEOUT"`
	WriteTimeout time.Duration `env:"SERVER_WRITE_TIMEOUT"`
}

type KafkaConfig struct {
	Broker []string `env:"KAFKA_BROKER"`
	Topic  string   `env:"KAFKA_TOPIC"`
}

type Postgresql struct {
	Host        string        `env:"POSTGRES_HOST"`
	Port        string        `env:"POSTGRES_PORT"`
	User        string        `env:"POSTGRES_USER"`
	Password    string        `env:"POSTGRES_PASSWORD"`
	Name        string        `env:"POSTGRES_NAME"`
	MaxOpenConn int           `env:"POSTGRES_MAX_OPEN_CONN"`
	MaxIdleConn int           `env:"POSTGRES_MAX_IDLE_CONN"`
	MaxIdleTime time.Duration `env:"POSTGRES_MAX_IDLE_TIME"`
	SSLMode     string        `env:"POSTGRES_SSL_MODE"`
	Timeout     time.Duration `env:"POSTGRES_TIMEOUT"`
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
