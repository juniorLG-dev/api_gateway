package controller

import (
	"gateway/internal/register_routes/application/usecase"
	"gateway/internal/register_routes/application/query"
	"gateway/internal/configuration/handler_err"
	"gateway/internal/register_routes/adapter/input/dto"

	"github.com/gin-gonic/gin"
	
	"net/http"
	"fmt"
)

type controller struct {
	createRoute 	 usecase.CreateRoute
	getRouteByName query.GetRouteByName
}

func NewController(
	createRoute usecase.CreateRoute,
	getRouteByName query.GetRouteByName,
) *controller {
	return &controller{
		createRoute: createRoute,
		getRouteByName: getRouteByName,
	}
}

type PortController interface {
	CreateRoute(c *gin.Context)
	GetRoutesByName(c *gin.Context)
}

func (ct *controller) CreateRoute(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "could not receive file"})
		return
	}

	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not open file"})
		return
	}
	defer src.Close()

	if infoErr := ct.createRoute.Run(file.Filename, src); infoErr.Err != nil {
		msgErr := handler_err.HandlerErr(infoErr)
		c.JSON(msgErr.Status, msgErr)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "route created"})
}

func (ct *controller) GetRoutesByName(c *gin.Context) {
	apiName := c.Param("apiName")

	routes, infoErr := ct.getRouteByName.Run(apiName)
	if infoErr.Err != nil {
		fmt.Println(infoErr)
		msgErr := handler_err.HandlerErr(infoErr)
		c.JSON(msgErr.Status, msgErr)
		return
	}

	var routesResponse []dto.RouteResponse
	for _, route := range routes {
		routeInfo := dto.RouteResponse{
			ID: route.ID,
			APIName: route.APIName,
			Path: route.Path,
			ServiceURL: route.ServiceURL,
		}

		routesResponse = append(routesResponse, routeInfo)
	}

	c.JSON(http.StatusOK, routesResponse)
}