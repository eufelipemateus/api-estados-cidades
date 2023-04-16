package routes

import (
	"github.com/eufelipemateus/api-estados-cidades/controllers"
	"github.com/gin-gonic/gin"
)

func Setup(app *gin.Engine) {

	router := app.Group("/api")

	router.GET("/estados", controllers.GetEstadosList)
	router.GET("/cidades/:sigla", controllers.GetCidadeList)

}
