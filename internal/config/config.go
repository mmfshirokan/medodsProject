package config

type Config struct {
	postgresURL string `env:"POSTGRES_URL"`
	apiEndPoint string `env:"API_ENDPOINT"`
}
