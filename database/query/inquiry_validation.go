package query

import (
	"github.com/uyupun/regret/database"
	"github.com/uyupun/regret/models"
)

func AddInquiryValidation(serviceId int) error {
	inquiryValidation := new(models.InquiryValidation)
	inquiryValidation.ServiceId = serviceId
	db, err := database.ConnectGorm()
	if err != nil {
		return err
	}
	db.Create(&inquiryValidation)
	return nil
}
