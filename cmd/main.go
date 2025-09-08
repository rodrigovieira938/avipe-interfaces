package main

import (
	"fmt"

	"github.com/rodrigovieira938/avipe-interfaces/applications"
	"github.com/rodrigovieira938/avipe-interfaces/internal"
	"github.com/rodrigovieira938/avipe-interfaces/pkg/app"
)

func main() {
	apps := []app.Application{
		applications.CreateExampleApp(),
	}
	for _, app := range apps {
		fmt.Printf("App Name: %s\n", app.Name())
	}
	app := internal.CreateFramework(apps)
	app.Run()
}
