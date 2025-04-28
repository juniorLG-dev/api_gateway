package database

import (
	"gateway/internal/register_routes/adapter/output/model"

	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

func SetupDB() (*gorm.DB, error) {
	db, err := gorm.Open(
		sqlite.Open("routes.db"),
		&gorm.Config{},
	)
	if err != nil {
		return nil, err
	}

	if err = db.AutoMigrate(&model.RouteDB{}); err != nil {
		return nil, err
	}

	return db, nil
}
