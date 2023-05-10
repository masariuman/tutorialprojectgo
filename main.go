package main

import (
	"tutorialProject/routes"
	"tutorialProject/setup"

	"github.com/gin-gonic/gin"
)

func main() {
	setup.ConnectDatabase()

	router := gin.Default()
	v1 := router.Group("/v1")
	routes.AddRoutes(v1)

	router.Run()

}
