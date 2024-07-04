package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"postgresApi/models"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func GetAllTodos(writer http.ResponseWriter, res *http.Request) {
	todos, err := models.GetAll()
	if err != nil {
		log.Printf("Error getting todos lists: %v", err)

	}

	writer.Header().Add("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(todos)
}

func GetTodoById(writer http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(req, "id"))
	if err != nil {
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Printf("Error parsing id: %v", err)
		return
	}

	todo, err := models.Get(int64(id))
	if err != nil {
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Printf("Error while getting todo: %v", err)
		return
	}

	writer.Header().Add("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(todo)
}
