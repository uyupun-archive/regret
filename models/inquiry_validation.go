package models

type InquiryValidation struct {
	ID                 int  `json:"id" yaml:"id"`
	ServiceId          int  `json:"service_id" yaml:"serviceId" gorm:"default:1"`
	IsRequiredSubject  bool `json:"is_required_subject" yaml:"isRequiredSubject" gorm:"default:1"`
	IsRequiredEmail    bool `json:"is_required_email" yaml:"isRequiredEmail" gorm:"default:1"`
	IsRequiredCategory bool `json:"is_required_category" yaml:"isRequiredCategory" gorm:"default:1"`
	IsRequiredText     bool `json:"is_required_text" yaml:"isRequiredText" gorm:"default:1"`
}
