package usecase

import (
	"gateway/internal/configuration/handler_err"
	"gateway/internal/dynamic_routing/adapter/output/api"
	"gateway/internal/dynamic_routing/adapter/output/gateway"
	"gateway/internal/dynamic_routing/application/dto"

	"fmt"
	"io/ioutil"
	"strings"
)

type SendRequest struct {
	gateway gateway.PortGateway
}

func NewSendRequest(gateway gateway.PortGateway) *SendRequest {
	return &SendRequest{
		gateway: gateway,
	}
}

func (sr *SendRequest) definePathForQuery(pathArray []string) string {
	pathQuery := pathArray[1]
	if len(pathArray) > 2 && pathArray[2] != "" {
		pathQuery = pathArray[1]
	}
	return pathQuery
}

func (sr *SendRequest) Run(sendRequestInput dto.SendRequestInput) (dto.SendRequestOutput, *handler_err.InfoErr) {
	bodyBytes, err := ioutil.ReadAll(sendRequestInput.Body)
	if err != nil {
		return dto.SendRequestOutput{}, &handler_err.InfoErr{
			Message: "could not read the request body",
			Err:     handler_err.ErrInternal,
		}
	}

	pathArray := strings.Split(sendRequestInput.Path, "/")
	pathQuery := sr.definePathForQuery(pathArray)

	serviceID, err := sr.gateway.GetServiceByName(sendRequestInput.ServiceName)
	if err != nil {
		return dto.SendRequestOutput{}, &handler_err.InfoErr{
			Message: "service not found",
			Err:     handler_err.ErrNotFound,
		}
	}

	routeInfo, err := sr.gateway.GetRouteByServiceAndPath(serviceID, pathQuery)
	if err != nil {
		return dto.SendRequestOutput{}, &handler_err.InfoErr{
			Message: "this path does not exist in this particular service",
			Err:     handler_err.ErrNotFound,
		}
	}

	backendURL := fmt.Sprintf("%s%s", routeInfo.ServiceURL, pathArray[1])
	if len(pathArray) > 2 && pathArray[2] != "" {
		backendURL = fmt.Sprintf("%s%s/%s", routeInfo.ServiceURL, pathArray[1], pathArray[2])
	}

	apiReq := api.NewAPIReq(
		routeInfo.Method,
		backendURL,
		bodyBytes,
	)

	response, err := apiReq.SendRequest()
	if err != nil {
		fmt.Println(err)
		return dto.SendRequestOutput{}, &handler_err.InfoErr{
			Message: "could not perform the request",
			Err:     handler_err.ErrInternal,
		}
	}

	responseReturn := dto.SendRequestOutput{
		Response: response.Response,
		Status:   response.Status,
		Header:   response.Header,
	}

	return responseReturn, &handler_err.InfoErr{}
}
