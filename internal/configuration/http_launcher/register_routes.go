package http_launcher

import (
	"gateway/internal/register_routes/application/query"
	"gateway/internal/register_routes/application/usecase"
	"gateway/internal/register_routes/infra/repository"
	"gateway/internal/register_routes/infra/web"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRegisterRoutes(db *gorm.DB, router *gin.RouterGroup) {
	repository := repository.NewRepository(db)
	createService := usecase.NewCreateAPIService(repository)
	createRoute := usecase.NewCreateRoute(repository)
	getRoutesByServiceID := query.NewGetRouteByServiceID(db)
	controller := web.NewController(
		*createService,
		*createRoute,
		*getRoutesByServiceID,
	)
	web.InitRoutes(router, controller)
}
