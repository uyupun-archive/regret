package query

import (
	"github.com/uyupun/regret/models"
)

func GetServices() ([]models.Service, error) {
	db, err := connectGorm()
	if err != nil {
		return nil, err
	}
	services := []models.Service{}
	db.Find(&services)
	return services, nil
}

func AddService(model models.Service) error {
	db, err := connectGorm()
	if err != nil {
		return err
	}
	db.Create(&model)
	return nil
}
