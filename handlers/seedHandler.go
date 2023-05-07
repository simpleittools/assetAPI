package handlers

import (
	"github.com/go-faker/faker/v4"
	"github.com/simpleittools/assetapi/database"
	"github.com/simpleittools/assetapi/models"
)

func Seed() error {
	iteration := 10

	for i := 0; i < iteration; i++ {

		user := models.User{
			Email:     faker.Email(),
			FirstName: faker.FirstName(),
			LastName:  faker.LastName(),
			Username:  faker.Username(),
			Password:  []byte("password"),
		}

		if err := database.DB.Create(&user).Error; err != nil {
			return err
		}
	}

	for i := 0; i < iteration; i++ {
		client := models.Client{
			Slug:           faker.Username(),
			ClientName:     faker.Name(),
			Address:        "123 Fake Street",
			Address2:       "unit 531",
			Phone:          faker.Phonenumber(),
			PrimaryEmail:   faker.Email(),
			SecondaryEmail: faker.Email(),
		}
		if err := database.DB.Create(&client).Error; err != nil {
			return err
		}

	}
	return nil
}
