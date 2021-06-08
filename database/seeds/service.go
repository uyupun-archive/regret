package main

import (
	"github.com/uyupun/regret/database"
	"github.com/uyupun/regret/models"
)

func SeedService() {
	service := models.Service{
		ID:          1,
		Name:        "test",
		NameJa:      "テスト",
		Description: "テストデータです",
		AccessToken: "testtest",
	}

	db, err := database.ConnectGorm()
	if err != nil {
		panic(err)
	}
	db.Create(&service)
}
