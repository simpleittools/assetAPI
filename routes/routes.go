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
	app.Post("/client/create", handlers.ClientCreate)
	app.Post("/contact/create", handlers.ContactCreate)
}
