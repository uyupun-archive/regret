package models

type Service struct {
	ID          int    `json:"id" yaml:"id"`
	Name        string `json:"name" yaml:"name"`
	NameJa      string `json:"name_ja" yaml:"nameJa"`
	Description string `json:"description" yaml:"description"`
	AccessToken string `json:"access_token" yaml:"accessToken"`
}

type AddService struct {
	Name        string `json:"name" yaml:"name"`
	NameJa      string `json:"name_ja" yaml:"nameJa"`
	Description string `json:"description" yaml:"description"`
}
