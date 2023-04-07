package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/simpleittools/assetapi/routes"
	"log"
	"os"
)

func main() {
	// Load the environment variables
	err := godotenv.Load()
	if err != nil {
		log.Panicln("cannot find .env file. Please create the .env file")
	}

	PORT := os.Getenv("APIPORT")

	app := fiber.New()
	routes.APIRoutes(app)

	app.Listen(PORT)
}
