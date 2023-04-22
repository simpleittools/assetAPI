package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/simpleittools/assetapi/database"
	"github.com/simpleittools/assetapi/helpers"
	"github.com/simpleittools/assetapi/models"
)

// ClientIndex will show all registered clients
func ClientIndex(c *fiber.Ctx) error {
	var clients []models.Client

	database.DB.Find(&clients)
	//database.DB.Where("active = true").Find(&clients)
	//database.DB.Preload(clause.Associations).Where("is_active = 1").Find(&clients)

	return c.JSON(clients)
}

// ClientCreate will create a new client
func ClientCreate(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)
	if err != nil {
		return err
	}

	client := models.Client{
		Slug:           helpers.Slugify(data["client_name"]),
		ClientName:     data["client_name"],
		Address:        data["address"],
		Address2:       data["address_2"],
		PrimaryEmail:   data["primary_email"],
		SecondaryEmail: data["secondary_email"],
	}

	database.DB.Create(&client)
	return c.JSON(client)
}
