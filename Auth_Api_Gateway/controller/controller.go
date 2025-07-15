package controller

import (
	"encoding/json"
	"net/http"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Set header
	w.WriteHeader(500)
	response := map[string]string{"message": "pong"} // Or use a struct
	json.NewEncoder(w).Encode(response)              // Encode to JSON
}
