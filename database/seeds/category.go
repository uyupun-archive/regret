package main

import (
	"fmt"

	"github.com/uyupun/regret/database"
	"github.com/uyupun/regret/models"
)

func SeedCategory() {
	category := models.Category{
		ID:        1,
		ServiceId: 1,
		Name:      "test",
		NameJa:    "test",
	}

	db, err := database.ConnectGorm()
	if err != nil {
		panic(err)
	}
	db.Create(&category)

	fmt.Println("Category seeding succeeded!")
}
