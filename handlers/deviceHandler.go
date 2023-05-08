package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/simpleittools/assetapi/database"
	"github.com/simpleittools/assetapi/helpers"
	"github.com/simpleittools/assetapi/models"
)

// DeviceIndex will list all Devices
func DeviceIndex(c *fiber.Ctx) error {
	var devices []models.Device

	err := database.DB.Preload("Client").Preload("DeviceType").Find(&devices).Error
	if err != nil {
		return err
	}

	return c.JSON(devices)
}

// DeviceCreate will create a new client
func DeviceCreate(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)
	if err != nil {
		return err
	}

	device := models.Device{
		Slug:         helpers.Slugify(data["device_name"] + "-" + data["serial_number"]),
		DeviceName:   data["device_name"],
		SerialNumber: data["serial_number"],
		Make:         data["make"],
		DeviceModel:  data["device_model"],
		ClientID:     helpers.UintConv(data["client_id"]),
		DeviceTypeID: helpers.UintConv(data["device_type_id"]),
		// todo fix the Active and IsLoaner from the static assignments
		Active:   true,
		IsLoaner: false,
	}

	database.DB.Create(&device)

	//deviceCreateSuccess := models.TransactionLog{
	//	TransactionType: "Device Created",
	//	Name:            data["device_name"],
	// todo: get the logged in user
	//UserID:          user.ID,
	//IPAddress: c.IP(),
	//}
	//database.DB.Create(&deviceCreateSuccess)
	return c.JSON(device)
}

func DeviceShow(c *fiber.Ctx) error {
	slug := c.Params("slug")
	device := models.Device{}
	err := database.DB.Find(&device, "slug", slug).Error
	if err != nil {
		return err
	}
	return c.JSON(device)
}
