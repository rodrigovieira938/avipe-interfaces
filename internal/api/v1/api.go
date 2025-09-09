package v1

import "github.com/rodrigovieira938/avipe-interfaces/pkg/app"

type VersionResponse struct {
	Version           string   `json:"version"`
	SupportedVersions []string `json:"supported_versions"`
}
type ApplicationsResponse struct {
	Apps []string `json:"apps"`
}

func CreateApplicationsResponse(apps []app.Application) ApplicationsResponse {
	var appNames []string
	for _, app := range apps {
		appNames = append(appNames, app.Name())
	}
	return ApplicationsResponse{Apps: appNames}
}
