package service

import (
	"gateway/internal/register_routes/application/dto"
	"gateway/internal/configuration/handler_err"

	"io"
	"encoding/json"
)

func DecodeFile(reader io.Reader) ([]dto.RouteJSON, *handler_err.InfoErr) {
	decoder := json.NewDecoder(reader)
	var routesJSON []dto.RouteJSON

	for decoder.More() {
		var routeJSON dto.RouteJSON
		if err := decoder.Decode(&routeJSON); err != nil {
			return []dto.RouteJSON{}, &handler_err.InfoErr{
				Message: "error decoding json",
				Err: handler_err.ErrInternal,
			}
		}

		routesJSON = append(routesJSON, routeJSON)
	}

	return routesJSON, &handler_err.InfoErr{}
}