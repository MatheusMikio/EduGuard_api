package routes

import (
	"github.com/MatheusMikio/eduGuard_api/internal/container"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Init(db *gorm.DB) {
	container := container.NewContainer(db)
	router := gin.Default()

	initRoutes(container, router)
	router.Run(":8080")
}
