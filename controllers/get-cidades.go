package controllers

import (
	"net/http"
	"strings"

	"github.com/eufelipemateus/api-estados-cidades/file"
	"github.com/gin-gonic/gin"

	"github.com/eufelipemateus/api-estados-cidades/interfaces"
)

func GetCidadeList(c *gin.Context) {
	var cidade []interfaces.Cidade
	var filtered []interfaces.Cidade

	data := file.ReadJson("./data/cidades.new.json", cidade)

	sigla := c.Param("sigla")
	sigla = strings.ToUpper(sigla)

	for _, val := range data {
		if val.Estado == strings.ToUpper(sigla) {
			filtered = append(filtered, val)
		}
	}
	c.JSON(http.StatusOK, filtered)

}
