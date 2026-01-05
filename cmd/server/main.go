package main

import (
	"log"
	"net/http"
	"os"

	apphttp "github.com/jebesh67/golang-http/internal/http"

	"github.com/jebesh67/golang-http/internal/config"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("APP_ENV") == "local" {
		err := godotenv.Load(".env.local")
		if err != nil {
			log.Fatal("Error loading .env.local file")
		}
	}

	cfg := config.LoadConfig()

	mux := http.NewServeMux()
	apphttp.HandleRoutes(mux)

	log.Printf("Listening on port %s\n", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, mux))
}
