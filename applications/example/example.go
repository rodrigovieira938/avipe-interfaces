package example

import (
	"fmt"
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/gorilla/mux"
	"github.com/rodrigovieira938/avipe-interfaces/pkg/app"
)

type ExampleApp struct {
}

func (app *ExampleApp) Name() string {
	return "ExampleApp"
}
func (app *ExampleApp) InitializeRoutes(r *mux.Router) {
	subr := r.PathPrefix("/ExampleApp").Subrouter()
	fmt.Println("Initializing routes for ExampleApp")
	subr.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World from ExampleApp!")
	}).Methods("GET")
	subr.Handle("/interface", templ.Handler(MainPage())).Methods("GET")
	subr.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, time.Now().Format(time.TimeOnly))
	}).Methods("GET")
}

func CreateExampleApp() app.Application {
	return &ExampleApp{}
}
