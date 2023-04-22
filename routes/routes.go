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

	app.Post("/login", handlers.LoginHandler)
	app.Post("/register", handlers.Register)

	app.Post("/contacts/create", handlers.ContactCreate)
	app.Post("/devices/create", handlers.DeviceCreate)
	app.Post("/devices/devicetypes/create", handlers.DeviceTypeCreate)
}
