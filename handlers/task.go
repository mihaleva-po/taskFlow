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

	for _, v := range store.Tasks {
		if v.ID == id {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(v)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(common.Response{Message: "Задача не найдена!"})

}