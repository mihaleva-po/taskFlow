package handlers

import (
	"TaskFlow/common"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/jackc/pgx/v5"
)

func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")

	err := h.store.DeleteTask(r.Context(), id)

	if err != nil {

		if errors.Is(err, pgx.ErrNoRows) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(common.Response{Message: "Такой задачи не существует!"})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(common.Response{Message: "Ошибка удаления!"})
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
