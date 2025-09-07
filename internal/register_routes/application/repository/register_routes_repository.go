package repository

import "gateway/internal/register_routes/domain/entities"

type RegisterRoutesRepository interface {
	CreateAPIService(entities.APIService) error
	CreateRoute([]entities.Route) error
}
