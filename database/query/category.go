package query

import (
	"fmt"

	"github.com/uyupun/regret/database"
	"github.com/uyupun/regret/models"
)

func GetCategoriesByServiceId(serviceId int) ([]models.Category, error) {
	db, err := database.ConnectGorm()
	if err != nil {
		return nil, err
	}
	categories := []models.Category{}
	db.Find(&categories, "service_id=?", serviceId)
	return categories, nil
}

func GetCategoriesByAccessToken(accessToken string) ([]models.Category, error) {
	db, err := database.ConnectGorm()
	if err != nil {
		return nil, err
	}

	service, err := GetServiceByAccessToken(accessToken)
	if err != nil {
		return nil, err
	}

	fmt.Printf("%#v\n", service)

	categories := []models.Category{}
	db.Find(&categories, "service_id=?", service.ID)
	return categories, nil
}

func AddCategory(category models.Category) error {
	db, err := database.ConnectGorm()
	if err != nil {
		return err
	}
	db.Create(&category)
	return nil
}

func EditCategory(after models.Category) error {
	db, err := database.ConnectGorm()
	if err != nil {
		return err
	}

	before := models.Category{}
	before.ID = after.ID
	db.First(&before)
	db.Model(&before).Updates(models.Category{Name: after.Name, NameJa: after.NameJa})
	return nil
}

func DeleteCategory(category models.Category) error {
	db, err := database.ConnectGorm()
	if err != nil {
		return err
	}
	db.Delete(&category)
	return nil
}
