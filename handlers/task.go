package handlers

import (
	"TaskFlow/common"
	"encoding/json"
	"log"
	"net/http"
)

func (h *Handler) Task(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	id := r.PathValue("id")

	task, err := h.store.GetTask(r.Context(), id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("get task: %v", err)
		json.NewEncoder(w).Encode(common.Response{Message: "Произошла ошибка!"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}
