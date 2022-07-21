package service

import (
	"time"

	"github.com/trite8q1/go-license-v2/keygen-service/internal/entities"
	"github.com/trite8q1/go-license-v2/keygen-service/utils"
)

// TODO: bad practise to have this variables in global scope -> maybe struct?
// LATER?: storing in database

var AvailableLicenses []entities.License // TODO: change to private if database is configured: availableLicenses
var UsedLicenses []entities.License      // TODO: change to private if database is configured: usedLicenses

// Mocking License data
func GenerateAvailableLicenses() {
	license := entities.License{
		Key:         generateRandomLicenseKey(),
		IsUsed:      entities.AVAILABLE,
		CreatedAt:   time.Now().Format("2006-01-02"),
		ExpiresAt:   time.Now().AddDate(0, 0, 1).Format("2006-01-02"),
		CompanyId:   "1",
		CompanyName: "Company 1",
		TeamId:      "1",
		TeamName:    "Team 1",
	}
	AvailableLicenses = append(AvailableLicenses, license)
	utils.EntityLogger(AvailableLicenses)
}
