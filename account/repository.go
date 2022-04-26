package account

import (
	"account-service/db"
	"encoding/json"
	"io"
	"time"
)

type Account struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	FullName  string    `json:"fullName"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

const CreateTableSql = `
CREATE TABLE IF NOT EXISTS accounts
(
    id         integer primary key generated always as identity,
    email      varchar,
    username   varchar,
    full_name  varchar,
    created_at timestamp default now()
);`

func (a *Account) FromJson(r io.Reader) error {
	return json.NewDecoder(r).Decode(a)
}

func New(email, username, fullName string) Account {
	return Account{
		Email:    email,
		Username: username,
		FullName: fullName,
	}
}

func (a *Account) Insert() error {
	const query = "INSERT INTO accounts (email, username, full_name) VALUES ($1, $2, $3) RETURNING id;"
	err := db.Conn.QueryRow(query, a.Email, a.Username, a.FullName).Scan(&a.ID)
	if err != nil {
		return err
	}
	return nil
}
