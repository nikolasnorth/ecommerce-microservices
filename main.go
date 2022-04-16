package main

import (
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

	config := Config{}
	err = json.Unmarshal(configContents, &config)
	if err != nil {
		log.Fatalf("failed to unmarshal config file: %v", err)
	}

	db, err := sql.Open("postgres", fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s",
		config.DbHost, config.DbPort, config.DbUser, config.DbPassword, config.DbName,
	))
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello world"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	})

	log.Println("listening on http://localhost" + port)
	err = http.ListenAndServe(port, r)
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
