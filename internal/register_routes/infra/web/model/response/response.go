package response

type RouteResponse struct {
	ID 					 string `json:"id"`
	Path 				 string `json:"path"`
	ServiceURL 	 string `json:"service_url"`
	Method       string `json:"method"`
}