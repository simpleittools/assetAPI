package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/simpleittools/assetapi/database"
	"github.com/simpleittools/assetapi/helpers"
	"github.com/simpleittools/assetapi/models"
	"gorm.io/gorm/clause"
)

// ClientIndex will show all registered clients
func ClientIndex(c *fiber.Ctx) error {
	var clients []models.Client

	err := database.DB.Find(&clients).Error

	if err != nil {
		return err
	}
	//database.DB.Where("active = true").Find(&clients)
	//database.DB.Preload(clause.Associations).Where("is_active = 1").Find(&clients)

	return c.JSON(clients)
}

// ClientCreate will create a new client
func ClientCreate(c *fiber.Ctx) error {
	//var data map[interface{}]string
	//
	//err := c.BodyParser(&data)
	var data models.Client

	err := c.BodyParser(&data)
	if err != nil {
		return err
	}

	slug := helpers.Slugify(data.ClientName)

	client := models.Client{
		Slug:           slug,
		ClientName:     data.ClientName,
		Address:        data.Address,
		Address2:       data.Address2,
		Phone:          data.Phone,
		PrimaryEmail:   data.PrimaryEmail,
		SecondaryEmail: data.SecondaryEmail,
		ClientActive:   data.ClientActive,
	}

	database.DB.Create(&client)
	return c.JSON(client)
}

// ClientShow will return the results of a selected client
func ClientShow(c *fiber.Ctx) error {
	slug := c.Params("slug")
	client := models.Client{}
	err := database.DB.Preload(clause.Associations).Where("slug =?", slug).Find(&client).Error

	if err != nil {
		return err
	}
	return c.JSON(client)
}

// ClientUpdate will PATCH the client details after edited by the user
func ClientUpdate(c *fiber.Ctx) error {
	slug := c.Params("slug")

	var data models.Client

	err := c.BodyParser(&data)
	if err != nil {
		return err
	}

	client := &models.Client{
		ClientName:     data.ClientName,
		Address:        data.Address,
		Address2:       data.Address2,
		Phone:          data.Phone,
		PrimaryEmail:   data.PrimaryEmail,
		SecondaryEmail: data.SecondaryEmail,
		ClientActive:   data.ClientActive,
	}

	err = database.DB.Model(&data).Where("slug = ?", slug).Updates(&client).Error
	if err != nil {
		return err
	}

	return c.JSON(client)
}

// ClientSoftDelete will set the deleted_at entry in the database. This will prevent the database from returning these items on a default look-up
func ClientSoftDelete(c *fiber.Ctx) error {
	slug := c.Params("slug")

	var data models.Client

	database.DB.Model(&data).Where("slug = ?", slug).Delete(&slug)

	return c.JSON(slug)
}

// ClientHardDelete will permanently delete entries
func ClientHardDelete(c *fiber.Ctx) error {
	slug := c.Params("slug")

	var data models.Client

	database.DB.Unscoped().Model(&data).Where("slug = ?", slug).Delete(&slug)

	return c.JSON(slug)
}
