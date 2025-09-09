package main

import (
	"github.com/rodrigovieira938/avipe-interfaces/applications"
	"github.com/rodrigovieira938/avipe-interfaces/internal"
	"github.com/rodrigovieira938/avipe-interfaces/pkg/app"
)

func main() {
	apps := []app.Application{
		applications.CreateExampleApp(),
	}
	app := internal.CreateFramework(apps)
	app.Run()
}
