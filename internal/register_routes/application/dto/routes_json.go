package dto

type RouteJSON struct {
	Path         string 	`json:"path"`
	ServiceURL   string 	`json:"service_url"`
	Method       string 	`json:"method"`
}
