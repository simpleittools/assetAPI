package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/simpleittools/assetapi/handlers"
)

func APIRoutes(app *fiber.App) {
	app.Get("/api/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	app.Post("/login", handlers.LoginHandler)
	app.Post("/register", handlers.Register)
	app.Post("/api/clients/create", handlers.ClientCreate)
	app.Get("/api/clients", handlers.ClientIndex)
	app.Post("/contacts/create", handlers.ContactCreate)
	app.Post("/devices/create", handlers.DeviceCreate)
	app.Post("/devices/devicetypes/create", handlers.DeviceTypeCreate)
}
