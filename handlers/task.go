package handlers

import (
	"TaskFlow/common"
	"TaskFlow/store"
	"encoding/json"
	"net/http"
)

func Task(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	id := r.PathValue("id")

	if id == "" {
		response := common.Response{Message: "Некорректный ID!"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return

	}

	for _, v := range store.Tasks {
		if v.ID == id {
			json.NewEncoder(w).Encode(v)
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(common.Response{Message: "Задача не найдена!"})

}
