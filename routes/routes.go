package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/simpleittools/assetapi/handlers"
)

func APIRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	app.Post("/login", handlers.LoginHandler)
	app.Post("/register", handlers.Register)
	app.Post("/clients/create", handlers.ClientCreate)
	app.Post("/contacts/create", handlers.ContactCreate)
	app.Post("/devices/create", handlers.DeviceCreate)
	app.Post("/devices/devicetypes/create", handlers.DeviceTypeCreate)
}
