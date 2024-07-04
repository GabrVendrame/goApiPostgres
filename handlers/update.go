package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"postgresApi/models"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func Update(writer http.ResponseWriter, req *http.Request) {
    id, err := strconv.Atoi(chi.URLParam(req, "id"))
    if err != nil {
        http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
        log.Printf("Error parsing id: %v", err)
        return
    }
    
    var todo models.Todo

    err = json.NewDecoder(req.Body).Decode(&todo)
    if err != nil {
        http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
        log.Printf("Error decoding json: %v", err)
        return
    }

    rows, err := models.Update(int64(id), todo)
    if err != nil {
        http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
        log.Printf("Error while updating: %v", err)
        return
    }

    if rows > 1 {
        log.Printf("Error: updated %d rows", rows)
    }

    res := map[string]any{
        "Error": false,
        "Message": "Sucessful updated",
    }
    
    writer.Header().Add("Content-Type", "application/json")
    json.NewEncoder(writer).Encode(res)
}
