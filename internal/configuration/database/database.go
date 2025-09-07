package database

import (
	"gateway/internal/register_routes/infra/repository"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupDB() (*gorm.DB, error) {
	db, err := gorm.Open(
		sqlite.Open("routes.db"),
		&gorm.Config{},
	)
	if err != nil {
		return nil, err
	}

	if err = db.AutoMigrate(&repository.RouteDB{}, &repository.APIServiceDB{}); err != nil {
		return nil, err
	}

	return db, nil
}
