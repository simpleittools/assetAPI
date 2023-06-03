package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/simpleittools/assetapi/handlers"
)

func APIRoutes(app *fiber.App) {
	app.Get("/api/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	// Client Routes
	client := app.Group("api/clients")
	client.Get("/", handlers.ClientIndex)
	client.Get("/:slug", handlers.ClientShow)
	client.Post("/create", handlers.ClientCreate)
	client.Patch("/:slug", handlers.ClientUpdate)
	client.Delete("/:slug", handlers.ClientSoftDelete)
	// TODO: limit access to the ClientHardDelete to specific users
	client.Delete("/permanent/:slug", handlers.ClientHardDelete)

	device := app.Group("api/devices")
	device.Get("/", handlers.DeviceIndex)
	device.Get("/:slug", handlers.DeviceShow)
	device.Post("/create", handlers.DeviceCreate)

	app.Post("/api/login", handlers.LoginHandler)
	app.Post("/api/register", handlers.Register)

	app.Post("/contacts/create", handlers.ContactCreate)

	app.Post("/devices/devicetypes/create", handlers.DeviceTypeCreate)
}
