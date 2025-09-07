package web

import (
	"gateway/internal/configuration/handler_err"
	"gateway/internal/dynamic_routing/application/dto"
	"gateway/internal/dynamic_routing/application/usecase"

	"github.com/gin-gonic/gin"
)

type controller struct {
	sendRequest usecase.SendRequest
}

func NewController(sendRequest usecase.SendRequest) *controller {
	return &controller{
		sendRequest: sendRequest,
	}
}

func (ct *controller) SendRequest(c *gin.Context) {
	serviceName := c.Param("serviceName")
	path := c.Param("path")

	sendRequestInput := dto.SendRequestInput{
		Path:        path,
		ServiceName: serviceName,
		Body:        c.Request.Body,
	}

	response, infoErr := ct.sendRequest.Run(sendRequestInput)
	if infoErr.Err != nil {
		msgErr := handler_err.HandlerErr(infoErr)
		c.JSON(msgErr.Status, msgErr)
		return
	}

	c.Data(response.Status, response.Header, []byte(response.Response))
}
