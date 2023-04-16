package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/eufelipemateus/api-estados-cidades/file"

	"github.com/eufelipemateus/api-estados-cidades/interfaces"
)

func GetEstadosList(w http.ResponseWriter, r *http.Request) {
	var estado []interfaces.Estado
	
	data := file.ReadJson("./data/estados.new.json", estado)
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)

}
