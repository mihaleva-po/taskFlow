package handlers

import (
	"TaskFlow/common"
	"TaskFlow/store"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type UpdateTaskRequest struct {
	Title  string `json:"title"`
	Status string `json:"status"`
}

func (h *Handler) UpdateTask(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	id := r.PathValue("id")

	if _, err := uuid.Parse(id); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(common.Response{Message: "Некорректный UUID!"})
		return 
	}

	var task UpdateTaskRequest

	if err := decoder.Decode(&task); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(common.Response{Message: "Некорректные данные!"})
		return
	}

	if task.Status == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(common.Response{Message: "Поле status обязательно для заполнения!"})
		return
	}

	if task.Title == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(common.Response{Message: "Поле title обязательно для заполнения!"})
		return
	}

	if !common.IsValidStatus(task.Status) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(common.Response{Message: "Недопустимый статус!"})
		return
	}

	err := h.store.UpdateTask(r.Context(), store.Task{ID: id, Title: task.Title, Status: task.Status})

	if err != nil {

		if errors.Is(err, pgx.ErrNoRows) {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(common.Response{Message: "Такой задачи не существует!"})
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("update task: %v", err)
		json.NewEncoder(w).Encode(common.Response{Message: "Не удалось изменить задачу!"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(common.Response{Message: "Задача изменена!"})

}
