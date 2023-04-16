package routes

import (
	"net/http"

	"github.com/eufelipemateus/api-estados-cidades/controllers"
	"github.com/gorilla/mux"
)

func Setup(app *mux.Router) {

	router := app.PathPrefix("/api").Subrouter()

	router.HandleFunc("/estados", controllers.GetEstadosList).Methods(http.MethodGet)
	router.HandleFunc("/cidades/{sigla}", controllers.GetCidadeList).Methods(http.MethodGet)

}
