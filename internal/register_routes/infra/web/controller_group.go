package web

import "github.com/gin-gonic/gin"

type ControllerGroup interface {
	CreateAPIService(c *gin.Context)
	CreateRoute(c *gin.Context)
	GetRoutesByServiceID(c *gin.Context)
}
