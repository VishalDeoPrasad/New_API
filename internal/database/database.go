package database

import (
	"job-application-api/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Open() (*gorm.DB, error) {
	database := "host=localhost user=postgres password=admin dbname=NewDatabase port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(database))
	if err != nil {
		return nil, err
	}

	err = db.Migrator().AutoMigrate(&models.User{})
	if err != nil {
		return nil, err
	}

	err = db.Migrator().AutoMigrate(&models.Company{})
	if err != nil {
		return nil, err
	}
	err = db.Migrator().AutoMigrate(&models.Job{}, &models.Location{}, &models.Qualification{}, &models.Shift{}, &models.TechnologyStack{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
