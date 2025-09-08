package example

import "github.com/rodrigovieira938/avipe-interfaces/pkg/app"

type ExampleApp struct {
}

func (app *ExampleApp) Name() string {
	return "ExampleApp"
}

func CreateExampleApp() app.Application {
	return &ExampleApp{}
}
