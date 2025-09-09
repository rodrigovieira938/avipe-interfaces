package handlers

import (
	"encoding/json"
	"net/http"

	v1 "github.com/rodrigovieira938/avipe-interfaces/internal/api/v1"
	"github.com/rodrigovieira938/avipe-interfaces/pkg/app"
)

type ApplicationsHandler struct {
	Apps     []app.Application
	response v1.ApplicationsResponse
}

func CreateApplicationsHandler(apps []app.Application) *ApplicationsHandler {
	return &ApplicationsHandler{
		Apps:     apps,
		response: v1.CreateApplicationsResponse(apps),
	}
}

func (h *ApplicationsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(h.response)
}
