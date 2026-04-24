package routes

import (
	"github.com/MatheusMikio/eduGuard_api/internal/container"
	"github.com/gin-gonic/gin"
)

func initRoutes(container *container.Container, router *gin.Engine) {
	basePath := "/api/v1"

	v1 := router.Group(basePath)
	{
		v1.POST("/")
	}
}
