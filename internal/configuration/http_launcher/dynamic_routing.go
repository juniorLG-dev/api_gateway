package http_launcher

import (
	"gateway/internal/dynamic_routing/adapter/output/gateway"
	"gateway/internal/dynamic_routing/application/usecase"
	"gateway/internal/dynamic_routing/adapter/input/controller"
	"gateway/internal/dynamic_routing/adapter/input/routes"

	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
)

func InitDynamicRouting(db *gorm.DB, router *gin.RouterGroup) {
	gateway := gateway.NewGateway(db)
	sendRequest := usecase.NewSendRequest(gateway)
	controller := controller.NewController(
		*sendRequest,
	)
	routes.InitRoutes(router, controller)
}