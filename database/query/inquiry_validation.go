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

func GetInquiryValidationByServiceId(serviceId int) (models.InquiryValidation, error) {
	inquiryValidation := models.InquiryValidation{}
	db, err := database.ConnectGorm()
	if err != nil {
		return inquiryValidation, err
	}
	db.First(&inquiryValidation, "service_id=?", serviceId)
	return inquiryValidation, nil
}

func EditInquiryValidation(after models.InquiryValidation) error {
	db, err := database.ConnectGorm()
	if err != nil {
		return err
	}

	before := models.InquiryValidation{}
	before.ID = after.ID
	db.First(&before)
	db.Model(&before).Updates(map[string]interface{}{"IsRequiredSubject": after.IsRequiredSubject, "IsRequiredEmail": after.IsRequiredEmail, "IsRequiredCategory": after.IsRequiredCategory, "IsRequiredText": after.IsRequiredText})
	return nil
}
