package model

type APIServiceDB struct {
	ID 	 string  `gorm:"primaryKey"`
	Name string
}