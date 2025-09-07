package web

import (
	"gateway/internal/configuration/handler_err"
	"gateway/internal/register_routes/application/dto"
	"gateway/internal/register_routes/application/query"
	"gateway/internal/register_routes/application/usecase"
	"gateway/internal/register_routes/infra/web/model/request"
	"gateway/internal/register_routes/infra/web/model/response"

	"github.com/gin-gonic/gin"

	"net/http"
)

type controller struct {
	createAPIService    usecase.CreateAPIService
	createRoute         usecase.CreateRoute
	getRouteByServiceID query.GetRouteByServiceID
}

func NewController(
	createAPIService usecase.CreateAPIService,
	createRoute usecase.CreateRoute,
	getRouteByServiceID query.GetRouteByServiceID,
) *controller {
	return &controller{
		createAPIService:    createAPIService,
		createRoute:         createRoute,
		getRouteByServiceID: getRouteByServiceID,
	}
}

func (ct *controller) CreateAPIService(c *gin.Context) {
	var createAPIService request.CreateAPIServiceRequest
	if err := c.ShouldBindJSON(&createAPIService); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid fields"})
		return
	}

	token, infoErr := ct.createAPIService.Run(createAPIService.Name)
	if infoErr.Err != nil {
		msgErr := handler_err.HandlerErr(infoErr)
		c.JSON(msgErr.Status, msgErr)
		return
	}

	c.Header("Authorization", token)
	c.JSON(http.StatusCreated, gin.H{"message": "service created"})
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

	createRouteInput := dto.CreateRouteInput{
		Filename: file.Filename,
		File:     src,
		Token:    c.Request.Header.Get("Authorization"),
	}

	if infoErr := ct.createRoute.Run(createRouteInput); infoErr.Err != nil {
		msgErr := handler_err.HandlerErr(infoErr)
		c.JSON(msgErr.Status, msgErr)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "route created"})
}

func (ct *controller) GetRoutesByServiceID(c *gin.Context) {
	routes, infoErr := ct.getRouteByServiceID.Run(c.Request.Header.Get("Authorization"))
	if infoErr.Err != nil {
		msgErr := handler_err.HandlerErr(infoErr)
		c.JSON(msgErr.Status, msgErr)
		return
	}

	var routesResponse []response.RouteResponse
	for _, route := range routes {
		routeInfo := response.RouteResponse{
			ID:         route.ID,
			Path:       route.Path,
			ServiceURL: route.ServiceURL,
			Method:     route.Method,
		}

		routesResponse = append(routesResponse, routeInfo)
	}

	c.JSON(http.StatusOK, routesResponse)
}
