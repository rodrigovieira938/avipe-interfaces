package main

import (
	"fmt"

	"github.com/rodrigovieira938/avipe-interfaces/applications/example"
	"github.com/rodrigovieira938/avipe-interfaces/internal"
	"github.com/rodrigovieira938/avipe-interfaces/pkg/app"
)

func main() {
	apps := []app.Application{
		example.CreateExampleApp(),
	}
	for _, app := range apps {
		fmt.Printf("App Name: %s\n", app.Name())
	}
	app := internal.CreateFramework()
	app.Run()
}
