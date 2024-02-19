package sensors

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	userGroup := router.Group("/sensors")
	{
		userGroup.POST("/", func(ctx *gin.Context) {
			sensorCallback(ctx, getReadings)
		})
		// userGroup.POST("/", CreateUser)
	}
}
