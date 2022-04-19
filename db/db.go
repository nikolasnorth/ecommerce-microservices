package db

import (
	"database/sql"
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

// InitDb opens a connection to the database specified by Config.
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
