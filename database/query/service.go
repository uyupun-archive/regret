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

func AddService(service models.Service) error {
	db, err := connectGorm()
	if err != nil {
		return err
	}
	db.Create(&service)
	return nil
}

func DeleteService(service models.Service) error {
	db, err := connectGorm()
	if err != nil {
		return err
	}
	db.Delete(&service)
	return nil
}
