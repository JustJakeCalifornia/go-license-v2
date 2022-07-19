package main

import (
	"log"
	"net/http"

	"github.com/trite8q1/go-keygen-service/internal/routes"
)

type Config struct {
	WebPort string
}

func main() {
	router := routes.NewRouter()

	config := Config{
		WebPort: ":8080",
	}

	log.Printf("Server started on port %s\n", config.WebPort)
	log.Fatal(http.ListenAndServe(config.WebPort, router))
}
