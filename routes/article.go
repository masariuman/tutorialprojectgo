package routes

import (
	"tutorialProject/controller"

	"github.com/gin-gonic/gin"
)

func articleRoute(superRoute *gin.RouterGroup) {
	articlesRoute := superRoute.Group("/article")
	{
		articlesRoute.POST("", controller.PostDataHandler)
		articlesRoute.GET("", controller.ReadHandler)
		articlesRoute.GET("/:id", controller.ReadidHandler)
		articlesRoute.PUT("/:id/update", controller.UpdateHandler)
		articlesRoute.DELETE("/", controller.DeleteHandler)
	}
}
