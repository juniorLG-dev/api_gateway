package repository

import (
	"gateway/internal/register_routes/application/domain/entities"
	"gateway/internal/register_routes/adapter/output/model"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

type PortRepository interface {
	CreateAPIService(entities.APIService) error
	CreateRoute([]entities.Route) error
}

func (r *repository) CreateAPIService(apiService entities.APIService) error {
	apiServiceDB := model.APIServiceDB{
		ID: apiService.GetID(),
		Name: apiService.GetName(),
	}

	return r.db.Create(&apiServiceDB).Error
}

func (r *repository) CreateRoute(routes []entities.Route) error {
	var routesDB []model.RouteDB
	for _, route := range routes {
		routeInfo := model.RouteDB{
			ID: route.GetID(),
			Path: route.GetPath(),
			ServiceURL: route.GetServiceURL(),
			Method: route.GetMethod(),
			APIServiceID: route.GetAPIServiceID(),
		}

		routesDB = append(routesDB, routeInfo)
	}

	return  r.db.Create(&routesDB).Error
}