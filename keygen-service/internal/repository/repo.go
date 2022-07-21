package repository

import (
	"github.com/trite8q1/go-license-v2/keygen-service/internal/entity"
)

type Repository interface {
	GetAllLicensesV1() ([]entity.License, error)
	GetAllLicensesV2() ([]entity.License, error)
}
