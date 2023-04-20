package config

import (
	"github.com/caarlos0/env/v7"
)

type Config struct {
	User     string `env:"USER" envDefault:"root"`
	Password string `env:"PASSWORD" envDefault:"root"`
	URL      string `env:"URL" envDefault:"localhost"`
	JWTKey   string `env:"JWT_KEY" envDefault:"supersecret"`
	Level    string `env:"APP_MODE" envDefault:"dev"`
	Port     string `env:"PORT" envDefault:":8080"`
	DB       struct {
		DSN          string `env:"dsn" envDefault:"postgres://onelab:postgre@localhost/onelab?sslmode=disable"`
		MaxOpenConns int    `env:"db-max-open-conns" envDefault:"25"`
		MaxIdleConns int    `env:"db-max-idle-conns" envDefault:"25"`
		MaxIdleTime  string `env:"db-max-idle-time" envDefault:"15m"`
	}
}

func NewConfig() *Config {
	cfg := Config{}
	env.Parse(&cfg)
	return &cfg
}
