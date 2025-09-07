package web

import "github.com/gin-gonic/gin"

type ControllerGroup interface {
	SendRequest(c *gin.Context)
}
