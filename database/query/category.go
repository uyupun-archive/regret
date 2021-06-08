package query

import "github.com/uyupun/regret/models"

func GetCategories(serviceId int) ([]models.Category, error) {
	db, err := connectGorm()
	if err != nil {
		return nil, err
	}
	categories := []models.Category{}
	db.Find(&categories, "service_id=?", serviceId)
	return categories, nil
}

func AddCategory(category models.Category) error {
	db, err := connectGorm()
	if err != nil {
		return err
	}
	db.Create(&category)
	return nil
}
