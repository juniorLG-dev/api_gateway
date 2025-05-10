package controller 

import (
	"gateway/internal/dynamic_routing/application/usecase"
	"gateway/internal/dynamic_routing/application/dto"
	"gateway/internal/configuration/handler_err"

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

type PortController interface {
	SendRequest(c *gin.Context)
}

func (ct *controller) SendRequest(c *gin.Context) {
	serviceName := c.Param("serviceName")
	path := c.Param("path")

	sendRequestInput := dto.SendRequestInput{
		Path: path,
		ServiceName: serviceName,
		Body: c.Request.Body,
	}

	response, infoErr := ct.sendRequest.Run(sendRequestInput)
	if infoErr.Err != nil {
		msgErr := handler_err.HandlerErr(infoErr)
		c.JSON(msgErr.Status, msgErr)
		return
	}

	c.Data(response.Status, response.Header, []byte(response.Response))
}