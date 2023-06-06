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
	device.Patch("/:slug", handlers.DeviceUpdate)
	device.Delete("/:slug", handlers.DeviceSoftDelete)
	// TODO: limit access to the DeviceHardDelete to specific users
	device.Delete("/permanent/:slug", handlers.DeviceHardDelete)

	deviceType := app.Group("api/devicetypes")
	deviceType.Get("/", handlers.DeviceTypeIndex)
	deviceType.Post("/create", handlers.DeviceTypeCreate)
	deviceType.Patch("/:slug", handlers.DeviceTypeUpdate)
	// TODO: don't expose the delete endpoints yet. Want to decide if I should cascade the deletion or if I want to prevent it if items exist
	//deviceType.Post("/:slug", handlers.DeviceTypeSoftDelete)
	//deviceType.Delete("/permanent/:slug", handlers.DeviceTypeHardDelete)

	app.Post("/api/login", handlers.LoginHandler)
	app.Post("/api/register", handlers.Register)

	app.Post("/contacts/create", handlers.ContactCreate)

}
