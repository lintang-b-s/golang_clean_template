package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		App       `yaml:"app"`
		HTTP      `yaml:"http"`
		Postgres  `yaml:"postgres"`
		LogConfig `yaml:"logger"`
		RabbitMQ  `yaml:"rabbitmq"`
		GRPC      `yaml:"grpc"`
	}

	// App -.
	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	// HTTP -.
	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	Postgres struct {
		Username string `env-required:"true" yaml:"username"`
		Password string `env-required:"true" yaml:"password"`
	}

	LogConfig struct {
		Level string `json:"level" yaml:"level"`
		// Filename   string `json:"filename" yaml:"filename"`
		// MaxSize    int    `json:"maxsize" yaml:"maxsize"`
		MaxAge     int `json:"max_age" yaml:"max_age"`
		MaxBackups int `json:"max_backups" yaml:"max_backups"`
	}

	RabbitMQ struct {
		RMQAddress string `json:"rabbitmqAddress" yaml:"rmqAddress" env:"RABBITMQ_ADDRESS"`
	}

	GRPC struct {
		URLGrpc string `json:"urlGRPC" yaml:"urlGRPC" env:"URL_GRPC"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	// err := cleanenv.ReadConfig(",./config/config.yml", cfg)
	err := cleanenv.ReadConfig("./config/config.yml", cfg)

	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
