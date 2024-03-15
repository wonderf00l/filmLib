package configs

import "os"

type PostgresConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

func NewPostgresConfig() PostgresConfig {
	return PostgresConfig{
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Database: os.Getenv("POSTGRES_DB"),
	}
}
