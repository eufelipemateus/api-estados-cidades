package main

import (
	"fmt"

	config "github.com/eufelipemateus/api-estados-cidades/configs"
	"github.com/eufelipemateus/api-estados-cidades/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	err := config.Load()
	if err != nil {
		panic("Erro on load config.toml")
	}
	fmt.Printf("%s \nRuning in %s ", config.GetAppName(), config.GetServerPort())
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.SetTrustedProxies(nil)
	routes.Setup(r)
	r.Run(config.GetServerPort())
}
