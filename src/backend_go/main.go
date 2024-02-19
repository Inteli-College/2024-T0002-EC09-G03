package main

import (
	"github.com/gin-gonic/gin"
  "grupo-3-backend/sensors"
)

func main() {
	router := gin.Default()
  
  sensors.SetupRoutes(router)

	router.Run("0.0.0.0:3000")
}
