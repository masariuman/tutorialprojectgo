package main

import (
	"tutorialProject/controller"
	"tutorialProject/struct/data"

	"github.com/gin-gonic/gin"
)

func main() {
	data.ConnectDatabase()

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", controller.RootHandler)
	v1.GET("/halo", controller.HaloHandler)
	v1.GET("/data/:id", controller.DataHandler)
	v1.GET("/query", controller.QueryHandler)
	v1.GET("/data/:id/:judul", controller.DataLebihVarHandler)
	v1.POST("/data", controller.PostDataHandler)
	v1.GET("/artikel", controller.ReadHandler)

	router.Run()

}
