package repository

type APIServiceDB struct {
	ID   string `gorm:"primaryKey"`
	Name string
}

type RouteDB struct {
	ID           string `gorm:"primaryKey"`
	Path         string
	ServiceURL   string
	Method       string
	APIServiceID string
}
