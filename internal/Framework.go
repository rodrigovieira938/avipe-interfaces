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
	api_router  ApiRouter
	apps        []app.Application
}

func (framework *Framework) initialize_router() {
	framework.main_router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!")
	}).Methods("GET")

	framework.api_router = CreateApiRouter(framework)
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
		//Only initialize required fields
		main_router: main_router,
		apps:        apps,
	}
	framework.Initialize()
	return framework
}
