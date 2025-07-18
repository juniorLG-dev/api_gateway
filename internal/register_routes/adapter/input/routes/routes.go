package routes

import (
	"gateway/internal/register_routes/adapter/input/controller"

	"github.com/gin-gonic/gin"
)

func InitRoutes(rg *gin.RouterGroup, ct controller.PortController) {
	rg.POST("/service", ct.CreateAPIService)
	rg.POST("/routes", ct.CreateRoute)
	rg.GET("/routes", ct.GetRoutesByServiceID)
}
