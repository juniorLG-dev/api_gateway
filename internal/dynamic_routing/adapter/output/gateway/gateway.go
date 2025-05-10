package gateway

import (
	"gateway/internal/register_routes/adapter/output/model"

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

type PortGateway interface {
	GetServiceByName(string) (string, error)
	GetRouteByServiceAndPath(string, string) (GatewayInfoOutput, error)
}

type GatewayInfoOutput struct {
	Path 			 string
	ServiceURL string
	Method 		 string
}

func (g *gateway) GetServiceByName(name string) (string, error) {
	var service model.APIServiceDB
  err := g.db.First(&service, "name = ?", name).Error

  return service.ID, err
}

func (g *gateway) GetRouteByServiceAndPath(serviceID, path string) (GatewayInfoOutput, error) {
	var route model.RouteDB
	err := g.db.Where("api_service_id = ? AND path = ?", serviceID, path).First(&route).Error

	routeOutput := GatewayInfoOutput{
		Path: path,
		ServiceURL: route.ServiceURL,
		Method: route.Method,
	}

	return routeOutput, err
}