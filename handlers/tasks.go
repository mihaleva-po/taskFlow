package handlers

import (
	"TaskFlow/store"
	"TaskFlow/common"

	"encoding/json"
	"log"
	"net/http"
)


func (h *Handler) Tasks(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	err := h.store.GetTasks(r.Context())

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("get tasks: %v", err)
		response := common.Response{Message: "Не удалось получить список задач!"}
		json.NewEncoder(w).Encode(response)
		return 
	} 

		w.WriteHeader(http.StatusOK)
}
