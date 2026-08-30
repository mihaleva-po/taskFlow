package handlers

import (
	"TaskFlow/common"
	"TaskFlow/store"
	"encoding/json"
	"net/http"
)

func DeleteTask(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	id := r.PathValue("id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(common.Response{Message: "Некорректный идентификатор задачи!"})
		return
	}

	hasTask := false

	for i, v := range store.Tasks {
		if v.ID == id {
			hasTask = true
			store.Tasks = append(store.Tasks[:i], store.Tasks[i+1:]...)
			break
		}
	}

	if hasTask {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(common.Response{Message: "Такой задачи не существует!"})

}