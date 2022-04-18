package db

import (
	"database/sql"
)

type Config struct {
	Host     string `json:"dbHost"`
	Port     int    `json:"dbPort"`
	User     string `json:"dbUser"`
	Password string `json:"dbPassword"`
	Name     string `json:"dbName"`
}

// Conn maintains a database connection.
var Conn *sql.DB
