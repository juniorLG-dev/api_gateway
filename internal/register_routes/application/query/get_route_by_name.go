package query

import (
	"gateway/internal/register_routes/adapter/output/model"
	"gateway/internal/register_routes/application/dto"
	"gateway/internal/configuration/handler_err"

	"gorm.io/gorm"

	"fmt"
)

type GetRouteByName struct {
	db *gorm.DB
}

func NewGetRouteByName(db *gorm.DB) *GetRouteByName {
	return &GetRouteByName{
		db: db,
	}
}

func (gr *GetRouteByName) Run(apiName string) ([]dto.GetRouteByNameOutput, *handler_err.InfoErr) {
	var routes []model.RouteDB
	if err := gr.db.Where("api_name = ?", apiName).Find(&routes).Error; err != nil {
		fmt.Println("erro ", err)
		return []dto.GetRouteByNameOutput{}, &handler_err.InfoErr{
			Message: "error querying routes",
			Err: handler_err.ErrInternal,
		}
	}

	if len(routes) == 0 {
		return []dto.GetRouteByNameOutput{}, &handler_err.InfoErr{
			Message: "routes not found",
			Err: handler_err.ErrNotFound,
		}
	}
	var routesOutput []dto.GetRouteByNameOutput
	for _, route := range routes {
		routeInfo := dto.GetRouteByNameOutput{
			ID: route.ID,
			APIName: route.APIName,
			Path: route.Path,
			ServiceURL: route.ServiceURL,
			Method: route.Method,
		}

		routesOutput = append(routesOutput, routeInfo)
	}

	return routesOutput, &handler_err.InfoErr{}
}
