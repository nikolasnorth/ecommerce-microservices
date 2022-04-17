package account

import (
	"encoding/json"
	"io"
	"time"
)

type Account struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (a *Account) FromJson(r io.Reader) error {
	return json.NewDecoder(r).Decode(a)
}
