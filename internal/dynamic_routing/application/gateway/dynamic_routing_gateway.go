package gateway

import "gateway/internal/dynamic_routing/infra/gateway"

type DynamicRoutingGateway interface {
	GetServiceByName(string) (string, error)
	GetRouteByServiceAndPath(string, string) (gateway.GatewayInfoOutput, error)
}
