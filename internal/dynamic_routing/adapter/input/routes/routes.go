package routes

import (
	"gateway/internal/dynamic_routing/adapter/input/controller"

	"github.com/gin-gonic/gin"
)

func InitRoutes(rg *gin.RouterGroup, ct controller.PortController) {
	rg.Any("/:serviceName/*path", ct.SendRequest)
}