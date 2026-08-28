package handlers

import (
	"TaskFlow/common"
	"TaskFlow/store"
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
)

type CreateTaskRequest struct {
	Title  string `json:"title"`
	Status string `json:"status"`
}

func CreateTask(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var task store.Task
	var response common.Response
	var request CreateTaskRequest

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response = common.Response{Message: "Некорректные данные!"}
		json.NewEncoder(w).Encode(response)
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
	w.WriteHeader(http.StatusOK)
	response = common.Response{Message: "Задача успешно создана!"}

	json.NewEncoder(w).Encode(response)

}
