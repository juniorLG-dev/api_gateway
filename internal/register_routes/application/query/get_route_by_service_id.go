package query

import (
	"gateway/internal/configuration/handler_err"
	"gateway/internal/register_routes/application/dto"
	"gateway/internal/register_routes/domain/service"
	"gateway/internal/register_routes/infra/repository"

	"gorm.io/gorm"
)

type GetRouteByServiceID struct {
	db *gorm.DB
}

func NewGetRouteByServiceID(db *gorm.DB) *GetRouteByServiceID {
	return &GetRouteByServiceID{
		db: db,
	}
}

func (gr *GetRouteByServiceID) Run(token string) ([]dto.GetRouteByServiceIDOutput, *handler_err.InfoErr) {
	tokenGenerator := service.NewTokenGenerator()
	apiService, msgErr := tokenGenerator.CheckToken(token)
	if msgErr.Err != nil {
		return []dto.GetRouteByServiceIDOutput{}, msgErr
	}
	var routes []repository.RouteDB
	if err := gr.db.Where("api_service_id = ?", apiService.ID).Find(&routes).Error; err != nil {
		return []dto.GetRouteByServiceIDOutput{}, &handler_err.InfoErr{
			Message: "error querying routes",
			Err:     handler_err.ErrInternal,
		}
	}

	if len(routes) == 0 {
		return []dto.GetRouteByServiceIDOutput{}, &handler_err.InfoErr{
			Message: "routes not found",
			Err:     handler_err.ErrNotFound,
		}
	}
	var routesOutput []dto.GetRouteByServiceIDOutput
	for _, route := range routes {
		routeInfo := dto.GetRouteByServiceIDOutput{
			ID:         route.ID,
			Path:       route.Path,
			ServiceURL: route.ServiceURL,
			Method:     route.Method,
		}

		routesOutput = append(routesOutput, routeInfo)
	}

	return routesOutput, &handler_err.InfoErr{}
}
