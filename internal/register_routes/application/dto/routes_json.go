package dto

type RouteJSON struct {
	APIName 		string `json:"api_name"`
	Path    		string `json:"path"`
	ServiceURL  string `json:"service_url"`
	Method      string `json:"method"`
}