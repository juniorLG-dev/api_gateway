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
	CreateRoute(entities.Route) error
}

func (r *repository) CreateAPIService(apiService entities.APIService) error {
	apiServiceDB := model.APIServiceDB{
		ID: apiService.ID.Value,
		Name: apiService.Name.Value,
	}

	return r.db.Create(&apiServiceDB).Error
}

func (r *repository) CreateRoute(route entities.Route) error {
	routeDB := model.RouteDB{
		ID: route.ID.Value,
		APIName: route.APIName,
		Path: route.Path,
		ServiceURL: route.ServiceURL.Value,
		Method: route.Method,
	}

	return  r.db.Create(&routeDB).Error
}