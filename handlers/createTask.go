package handlers

import (
	"TaskFlow/common"
	"TaskFlow/store"
	"encoding/json"

	"log"
	"net/http"

	"github.com/google/uuid"
)

type CreateTaskRequest struct {
	Title  string `json:"title"`
	Status string `json:"status"`
}

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var request CreateTaskRequest

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(common.Response{Message: "Некорректные данные!"})
		return
	}

	if request.Status == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(common.Response{Message: "Поле status обязательно для заполнения!"})
		return
	}

	if request.Title == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(common.Response{Message: "Поле title обязательно для заполнения!"})
		return
	}

	if !common.IsValidStatus(request.Status) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(common.Response{Message: "Недопустимый статус!"})
		return
	}

	task := store.Task{
		ID:     uuid.New().String(),
		Title:  request.Title,
		Status: request.Status,
	}

	err := h.store.CreateTask(r.Context(), task)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("create task: %v", err)
		json.NewEncoder(w).Encode(common.Response{Message: "Не удалось создать задачу!"})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(common.Response{Message: "Задача успешно создана!"})

}
