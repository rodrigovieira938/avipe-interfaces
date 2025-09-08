package app

import "github.com/gorilla/mux"

type Application interface {
	Name() string
	InitializeRoutes(r *mux.Router)
}
