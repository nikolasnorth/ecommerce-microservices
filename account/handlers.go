package account

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

// GetByIdHandler returns the account associated with given `id` URL param.
func GetByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	_, err := w.Write([]byte(id))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// GetByEmailHandler returns the account associated with the given `email` URL query param.
func GetByEmailHandler(w http.ResponseWriter, r *http.Request) {
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

func PostHandler(w http.ResponseWriter, r *http.Request) {
	var a Account
	err := a.FromJson(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	err = a.SendJson(w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
