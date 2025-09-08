package example

import (
	"github.com/gorilla/mux"
	"github.com/rodrigovieira938/avipe-interfaces/pkg/app"
)

type ExampleApp struct {
}

func (app *ExampleApp) Name() string {
	return "ExampleApp"
}
func (app *ExampleApp) InitializeRoutes(r *mux.Router) {
	// Initialize your application routes here
}

func CreateExampleApp() app.Application {
	return &ExampleApp{}
}
