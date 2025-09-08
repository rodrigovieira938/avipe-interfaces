package internal

import (
	"log/slog"

	"github.com/gorilla/mux"

	handlers "github.com/rodrigovieira938/avipe-interfaces/internal/api/v1/handlers"
)

type ApiRouter struct {
	r *mux.Router
}

func CreateApiRouter(framework *Framework) ApiRouter {
	r := ApiRouter{r: framework.main_router.PathPrefix("/api").Subrouter()}

	framework.main_router.Handle("/version", &handlers.VersionHandler{}).Methods("GET")
	v1_router := r.r.PathPrefix("/v1").Subrouter()
	v1_router.Handle("/apps", handlers.CreateApplicationsHandler(framework.GetApps())).Methods("GET")

	for _, application := range framework.apps {
		slog.Info("Registering application", "app", application.Name())
		app_router := v1_router.PathPrefix("/" + application.Name()).Subrouter()
		application.InitializeRoutes(app_router)
	}
	return r
}
