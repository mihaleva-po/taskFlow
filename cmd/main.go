package main

import (
	"TaskFlow/handlers"
	"TaskFlow/store"
	"context"
	"fmt"
	"net/http"
	"os"
)

func main() {

	// БД
	ctx := context.Background() // создаем пустой корневой контекст
	databaseURL := os.Getenv("DATABASE_URL")

	pool, err := store.NewPostgresPool(ctx, databaseURL)

	if err != nil {
		panic(err)
	}

	defer pool.Close()

	fmt.Println("Подсоединились к бд!")

	tasksStore := store.NewStore(pool)
	handler := handlers.NewHandler(tasksStore)

	// Методы
	mux := http.NewServeMux()

	mux.HandleFunc("/health", handlers.HealthHandler)

	mux.HandleFunc("POST /tasks", handler.CreateTask)
	mux.HandleFunc("GET /tasks", handler.Tasks)
	mux.HandleFunc("GET /tasks/{id}", handler.Task)
	mux.HandleFunc("PUT /tasks/{id}", handler.UpdateTask)
	mux.HandleFunc("DELETE /tasks/{id}", handler.DeleteTask)

	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		panic(err)
	}

}
