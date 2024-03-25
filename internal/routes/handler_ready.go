package routes

import (
	"net/http"

	responses "github.com/yash-278/gosync/pkg"
)

func (cfg *apiConfig) readinessSuccessHandler(w http.ResponseWriter, r *http.Request) {
	type status struct {
		Status string `json:"status"`
	}
	responses.
		RespondWithJSON(w, 200, status{
			Status: "ok",
		})
}

func (cfg *apiConfig) readinessErrHandler(w http.ResponseWriter, r *http.Request) {
	responses.RespondWithError(w, 500, "Internal Server Error")
}
