package routes

import (
	"github.com/gin-gonic/gin"
)

func AddRoutes(superRoute *gin.RouterGroup) {
	articleRoute(superRoute)
	tutorialRoute(superRoute)
}
