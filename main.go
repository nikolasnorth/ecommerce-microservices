package main

import (
	"account-service/account"
	"account-service/db"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
)

const (
	configFilename = "config.json"
	port           = ":8000"
)

func seedDb() error {
	// Reset database
	err := db.Drop()
	if err != nil {
		return errors.New(fmt.Sprintf("could not drop table: %v", err))
	}

	// Create tables
	_, err = db.Conn.Exec(account.CreateTableSql)
	if err != nil {
		return errors.New(fmt.Sprintf("could not create tables: %v", err))
	}

	// Insert accounts
	a1 := account.New("nikolas@email.com", "nikolas", "Nikolas N")
	a2 := account.New("johndoe@email.com", "johndoe", "John Doe")
	a3 := account.New("janedoe@email.com", "janedoe", "Jane Doe")

	err = a1.Insert()
	if err != nil {
		return errors.New(fmt.Sprintf("could not insert account into db: %v", err))
	}

	err = a2.Insert()
	if err != nil {
		return errors.New(fmt.Sprintf("could not insert account into db: %v", err))
	}

	err = a3.Insert()
	if err != nil {
		return errors.New(fmt.Sprintf("could not insert account into db: %v", err))
	}

	return nil
}

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

	err = db.Init(config)
	if err != nil {
		log.Fatalf("could not open database connection: %v", err)
	}

	err = seedDb()
	if err != nil {
		log.Fatalf("could not seed database: %v", err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Account service handlers
	r.Get("/api/v1/accounts/{id}", account.GetByIdHandler)
	r.Get("/api/v1/accounts", account.GetHandler)
	r.Post("/api/v1/accounts", account.PostHandler)
	r.Put("/api/v1/accounts/{id}", account.UpdateByIdHandler)
	r.Delete("/api/v1/accounts/{id}", account.DeleteAccountByIdHandler)

	go func() {
		err = http.ListenAndServe(port, r)
		if err != nil {
			log.Fatalf("failed to start server: %v", err)
		}
	}()
	log.Println("listening on http://localhost" + port)

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	log.Println("received", sig)

	log.Println("dropping database...")
	err = db.Drop()
	if err != nil {
		log.Println(fmt.Errorf("could not drop database: %v", err).Error())
	}

	log.Println("server closed")
}
