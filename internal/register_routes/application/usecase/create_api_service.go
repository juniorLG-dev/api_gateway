package usecase

import (
	"gateway/internal/register_routes/application/domain/entities"
	"gateway/internal/register_routes/application/domain/service"
	"gateway/internal/register_routes/adapter/output/repository"
	"gateway/internal/configuration/handler_err"
)

type CreateAPIService struct {
	repository repository.PortRepository
}

func NewCreateAPIService(repository repository.PortRepository) *CreateAPIService {
	return &CreateAPIService{
		repository: repository,
	}
}

func (cas *CreateAPIService) Run(name string) (string, *handler_err.InfoErr) {
	apiService, msgErr := entities.NewAPIService(name)
	if msgErr.Err != nil {
		return "", msgErr
	}

	if err := cas.repository.CreateAPIService(*apiService); err != nil {
		return "", &handler_err.InfoErr{
			Message: "could not create your api service",
			Err: handler_err.ErrInternal,
		}
	}

	tokenGenerator := service.NewTokenGenerator()

	token, msgErr := tokenGenerator.GenerateToken(*apiService)
	if msgErr.Err != nil {
		return "", msgErr
	}

	return token, &handler_err.InfoErr{}
}