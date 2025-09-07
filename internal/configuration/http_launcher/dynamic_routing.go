package http_launcher

import (
	"gateway/internal/dynamic_routing/application/usecase"
	"gateway/internal/dynamic_routing/infra/gateway"
	"gateway/internal/dynamic_routing/infra/web"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitDynamicRouting(db *gorm.DB, router *gin.RouterGroup) {
	gateway := gateway.NewGateway(db)
	sendRequest := usecase.NewSendRequest(gateway)
	controller := web.NewController(
		*sendRequest,
	)
	web.InitRoutes(router, controller)
}
