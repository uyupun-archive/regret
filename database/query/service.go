package query

import (
	"github.com/uyupun/regret/models"
)

func AddService(model models.Service) error {
	db, err := connectGorm()
	if err != nil {
		return err
	}
	db.Create(&model)
	return nil
}
