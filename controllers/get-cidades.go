package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/eufelipemateus/api-estados-cidades/file"

	"github.com/eufelipemateus/api-estados-cidades/interfaces"
	"github.com/gorilla/mux"
)

func GetCidadeList(w http.ResponseWriter, r *http.Request) {
	var cidade []interfaces.Cidade
	var filtered []interfaces.Cidade

	data := file.ReadJson("./data/cidades.new.json", cidade)

	sigla := mux.Vars(r)["sigla"]
	sigla = strings.ToUpper(sigla)

	for _, val := range data {
		if val.Estado == strings.ToUpper(sigla) {
			filtered = append(filtered, val)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(filtered)

}
