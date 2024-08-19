package config

type Config struct {
	postgresURL string `env:"POSTGRES_URL" default:"postgres://user:password@localhost:5432/db?sslmode=disable" validate:"uri"`
	apiEndPoint string `env:"API_ENDPOINT" default:"localhost:1323" validate:"uri"`
}
