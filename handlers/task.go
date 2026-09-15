package handlers

import (
	"TaskFlow/common"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (h *Handler) Task(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	id := r.PathValue("id")

	if _, err := uuid.Parse(id); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(common.Response{Message: "Некорректный UUID!"})
		return
	}

	task, err := h.store.GetTask(r.Context(), id)

	if err != nil {

		if errors.Is(err, pgx.ErrNoRows) {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(common.Response{Message: "Задача не найдена!"})
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("get task: %v", err)
		json.NewEncoder(w).Encode(common.Response{Message: "Произошла ошибка!"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}
