package config

type Config struct {
	postgresURL string `env:"POSTGRES_URL"`
	echoPort    string `env:"ECHO_PORT"`
}
