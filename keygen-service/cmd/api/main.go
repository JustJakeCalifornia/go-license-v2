package main

import (
	"log"
	"net/http"

	"github.com/trite8q1/go-license-v2/keygen-service/internal/config"
	"github.com/trite8q1/go-license-v2/keygen-service/internal/repository/postgres"
	"github.com/trite8q1/go-license-v2/keygen-service/internal/rest/routes"
)

func main() {
	router := routes.NewRouter()
	config := config.NewConfig()
	postgres.ConnectToDB()

	log.Printf("Server started on port %s\n", config.WebPort)
	log.Fatal(http.ListenAndServe(config.WebPort, router))
}
