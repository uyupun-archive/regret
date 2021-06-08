package query

import (
	"github.com/uyupun/regret/database"
	"github.com/uyupun/regret/models"
)

func GetServices() ([]models.Service, error) {
	db, err := database.ConnectGorm()
	if err != nil {
		return nil, err
	}
	services := []models.Service{}
	db.Find(&services)
	return services, nil
}

func GetServiceByAccessToken(accessToken string) (models.Service, error) {
	service := models.Service{}
	db, err := database.ConnectGorm()
	if err != nil {
		return service, err
	}
	db.First(&service, "access_token=?", accessToken)
	return service, nil
}

func AddService(service models.Service) error {
	db, err := database.ConnectGorm()
	if err != nil {
		return err
	}
	db.Create(&service)
	return nil
}

func EditService(after models.Service) error {
	db, err := database.ConnectGorm()
	if err != nil {
		return err
	}

	before := models.Service{}
	before.ID = after.ID
	db.First(&before)
	db.Model(&before).Updates(models.Service{Name: after.Name, NameJa: after.NameJa, Description: after.Description})
	return nil
}

func DeleteService(service models.Service) error {
	db, err := database.ConnectGorm()
	if err != nil {
		return err
	}
	db.Delete(&service)
	return nil
}
