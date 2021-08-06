package main

import (
	"github.com/bxcodec/faker/v3"
	"go-ambassador/src/database"
	"go-ambassador/src/models"
)

func main() {
	database.Connect()

	for i := 0; i < 30; i++ {
		ambassador := models.User{
			FristName:    faker.FirstName(),
			LastName:     faker.LastName(),
			Email:        faker.Email(),
			IsAmbassador: true,
		}

		ambassador.SetPassword("1234")

		database.DB.Create(&ambassador)
	}
}
