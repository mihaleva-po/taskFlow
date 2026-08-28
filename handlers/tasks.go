package handlers

import (
	"TaskFlow/store"
	"encoding/json"
	"net/http"
)

func Tasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(store.Tasks)
}
