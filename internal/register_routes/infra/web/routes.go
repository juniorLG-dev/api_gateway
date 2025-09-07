package web

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes(rg *gin.RouterGroup, ct ControllerGroup) {
	rg.POST("/service", ct.CreateAPIService)
	rg.POST("/routes", ct.CreateRoute)
	rg.GET("/routes", ct.GetRoutesByServiceID)
}
