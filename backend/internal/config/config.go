package config

import (
	"crypto/tls"
	"database/sql"
	"errors"
	"fmt"
	"sync"

	"github.com/caarlos0/env/v11"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Config struct {
	DebugMode         bool              `env:"DEBUG_MODE" envDefault:"false"`
	LogConfig         LogConfig         `envPrefix:"LOG_"`
	ServerConfig      ServerConfig      `envPrefix:"SERVER_"`
	DatabaseConfig    DatabaseConfig    `envPrefix:"DATABASE_"`
	GoogleOAuthConfig GoogleOAuthConfig `envPrefix:"GOOGLE_"`
}

// GoogleOAuthConfig - Google OAuth configuration
//
// Values are loaded from environment variables
// RedirectURI is expanded from SERVER_METHOD, SERVER_HOST, SERVER_PORT
// This variables should be set in the .env file BEFORE the GoogleOAuthConfig is parsed
type GoogleOAuthConfig struct {
	ClientID     string `env:"CLIENT_ID" envDefault:""`
	ClientSecret string `env:"CLIENT_SECRET" envDefault:""`

	RedirectURI string   `env:"REDIRECT_URI,expand" envDefault:"${SERVER_METHOD}://${SERVER_HOST}:${SERVER_PORT}/api/auth/google/callback"`
	Scopes      []string `env:"SCOPES" envDefault:"https://www.googleapis.com/auth/userinfo.email"`
	config      *oauth2.Config
}

type LogConfig struct {
	Path   string `env:"PATH" envDefault:"wstore.log"`
	Level  string `env:"LEVEL" envDefault:"info"`  // panic, fatal, error, warn, info, debug, trace
	Format string `env:"FORMAT" envDefault:"text"` // json, text
}

type DatabaseConfig struct {
	Host     string `env:"HOST" envDefault:"localhost"`
	Port     int    `env:"PORT" envDefault:"5432"`
	User     string `env:"USER,required" envDefault:"postgres"`
	Password string `env:"PASSWORD,required" envDefault:"postgres"`
	Database string `env:"DATABASE,required" envDefault:"wstore"`
	db       *sql.DB
}

func (c *DatabaseConfig) GetDB() *sql.DB {
	return c.db
}

func (c *DatabaseConfig) openDB() error {
	if c.db == nil {
		db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", c.Host, c.Port, c.User, c.Password, c.Database))
		if err != nil {
			return err
		}
		c.db = db
		return nil
	}
	return errors.New("database already opened")
}

type ServerConfig struct {
	Host      string `env:"HOST" envDefault:"0.0.0.0"`
	Port      string `env:"PORT" envDefault:"8080"`
	Method    string `env:"METHOD" envDefault:"http"`
	JwtSecret string `env:"JWT_SECRET,required"`

	TLS       bool   `env:"TLS" envDefault:"false"`
	CertFile  string `env:"CERT_FILE" envDefault:"certificates/server.crt"`
	KeyFile   string `env:"KEY_FILE" envDefault:"certificates/server.key"`
	TLSConfig *tls.Config
}

var config *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		config = &Config{}

		if err := env.Parse(config); err != nil {
			logrus.Fatalf("Failed to parse config: %v", err)
		}

		if err := config.DatabaseConfig.openDB(); err != nil {
			logrus.Fatalf("Failed to open database: %v", err)
		}

		if config.ServerConfig.TLS {
			cert, err := tls.LoadX509KeyPair(config.ServerConfig.CertFile, config.ServerConfig.KeyFile)
			if err != nil {
				logrus.Fatalf("Failed to load TLS certificate: %v", err)
			}
			config.ServerConfig.TLSConfig = &tls.Config{
				Certificates: []tls.Certificate{
					cert,
				},
			}
		} else {
			config.ServerConfig.TLSConfig = nil
		}
	})
	return config
}

var googleOnce sync.Once

func (c *Config) GetGoogleOAuthConfig() *oauth2.Config {
	googleOnce.Do(func() {
		c.GoogleOAuthConfig.config = &oauth2.Config{
			ClientID:     c.GoogleOAuthConfig.ClientID,
			ClientSecret: c.GoogleOAuthConfig.ClientSecret,
			RedirectURL:  c.GoogleOAuthConfig.RedirectURI,
			Scopes:       c.GoogleOAuthConfig.Scopes,
			Endpoint:     google.Endpoint,
		}
	})
	return c.GoogleOAuthConfig.config
}
