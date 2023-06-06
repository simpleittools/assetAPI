package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/simpleittools/assetapi/database"
	"github.com/simpleittools/assetapi/helpers"
	"github.com/simpleittools/assetapi/models"
)

// DeviceTypeIndex will show all device Types
func DeviceTypeIndex(c *fiber.Ctx) error {
	var data []models.DeviceType

	database.DB.Find(&data)
	//database.DB.Where("active = true").Find(&clients)
	//database.DB.Preload(clause.Associations).Where("is_active = 1").Find(&clients)

	return c.JSON(data)
}

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

// DeviceTypeUpdate will PATCH the device type details after edited by the user
func DeviceTypeUpdate(c *fiber.Ctx) error {
	slug := c.Params("slug")

	var data models.DeviceType

	err := c.BodyParser(&data)
	if err != nil {
		return err
	}

	device := &models.DeviceType{
		DeviceType: data.DeviceType,
	}

	err = database.DB.Model(&data).Where("slug = ?", slug).Updates(&device).Error
	if err != nil {
		return err
	}

	return c.JSON(device)
}

// DeviceTypeSoftDelete will set the deleted_at entry in the database. This will prevent the database from returning these items on a default look-up
func DeviceTypeSoftDelete(c *fiber.Ctx) error {
	slug := c.Params("slug")

	var data models.DeviceType

	database.DB.Model(&data).Where("slug = ?", slug).Delete(&slug)

	return c.JSON(slug)
}

// DeviceHardDelete will permanently delete entries
func DeviceTypeHardDelete(c *fiber.Ctx) error {
	slug := c.Params("slug")

	var data models.DeviceType

	database.DB.Unscoped().Model(&data).Where("slug = ?", slug).Delete(&slug)

	return c.JSON(slug)
}
