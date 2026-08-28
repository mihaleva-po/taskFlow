package main

import (
	"TaskFlow/handlers"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/health", handlers.HealthHandler)

	mux.HandleFunc("POST /tasks", handlers.CreateTask)
	mux.HandleFunc("GET /tasks", handlers.Tasks)
	mux.HandleFunc("GET /tasks/{id}", handlers.Task)
	mux.HandleFunc("PUT /tasks/{id}", handlers.EditTask)
	mux.HandleFunc("DELETE /tasks/{id}", handlers.DeleteTask)

	http.ListenAndServe("localhost:8080", mux)

}
