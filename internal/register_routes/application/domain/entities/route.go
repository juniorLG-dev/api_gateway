package entities

import (
	"gateway/internal/register_routes/application/domain/vo"
	"gateway/internal/configuration/handler_err"
)

type Route struct {
	ID 				 vo.ID
	APIName    string
	Path 			 string
	ServiceURL vo.ServiceURL
	Method     string
}

func NewRoute(
	APIName    string,
	path 			 string,
	service 	 string,
	method 		 string,
) (*Route, *handler_err.InfoErr) {
	serviceURL, msgErr := vo.NewServiceURL(service)
	if msgErr.Err != nil {
		return &Route{}, msgErr
	}

	return &Route{
		ID: *vo.NewID(),
		APIName: APIName,
		Path: path,
		ServiceURL: *serviceURL,
		Method: method,
	}, &handler_err.InfoErr{}
}