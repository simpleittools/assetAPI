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
	{
		Email:     "joe@joe.com",
		FirstName: "Joe",
		LastName:  "Logan",
		Username:  "joe",
		Password:  []byte("password"),
	},
}
