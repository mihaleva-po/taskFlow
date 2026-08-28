package handlers

import (
	"TaskFlow/common"
	"TaskFlow/store"
	"encoding/json"
	"net/http"
)

type EditTaskRequest struct {
	Title  string `json:"title"`
	Status string `json:"status"`
}

func EditTask(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	id := r.PathValue("id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(common.Response{Message: "Не корректный идентификатор задачи!"})
		return
	}

	var task EditTaskRequest

	err := json.NewDecoder(r.Body).Decode(&task)

	if err != nil {
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

	found := false
	for i, v := range store.Tasks {
		if v.ID == id {
			store.Tasks[i] = store.Task{ID: id, Title: task.Title, Status: task.Status}
			found = true
			break
		}
	}

	if !found {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(common.Response{Message: "Такой задачи не существует!"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(common.Response{Message: "Задача изменена!"})

}
