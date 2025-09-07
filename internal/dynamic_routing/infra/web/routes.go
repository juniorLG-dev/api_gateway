package web

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes(rg *gin.RouterGroup, ct ControllerGroup) {
	rg.Any("/:serviceName/*path", ct.SendRequest)
}
