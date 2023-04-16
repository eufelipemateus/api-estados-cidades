package controllers

import (
	"net/http"

	"github.com/eufelipemateus/api-estados-cidades/file"
	"github.com/gin-gonic/gin"

	"github.com/eufelipemateus/api-estados-cidades/interfaces"
)

func GetEstadosList(c *gin.Context) {
	var estado []interfaces.Estado

	data := file.ReadJson("./data/estados.new.json", estado)

	c.JSON(http.StatusOK,  data)
}
