package metereological_station

import (
	"strconv"

	"github.com/a-h/templ"
	"github.com/gorilla/mux"
	"github.com/rodrigovieira938/avipe-interfaces/pkg/app"

	"encoding/json"
	"fmt"
	"net/http"
)

type MetereologicalStationApp struct {
}

func (app *MetereologicalStationApp) Name() string {
	return "MetereologicalStationApp"
}

type DadosEstacao struct {
	EstacaoID    int     `json:"id"`
	Temperatura  float64 `json:"temperatura"`
	Precipitacao float64 `json:"precipitacao"`
	Folha        float64 `json:"folha"`
	Vento        float64 `json:"vento"`
	Direcao      string  `json:"direcao"`
	Humidade     float64 `json:"humidade"`
}

var Dados = [3]DadosEstacao{
	{EstacaoID: 0},
	{EstacaoID: 1},
	{EstacaoID: 2},
}

var estacaoNome = [3]string{
	"Azeitao",
	"Faias",
	"Algeruz",
}

func NomeEstacao(ID int) string {
	if ID < 0 || ID >= len(Dados) {
		panic(fmt.Errorf("Invalid ID"))
	}
	return estacaoNome[ID]
}

func receberDados(w http.ResponseWriter, r *http.Request) {
	var dados DadosEstacao

	err := json.NewDecoder(r.Body).Decode(&dados)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if dados.EstacaoID < 0 || dados.EstacaoID >= len(Dados) {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	Dados[dados.EstacaoID] = dados

	fmt.Println("Dados recebidos:", dados, dados.EstacaoID)

	w.WriteHeader(http.StatusOK)
}

func (app *MetereologicalStationApp) InitializeRoutes(r *mux.Router) {
	subr := r.PathPrefix("/MetereologicalStationApp").Subrouter()
	fmt.Println("Initializing routes for Metereological Station App")
	subr.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World from Metereological Station App!")
	}).Methods("GET")

	subr.HandleFunc("/interface", func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(MainPage(Dados[0])).ServeHTTP(w, r)
	}).Methods("GET")

	subr.HandleFunc("/interface/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		idStr := vars["id"]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}
		templ.Handler(MainPage(Dados[id])).ServeHTTP(w, r)
	}).Methods("GET")

	subr.HandleFunc("/dados", receberDados).Methods("POST")
}

func CreateMetereologicalStationApp() app.Application {
	return &MetereologicalStationApp{}
}
