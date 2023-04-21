package main

import (
	"github.com/eufelipemateus/api-estados-cidades/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.SetTrustedProxies(nil)
	routes.Setup(r)
	r.Run(":3001")
}
