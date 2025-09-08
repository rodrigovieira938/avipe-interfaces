package main

import (
	"github.com/rodrigovieira938/avipe-interfaces/internal"
)

func main() {
	app := internal.CreateFramework()
	app.Run()
}
