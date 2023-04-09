package database

import "github.com/simpleittools/assetapi/models"

var UserSeed = []models.User{
	{
		Email:     "ryan@ryan.com",
		FirstName: "Ryan",
		LastName:  "Mooney",
		Username:  "ryan",
		Password:  []byte("password"),
	},
}
