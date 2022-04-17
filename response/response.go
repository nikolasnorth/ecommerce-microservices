package response

import (
	"encoding/json"
	"net/http"
)

// Json encodes `payload` as JSON and writes to `w`.
func Json(w http.ResponseWriter, status int, payload map[string]any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
