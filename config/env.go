package config

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	GinMode     string `envconfig:"GIN_MODE"`
	HttpPort    string `envconfig:"HTTP_PORT"`
	SocketPort  string `envconfig:"SOCKET_PORT"`
	PostgresUrl string `envconfig:"POSTGRES_URL"`
	Database    string `envconfig:"DATABASE"`
	JwtSecret   string `envconfig:"JWT_SECRET"`
}

var ENV Config

func GetConfig() (*Config, error) {
	err := envconfig.Process("", &ENV)
	if err != nil {
		return nil, err
	}
	return &ENV, nil
}
