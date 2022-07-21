package config

import (
	"database/sql"

	"github.com/trite8q1/go-license-v2/keygen-service/internal/repository"
	"github.com/trite8q1/go-license-v2/keygen-service/internal/repository/postgres"
)

type Config struct {
	WebPort string
	Repo    repository.Repository
}

func NewConfig() *Config {
	return &Config{
		WebPort: ":8080",
	}
}

func (app *Config) SetupPostgresRepository(conn *sql.DB) {
	db := postgres.NewPostgresRepository(conn)
	app.Repo = db
}
