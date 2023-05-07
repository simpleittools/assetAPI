package database

import "github.com/simpleittools/assetapi/models"

var UserSeed = []models.User{
	{
		Email:     "ryan@ryan.com",
		FirstName: "Ryan",
		LastName:  "Mooney",
		Username:  "rya",
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

var ClientSeed = []models.Client{
	{
		Slug:           "first-client",
		ClientName:     "First Client",
		Address:        "123 Fake Street",
		Address2:       "unit 531",
		Phone:          "555-555-5555",
		PrimaryEmail:   "email@fistclient.com",
		SecondaryEmail: "alternate@firstclient.com",
	},
}
