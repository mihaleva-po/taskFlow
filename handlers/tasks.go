package handlers

import (
	"TaskFlow/common"
	"TaskFlow/store"
	"encoding/json"
	"net/http"
)

func Tasks(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(store.Tasks)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := common.Response{Message: "Произошла ошибка!"}
		json.NewEncoder(w).Encode(response)
	} else {
		w.WriteHeader(http.StatusOK)
	}

}
