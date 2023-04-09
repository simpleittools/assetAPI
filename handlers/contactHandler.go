package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/simpleittools/assetapi/database"
	"github.com/simpleittools/assetapi/helpers"
	"github.com/simpleittools/assetapi/models"
)

func ContactCreate(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)
	if err != nil {
		return err
	}

	contact := models.Contact{
		Slug:      helpers.Slugify(data["first_name"] + data["last_name"]),
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
		CellPhone: data["cell_phone"],
		ClientID:  helpers.UintConv(data["client_id"]),
	}

	database.DB.Create(&contact)
	return c.JSON(contact)
}
