package main

import (
	"fmt"
	"net/http"

	"github.com/eufelipemateus/api-estados-cidades/routes"
	"github.com/gorilla/mux"

)

func main() {

	r := mux.NewRouter()
	routes.Setup(r)

	r.Use(mux.CORSMethodMiddleware(r))

	fmt.Println(http.ListenAndServe(":8080", r))
}
