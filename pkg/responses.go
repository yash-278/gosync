package responses

import (
	"encoding/json"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	dat, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(dat)
}

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	type customErrors struct {
		Error string `json:"error"`
	}

	respBody := customErrors{
		Error: msg,
	}

	RespondWithJSON(w, code, respBody)
}
