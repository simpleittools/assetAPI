package handlers

import (
	"github.com/simpleittools/assetapi/database"
	"github.com/simpleittools/assetapi/models"
	"golang.org/x/crypto/bcrypt"
)

func Seed() error {
	for _, userSeed := range database.UserSeed {
		hashedPassword, err := bcrypt.GenerateFromPassword(userSeed.Password, 12)
		if err != nil {
			return err
		}

		user := models.User{
			Email:     userSeed.Email,
			FirstName: userSeed.FirstName,
			LastName:  userSeed.LastName,
			Username:  userSeed.Username,
			Password:  hashedPassword,
		}
		if err = database.DB.Create(&user).Error; err != nil {
			return err
		}
	}
	return nil
}
