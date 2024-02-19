package sensors

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ServiceBody struct {
	Reading string
}

type serviceCallback func(*ServiceBody) error

func sensorCallback(ctx *gin.Context, service serviceCallback) {
	bodyArgs := ServiceBody{}

	if err := ctx.Bind(&bodyArgs); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "error parsing JSON",
		})

		return
	}

	fmt.Println(bodyArgs)

	if error := service(&bodyArgs); error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": error.Error(),
		})

		return
	}
}
