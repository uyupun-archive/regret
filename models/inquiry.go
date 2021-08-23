package models

type Inquiry struct {
	Subject    string `json:"subject" yaml:"subject"`
	Email      string `json:"email" yaml:"email"`
	CategoryId int    `json:"category_id" yaml:"categoryId"`
	Text       string `json:"text" yaml:"text"`
}
