package usecase

import (
	"gateway/internal/register_routes/application/domain/entities"
	"gateway/internal/register_routes/application/domain/service"
	"gateway/internal/register_routes/adapter/output/repository"
	"gateway/internal/configuration/handler_err"

	"io"
)

type CreateRoute struct {
	repository repository.PortRepository
}

func NewCreateRoute(repository repository.PortRepository) *CreateRoute {
	return &CreateRoute{
		repository: repository,
	}
}

func (cr *CreateRoute) Run(filename string, file io.Reader) *handler_err.InfoErr {
	checkFile := service.NewCheckFile(filename)
	if !checkFile.Check(".json") {
		return &handler_err.InfoErr{
			Message: "invalid extension",
			Err: handler_err.ErrInvalidInput,
		}
	}

	routeJSON, msgErr := service.DecodeFile(file)
	if msgErr.Err != nil {
		return msgErr
	}

	route, msgErr := entities.NewRoute(
		routeJSON.APIName,
		routeJSON.Path,
		routeJSON.ServiceURL,
		routeJSON.Method,
	)
	if msgErr.Err != nil {
		return msgErr
	}

	if err := cr.repository.CreateRoute(*route); err != nil {
		return &handler_err.InfoErr{
			Message: "could not create the route",
			Err: handler_err.ErrInternal,
		}
	}

	return &handler_err.InfoErr{}
}