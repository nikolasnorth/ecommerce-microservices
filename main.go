package main

import (
	"account-service/account"
	"account-service/db"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"os"
)

const (
	configFilename = "config.json"
	port           = ":8000"
)

func main() {
	configContents, err := os.ReadFile(configFilename)
	if err != nil {
		log.Fatalf("failed to open %s: %v", configFilename, err)
	}

	config := db.Config{}
	err = json.Unmarshal(configContents, &config)
	if err != nil {
		log.Fatalf("failed to unmarshal config file: %v", err)
	}

	err = db.InitDb(config)
	if err != nil {
		log.Fatalf("could not open database connection: %v", err)
	}

	dbSeed := flag.Bool("seed", false, "run the seed script to populate the database")
	if *dbSeed {
		// Seed database
		fmt.Println("seeding database...")
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Account service handlers
	r.Get("/api/v1/accounts/{id}", account.GetByIdHandler)
	r.Get("/api/v1/accounts", account.GetHandler)
	r.Post("/api/v1/accounts", account.PostHandler)
	r.Put("/api/v1/accounts/{id}", account.UpdateByIdHandler)
	r.Delete("/api/v1/accounts/{id}", account.DeleteAccountByIdHandler)

	log.Println("listening on http://localhost" + port)
	err = http.ListenAndServe(port, r)
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
