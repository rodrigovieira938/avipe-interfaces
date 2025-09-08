package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/rodrigovieira938/avipe-interfaces/internal/api/utils"
	v1 "github.com/rodrigovieira938/avipe-interfaces/internal/api/v1"
)

type VersionHandler struct {
	response v1.VersionResponse
}

func CreateVersionHandler(version string, supportedVersions []string) *VersionHandler {
	return &VersionHandler{
		response: v1.VersionResponse{
			Version:           version,
			SupportedVersions: supportedVersions,
		},
	}
}

func (vh *VersionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := v1.VersionResponse{
		Version: "1.0",
		//V1.0 can only support itself
		SupportedVersions: []string{"1.0"},
	}
	_, error := json.Marshal(response)
	if error != nil {
		slog.Error("Failed to marshal version response", "error", error.Error())
		utils.CreateInternalStatusError().Write(w)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
