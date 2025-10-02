package internal

import (
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gorilla/mux"
	"github.com/rodrigovieira938/avipe-interfaces/internal/mainpage"
	"github.com/rodrigovieira938/avipe-interfaces/pkg/app"
)

type Framework struct {
	main_router *mux.Router
	api_router  ApiRouter
	apps        []app.Application
	mainpage    templ.Component
}

func (framework *Framework) initialize_router() {
	framework.mainpage = mainpage.MainPage(framework.apps)
	framework.main_router.Handle("/", templ.Handler(framework.mainpage))

	framework.main_router.PathPrefix("/static/").
        Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	framework.api_router = CreateApiRouter(framework)
}

func (framework *Framework) Initialize() {
	framework.initialize_router()
}

func (framework *Framework) Run() {
	slog.Info("Server is running on http://127.0.0.1:8080")
	err := http.ListenAndServe(":8080", framework.main_router)
	if err != nil {
		slog.Error("Error starting server: ", "err,", err.Error())
	}
}

func (framework *Framework) GetApps() []app.Application {
	return framework.apps
}

func CreateFramework(apps []app.Application) Framework {
	main_router := mux.NewRouter()
	framework := Framework{
		//Only initialize required fields
		main_router: main_router,
		apps:        apps,
	}
	framework.Initialize()
	return framework
}
