package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kaushik13k/rest-GoApi/config"
	"github.com/kaushik13k/rest-GoApi/routes"
)

func main() {
	router := gin.New()
	config.Connect()
	routes.UserRoute(router)
	router.Run(":8080")
}
