package http_launcher

import (
	"gateway/internal/register_routes/adapter/output/repository"
	"gateway/internal/register_routes/adapter/input/controller"
	"gateway/internal/register_routes/adapter/input/routes"
	"gateway/internal/register_routes/application/usecase"
	"gateway/internal/register_routes/application/query"

	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
)

func InitRegisterRoutes(db *gorm.DB, router *gin.RouterGroup) {
	repository := repository.NewRepository(db)
	createService := usecase.NewCreateAPIService(repository)
	createRoute := usecase.NewCreateRoute(repository)
	getRoutesByServiceID := query.NewGetRouteByServiceID(db)
	controller := controller.NewController(
		*createService,
		*createRoute,
		*getRoutesByServiceID,
	)
	routes.InitRoutes(router, controller)
}