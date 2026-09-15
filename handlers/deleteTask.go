package handlers

import (
	"TaskFlow/common"
	"encoding/json"
	"errors"
	"net/http"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	w.Header().Set("Content-Type", "application/json")

	if _, err := uuid.Parse(id); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(common.Response{Message: "Некорректный UUID!"})
		return
	}

	err := h.store.DeleteTask(r.Context(), id)

	if err != nil {

		if errors.Is(err, pgx.ErrNoRows) {

			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(common.Response{Message: "Такой задачи не существует!"})
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(common.Response{Message: "Ошибка удаления!"})
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
