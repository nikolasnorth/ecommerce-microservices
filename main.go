package main

import (
	"account-service/account"
	"account-service/db"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

type Config struct {
	DbHost     string `json:"dbHost"`
	DbPort     int    `json:"dbPort"`
	DbUser     string `json:"dbUser"`
	DbPassword string `json:"dbPassword"`
	DbName     string `json:"dbName"`
}

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

	db.Conn, err = sql.Open("postgres", fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s",
		config.Host, config.Port, config.User, config.Password, config.Name,
	))
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	err = db.Conn.Ping()
	if err != nil {
		log.Fatalf("failed to ping database: %v", err)
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
