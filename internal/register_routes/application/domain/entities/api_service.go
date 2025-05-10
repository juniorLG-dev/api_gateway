package entities

import (
	"gateway/internal/register_routes/application/domain/vo"
	"gateway/internal/configuration/handler_err"
)

type APIService struct {
	ID 	 vo.ID
	Name vo.Name
}

func NewAPIService(name string) (*APIService, *handler_err.InfoErr) {
	serviceName, msgErr := vo.NewName(name)
	if msgErr.Err != nil {
		return &APIService{}, msgErr
	}

	return &APIService{
		ID: *vo.NewID(),
		Name: *serviceName,
	}, &handler_err.InfoErr{}
}

func (as *APIService) GetID() string {
	return as.ID.Value
}

func (as *APIService) GetName() string {
	return as.Name.Value
}