package db

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
)

type Config struct {
	Host     string `json:"dbHost"`
	Port     int    `json:"dbPort"`
	User     string `json:"dbUser"`
	Password string `json:"dbPassword"`
	Name     string `json:"dbName"`
}

// Conn maintains a database connection.
var db *sql.DB

// InitDb opens a connection to the database specified by the given Config.
func InitDb(c Config) error {
	var err error

	db, err = sql.Open("postgres", fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s",
		c.Host, c.Port, c.User, c.Password, c.Name,
	))
	if err != nil {
		return err
	}

	return db.Ping()
}

func SeedDb() error {
	// Reset database
	err := DropDb()
	if err != nil {
		return errors.New(fmt.Sprintf("could not drop table: %v", err))
	}

	// Create tables
	_, err = db.Exec(`
CREATE TABLE IF NOT EXISTS accounts
(
    id         integer primary key generated always as identity,
    email      varchar,
    username   varchar,
    full_name  varchar,
    created_at timestamp default now()
);
`)
	if err != nil {
		return errors.New(fmt.Sprintf("could not create tables: %v", err))
	}

	// Insert rows
	_, err = db.Exec(`
INSERT INTO accounts (email, username, full_name)
VALUES ('nikolas@email.com', 'nikolas', 'Nikolas N')`)
	if err != nil {
		return errors.New(fmt.Sprintf("could not insert rows: %v", err))
	}

	return nil
}

func DropDb() error {
	// Reset database
	_, err := db.Exec(`
DROP SCHEMA public CASCADE;
CREATE SCHEMA public;

GRANT ALL ON SCHEMA public TO postgres;
GRANT ALL ON SCHEMA public to public;`)
	if err != nil {
		return errors.New(fmt.Sprintf("could not drop table: %v", err))
	}

	return nil
}
