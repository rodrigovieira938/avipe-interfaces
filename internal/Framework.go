package internal

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rodrigovieira938/avipe-interfaces/pkg/app"
)

type Framework struct {
	main_router mux.Router
	api_router  *mux.Router
	apps        []app.Application
}

func (framework *Framework) initialize_router() {
	framework.main_router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!")
	}).Methods("GET")
	framework.api_router.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "API Version 1.0")
	}).Methods("GET")
	for _, application := range framework.apps {
		slog.Info("Registering application", "app", application.Name())
		// Change to use a subrouter for each application
		framework.api_router.HandleFunc("/"+application.Name(), func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello from %s!", application.Name())
		}).Methods("GET")
	}
}

func (framework *Framework) Initialize() {
	framework.initialize_router()
}

func (framework *Framework) Run() {
	slog.Info("Server is running on http://127.0.0.1:8080")
	http.ListenAndServe(":8080", &framework.main_router)
}

func CreateFramework(apps []app.Application) Framework {
	main_router := mux.Router{}
	framework := Framework{
		main_router: main_router,
		api_router:  main_router.PathPrefix("/api").Subrouter(),
		apps:        apps,
	}
	framework.Initialize()
	return framework
}
