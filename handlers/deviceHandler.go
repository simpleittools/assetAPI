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
	var data models.Device

	err := c.BodyParser(&data)
	if err != nil {
		return err
	}

	device := models.Device{
		Slug:         helpers.Slugify(data.DeviceName + "-" + data.SerialNumber),
		DeviceName:   data.DeviceName,
		SerialNumber: data.SerialNumber,
		Make:         data.Make,
		DeviceModel:  data.DeviceModel,
		ClientID:     data.ClientID,
		DeviceTypeID: data.DeviceTypeID,
		Active:       data.Active,
		IsLoaner:     data.IsLoaner,
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
	err := database.DB.Find(&device, "slug", slug).Joins("Clients").Error
	if err != nil {
		return err
	}
	return c.JSON(device)
}

// DeviceUpdate will PATCH the device details after edited by the user
func DeviceUpdate(c *fiber.Ctx) error {
	slug := c.Params("slug")

	var data models.Device

	err := c.BodyParser(&data)
	if err != nil {
		return err
	}

	device := &models.Device{
		DeviceName:   data.DeviceName,
		SerialNumber: data.SerialNumber,
		Make:         data.Make,
		DeviceModel:  data.DeviceModel,
		ClientID:     data.ClientID,
		DeviceTypeID: data.DeviceTypeID,
		Active:       data.Active,
		IsLoaner:     data.IsLoaner,
	}

	err = database.DB.Model(&data).Where("slug = ?", slug).Updates(&device).Error
	if err != nil {
		return err
	}

	return c.JSON(device)
}

// DeviceSoftDelete will set the deleted_at entry in the database. This will prevent the database from returning these items on a default look-up
func DeviceSoftDelete(c *fiber.Ctx) error {
	slug := c.Params("slug")

	var data models.Device

	database.DB.Model(&data).Where("slug = ?", slug).Delete(&slug)

	return c.JSON(slug)
}

// DeviceHardDelete will permanently delete entries
func DeviceHardDelete(c *fiber.Ctx) error {
	slug := c.Params("slug")

	var data models.Device

	database.DB.Unscoped().Model(&data).Where("slug = ?", slug).Delete(&slug)

	return c.JSON(slug)
}
