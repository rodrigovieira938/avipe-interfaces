package metereological_station

import (
	"github.com/a-h/templ"
	"github.com/gorilla/mux"
	"github.com/rodrigovieira938/avipe-interfaces/pkg/app"

	"fmt"
	"net/http"
)

type MetereologicalStationApp struct {
}

func (app *MetereologicalStationApp) Name() string {
	return "MetereologicalStationApp"
}
func (app *MetereologicalStationApp) InitializeRoutes(r *mux.Router ) {
	subr := r.PathPrefix("/MetereologicalStationApp").Subrouter()
	fmt.Println("Initializing routes for Metereological Station App")
	subr.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World from Metereological Station App!")
	}).Methods("GET")
	subr.Handle("/interface", templ.Handler(MainPage())).Methods("GET")

}

func CreateMetereologicalStationApp() app.Application {
	return &MetereologicalStationApp{}
}
