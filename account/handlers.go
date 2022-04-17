package account

import (
	"account-service/response"
	"github.com/go-chi/chi/v5"
	"net/http"
)

// GetByIdHandler returns the account associated with given URL path parameter of `id`.
func GetByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	_, err := w.Write([]byte(id))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// GetHandler returns the account associated with the given URL query parameter(s).
func GetHandler(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	if email == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err := w.Write([]byte(email))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// PostHandler creates and returns a new account.
func PostHandler(w http.ResponseWriter, r *http.Request) {
	var a Account
	err := a.FromJson(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response.Json(w, http.StatusCreated, map[string]any{
		"account": &a,
	})
}

// UpdateByIdHandler updates the account associated with the given URL path parameter of `id`.
func UpdateByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	_, err := w.Write([]byte(id))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// DeleteAccountByIdHandler deletes the account associated with the given URL path parameter of `id`.
func DeleteAccountByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	_, err := w.Write([]byte(id))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
