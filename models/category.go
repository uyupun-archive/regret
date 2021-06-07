package models

type Category struct {
	ID        int    `json:"id" yaml:"id"`
	ServiceId int    `json:"service_id" yaml:"serviceId"`
	Name      string `json:"name" yaml:"name"`
	NameJa    string `json:"name_ja" yaml:"nameJa"`
}
