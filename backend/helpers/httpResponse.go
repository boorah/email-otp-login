package helpers

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, requestError RequestError) {
	RespondWithJSON(w, requestError.GetStatusCode(), requestError)
}

func RespondWithJSON(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	stringifiedResponse, err := json.Marshal(data)
	if err != nil {
		log.Println("error while stringifying response:", err)

		w.Write([]byte(`{"type": "internal_server_error", "message": "An internal server error occurred"}`))

		return
	}

	w.Write(stringifiedResponse)
}
