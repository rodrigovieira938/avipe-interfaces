package internal

import (
	"log/slog"

	"github.com/gorilla/mux"

	handlers "github.com/rodrigovieira938/avipe-interfaces/internal/api/v1/handlers"
)

type ApiRouter struct {
	r *mux.Router
}

func (r ApiRouter) initV1Routes(framework *Framework) {

	v1_router := r.r.PathPrefix("/v1").Subrouter()
	v1_router.Handle("/apps", handlers.CreateApplicationsHandler(framework.GetApps())).Methods("GET")

	for _, application := range framework.apps {
		slog.Info("Registering application", "app", application.Name())
		app_router := v1_router.PathPrefix("/" + application.Name()).Subrouter()
		application.InitializeRoutes(app_router)
	}
}

func CreateApiRouter(framework *Framework) ApiRouter {
	r := ApiRouter{r: framework.main_router.PathPrefix("/api").Subrouter()}

	r.r.Handle("/version", handlers.CreateVersionHandler("1.0", []string{"1.0"})).Methods("GET")
	r.initV1Routes(framework)

	return r
}
