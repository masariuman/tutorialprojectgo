package routes

import (
	"tutorialProject/controller"

	"github.com/gin-gonic/gin"
)

func tutorialRoute(superRoute *gin.RouterGroup) {
	tutorialRoute := superRoute.Group("/tutorial")
	{
		tutorialRoute.GET("/", controller.RootHandler)
		tutorialRoute.GET("/halo", controller.HaloHandler)
		tutorialRoute.GET("/data/:id", controller.DataHandler)
		tutorialRoute.GET("/query", controller.QueryHandler)
		tutorialRoute.GET("/data/:id/:judul", controller.DataLebihVarHandler)
	}
}
