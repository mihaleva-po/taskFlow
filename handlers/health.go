package handlers

import (
	"io"
	"net/http"
)

func HealthHandler(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, `{"status":"ok"}`)
}
