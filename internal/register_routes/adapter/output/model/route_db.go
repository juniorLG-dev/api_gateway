package model

type RouteDB struct {
	ID   			 		string `gorm:"primaryKey"`
	Path 			 		string
	ServiceURL 		string
	Method     		string
	APIServiceID  string
}