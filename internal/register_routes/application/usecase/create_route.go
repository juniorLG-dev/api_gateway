package usecase

import (
	"gateway/internal/register_routes/application/domain/entities"
	"gateway/internal/register_routes/application/domain/service"
	"gateway/internal/register_routes/application/dto"
	"gateway/internal/register_routes/adapter/output/repository"
	"gateway/internal/configuration/handler_err"
)

type CreateRoute struct {
	repository repository.PortRepository
}

func NewCreateRoute(repository repository.PortRepository) *CreateRoute {
	return &CreateRoute{
		repository: repository,
	}
}

func (cr *CreateRoute) Run(routeInput dto.CreateRouteInput) *handler_err.InfoErr {
	tokenGenerator := service.NewTokenGenerator()
	apiService, msgErr := tokenGenerator.CheckToken(routeInput.Token)
	if msgErr.Err != nil {
		return msgErr
	}


	checkFile := service.NewCheckFile(routeInput.Filename)
	if !checkFile.Check(".json") {
		return &handler_err.InfoErr{
			Message: "invalid extension",
			Err: handler_err.ErrInvalidInput,
		}
	}

	routesJSON, msgErr := service.DecodeFile(routeInput.File)
	if msgErr.Err != nil {
		return msgErr
	}

	var routes []entities.Route
	for _, route := range routesJSON {
		routeInfo, msgErr := entities.NewRoute(
			route.Path,
			route.ServiceURL,
			route.Method,
			apiService.ID,
		)
		if msgErr.Err != nil {
			return msgErr
		}

		routes = append(routes, *routeInfo)
	}

	if err := cr.repository.CreateRoute(routes); err != nil {
		return &handler_err.InfoErr{
			Message: "could not create the route",
			Err: handler_err.ErrInternal,
		}
	}

	return &handler_err.InfoErr{}
}