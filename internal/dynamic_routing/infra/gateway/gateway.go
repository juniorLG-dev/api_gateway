package gateway

import (
	"gateway/internal/register_routes/infra/repository"

	"gorm.io/gorm"
)

type gateway struct {
	db *gorm.DB
}

func NewGateway(db *gorm.DB) *gateway {
	return &gateway{
		db: db,
	}
}

type GatewayInfoOutput struct {
	Path       string
	ServiceURL string
	Method     string
}

func (g *gateway) GetServiceByName(name string) (string, error) {
	var service repository.APIServiceDB
	err := g.db.First(&service, "name = ?", name).Error

	return service.ID, err
}

func (g *gateway) GetRouteByServiceAndPath(serviceID, path string) (GatewayInfoOutput, error) {
	var route repository.RouteDB
	err := g.db.Where("api_service_id = ? AND path = ?", serviceID, path).First(&route).Error

	routeOutput := GatewayInfoOutput{
		Path:       path,
		ServiceURL: route.ServiceURL,
		Method:     route.Method,
	}

	return routeOutput, err
}
