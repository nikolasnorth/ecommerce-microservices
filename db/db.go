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

// Conn holds a database connection.
var Conn *sql.DB

// Init opens a connection to the database specified by the given Config.
func Init(c Config) error {
	var err error

	Conn, err = sql.Open("postgres", fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s",
		c.Host, c.Port, c.User, c.Password, c.Name,
	))
	if err != nil {
		return err
	}

	return Conn.Ping()
}

// Drop resets the database.
func Drop() error {
	_, err := Conn.Exec(`
DROP SCHEMA public CASCADE;
CREATE SCHEMA public;
GRANT ALL ON SCHEMA public TO postgres;
GRANT ALL ON SCHEMA public to public;`)
	if err != nil {
		return errors.New(fmt.Sprintf("could not drop table: %v", err))
	}

	return nil
}
