package main

import (
	"TaskFlow/handlers"
	"TaskFlow/middleware"
	"TaskFlow/store"
	"context"
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("Файл .env не найден, используются переменные окружения")
	}

	databaseURL := os.Getenv("DATABASE_URL")

	ctx := context.Background() // создаем пустой корневой контекст

	pool, err := store.NewPostgresPool(ctx, databaseURL)

	if err != nil {
		panic(err)
	}

	defer pool.Close()

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

	httpHandler := middleware.Logging(mux)

	if err := http.ListenAndServe("localhost:8080", httpHandler); err != nil {
		panic(err)
	}
}
