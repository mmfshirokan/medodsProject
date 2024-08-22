package config

import "github.com/caarlos0/env/v11"

type Config struct {
	PostgresURL string `env:"POSTGRES_URL" envDefault:"postgres://user:password@localhost:5432/db?sslmode=disable" validate:"uri"`
	ApiEndPoint string `env:"API_ENDPOINT" envDefault:"localhost:1323" validate:"uri"`
}

func New() (Config, error) {
	var cnf Config
	err := env.Parse(&cnf)
	if err != nil {
		return Config{}, err
	}

	return cnf, nil
}
