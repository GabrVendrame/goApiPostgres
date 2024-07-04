package main

import (
	"fmt"
	"net/http"
	"postgresApi/configs"
	"postgresApi/handlers"

	"github.com/go-chi/chi/v5"
)

func main() {
    err := configs.Load()
    if err != nil {
        panic(err)
    }

    router := chi.NewRouter()

    router.Post("/createTodo", handlers.Create)
    router.Put("/updateTodo/{id}", handlers.Update)
    router.Delete("/deleteTodo/{id}", handlers.Delete)
    router.Get("/getAllTodos", handlers.GetAllTodos)
    router.Get("/getTodo/{id}", handlers.GetTodoById)

    http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), router)

    fmt.Printf("Server running on port %s", configs.GetServerPort())
    fmt.Println("Routes")
    fmt.Printf("POST /createTodo")
    fmt.Printf("PUT /updatedTodo")
    fmt.Printf("DELETE /deleteTodo")
    fmt.Printf("GET /getAllTodos")
    fmt.Printf("GET /getTodo")
}

