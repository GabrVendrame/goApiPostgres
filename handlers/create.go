package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"postgresApi/models"
)

func Create(writer http.ResponseWriter, req *http.Request) {
	var todo models.Todo

	err := json.NewDecoder(req.Body).Decode(&todo)
	if err != nil {
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Printf("Error decoding json: %v", err)
		return
	}

	id, err := models.Insert(todo)
	var res map[string]any
	if err != nil {
		res = map[string]any{
			"Error":   true,
			"Message": fmt.Sprintf("Error while inserting: %v", err),
		}
	} else {
		res = map[string]any{
			"Error":   false,
			"Message": fmt.Sprintf("Insertion OK. ID: %d", id),
		}
	}

    writer.Header().Add("Content-Type", "application/json")
    json.NewEncoder(writer).Encode(res)
}
