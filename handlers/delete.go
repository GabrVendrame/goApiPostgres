package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"postgresApi/models"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func Delete(writer http.ResponseWriter, req *http.Request) {
    id, err := strconv.Atoi(chi.URLParam(req, "id"))
    if err != nil {
        http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
        log.Printf("Error parsing id: %v", err)
        return
    }

    rows, err := models.Delete(int64(id))
    if err != nil {
        http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
        log.Printf("Error while removing: %v", err)
        return
    }

    if rows > 1 {
        log.Printf("Error: removed %d rows", rows)
    }

    res := map[string]any{
        "Error": false,
        "Message": "Sucessful removed",
    }
    
    writer.Header().Add("Content-Type", "application/json")
    json.NewEncoder(writer).Encode(res)
}
