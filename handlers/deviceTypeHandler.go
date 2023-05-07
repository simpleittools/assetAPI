package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/simpleittools/assetapi/database"
	"github.com/simpleittools/assetapi/helpers"
	"github.com/simpleittools/assetapi/models"
)

func DeviceTypeCreate(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)
	if err != nil {
		return err
	}

	deviceType := models.DeviceType{
		Slug:       helpers.Slugify(data["device_type"]),
		DeviceType: data["device_type"],
	}

	database.DB.Create(&deviceType)

	//deviceTypeCreateSuccess := models.TransactionLog{
	//TransactionType: "Device Type Created",
	//Name:            data["device_type"],
	// todo: get the logged in user
	//UserID:          user.ID,
	//IPAddress: c.IP(),
	//}
	//database.DB.Create(&deviceTypeCreateSuccess)
	return c.JSON(deviceType)
}
