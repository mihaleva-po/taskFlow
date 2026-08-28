package handlers

import (
	"TaskFlow/common"
	"TaskFlow/store"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type CreateTaskRequest struct {
	Title  string `json:"title"`
	Status string `json:"status"`
}

func CreateTask(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var task store.Task
	var request CreateTaskRequest

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
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

	task = store.Task{
		ID:     uuid.New().String(),
		Title:  request.Title,
		Status: request.Status,
	}
	store.Tasks = append(store.Tasks, task)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(common.Response{Message: "Задача успешно создана!"})

}
