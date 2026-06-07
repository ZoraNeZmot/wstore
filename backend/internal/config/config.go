package config

import (
	"sync"

	"github.com/caarlos0/env/v11"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
)

type Config struct {
	DebugMode      bool           `env:"DEBUG_MODE" envDefault:"false"`
	LogConfig      LogConfig      `envPrefix:"LOG_"`
	DatabaseConfig DatabaseConfig `envPrefix:"DATABASE_"`
	ServerConfig   ServerConfig   `envPrefix:"SERVER_"`
}

type LogConfig struct {
	Path   string `env:"PATH" envDefault:"wstore.log"`
	Level  string `env:"LEVEL" envDefault:"info"`  // panic, fatal, error, warn, info, debug, trace
	Format string `env:"FORMAT" envDefault:"text"` // json, text
}

type DatabaseConfig struct {
	Host     string `env:"HOST" envDefault:"localhost"`
	Port     int    `env:"PORT" envDefault:"5432"`
	User     string `env:"USER" envDefault:"postgres"`
	Password string `env:"PASSWORD" envDefault:"postgres"`
	Database string `env:"DATABASE" envDefault:"wstore"`
}

type ServerConfig struct {
	Host string `env:"HOST" envDefault:"0.0.0.0"`
	Port string `env:"PORT" envDefault:"8080"`
}

var config *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		config = &Config{}
		if err := env.Parse(config); err != nil {
			logrus.Fatalf("Failed to parse config: %v", err)
		}
	})
	return config
}
