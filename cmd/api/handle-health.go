package main

import (
	"net/http"
)

func (api *Config) handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	payload := JsonResponse{
		Error:   false,
		Message: "The server is up and running!",
	}

	api.writeJSON(w, http.StatusOK, payload)
}
