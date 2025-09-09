package example

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rodrigovieira938/avipe-interfaces/pkg/app"
)

type ExampleApp struct {
}

func (app *ExampleApp) Name() string {
	return "ExampleApp"
}
func (app *ExampleApp) InitializeRoutes(r *mux.Router) {
	fmt.Println("Initializing routes for ExampleApp")
	r.HandleFunc("/exampleapp", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World from ExampleApp!")
	}).Methods("GET")
}

func CreateExampleApp() app.Application {
	return &ExampleApp{}
}
