package internal

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
)

type ApiRouter struct {
	r *mux.Router
}

func CreateApiRouter(framework *Framework) ApiRouter {
	r := ApiRouter{r: framework.main_router.PathPrefix("/api").Subrouter()}

	r.r.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "API Version 1.0")
	}).Methods("GET")
	for _, application := range framework.apps {
		slog.Info("Registering application", "app", application.Name())
		// Change to use a subrouter for each application
		r.r.HandleFunc("/"+application.Name(), func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello from %s!", application.Name())
		}).Methods("GET")
	}
	return r
}
