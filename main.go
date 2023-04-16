package main

import (
	"github.com/eufelipemateus/api-estados-cidades/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.Setup(r)
	r.Run(":8080")
}
