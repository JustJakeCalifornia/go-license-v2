package rest

import (
	"encoding/json"
	"net/http"

	"github.com/trite8q1/go-license-v2/keygen-service/core/service"
)

//Index route
func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	payload := map[string]string{
		"message": "test",
	}
	bytes, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}
	w.Write(bytes)
}

func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	payload := map[string]string{
		"status": "ok",
	}
	bytes, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}
	w.Write(bytes)
}

// exopse licenses to api
func GetLicenses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	for i := 0; i <= 100; i++ {
		service.GenerateAvailableLicenses()
	}
	payload := service.AvailableLicenses

	bytes, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}
	w.Write(bytes)
}
