package entities

import (
	"gateway/internal/register_routes/application/domain/vo"
	"gateway/internal/configuration/handler_err"
)

type Route struct {
	ID 				 	 vo.ID
	Path 			 	 string
	ServiceURL 	 vo.ServiceURL
	Method     	 string
	APIServiceID string
}

func NewRoute(
	path 			 string,
	service 	 string,
	method 		 string,
	apiServiceID string,
) (*Route, *handler_err.InfoErr) {
	serviceURL, msgErr := vo.NewServiceURL(service)
	if msgErr.Err != nil {
		return &Route{}, msgErr
	}

	return &Route{
		ID: *vo.NewID(),
		Path: path,
		ServiceURL: *serviceURL,
		Method: method,
		APIServiceID: apiServiceID,
	}, &handler_err.InfoErr{}
}

func (r *Route) GetID() string {
	return r.ID.Value
}

func (r *Route) GetPath() string {
	return r.Path
}

func (r *Route) GetServiceURL() string {
	return r.ServiceURL.Value
}

func (r *Route) GetMethod() string {
	return r.Method
}

func (r *Route) GetAPIServiceID() string {
	return r.APIServiceID
}