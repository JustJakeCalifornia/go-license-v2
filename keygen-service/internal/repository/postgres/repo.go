package postgres

import (
	"database/sql"
	"time"

	"github.com/trite8q1/go-license-v2/keygen-service/internal/entity"
)

type PostgresRepository struct {
	Conn *sql.DB
}

const dbTimeout = time.Second * 3

var db *sql.DB // refactor in config for better testing?

func NewPostgresRepository(pool *sql.DB) *PostgresRepository {
	db = pool
	return &PostgresRepository{
		Conn: pool,
	}
}

func (u *PostgresRepository) GetAllLicensesV1() ([]entity.License, error) {
	return nil, nil
}

func (u *PostgresRepository) GetAllLicensesV2() ([]entity.License, error) {
	return nil, nil
}
