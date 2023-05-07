package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/simpleittools/assetapi/database"
	"github.com/simpleittools/assetapi/handlers"
	"github.com/simpleittools/assetapi/routes"
	"log"
	"os"
)

func main() {
	// Load the environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("cannot find .env file. Please create the .env file")
	}
	status := os.Getenv("STATUS")

	if status == "DEVELOPMENT" {
		fmt.Println("We are in DEV")
	} else if status == "PRODUCTION" {
		fmt.Println("We are in PROD")
	} else {
		log.Panicln("Invalid Status configuration")
	}

	PORT := os.Getenv("APIPORT")

	app := fiber.New()

	database.Conn()
	seedDb := os.Getenv("SEEDDB")
	switch seedDb {
	case "TRUE":
		handlers.Seed()
	case "FAlSE":
		fmt.Println("You are not seeding the DB")
	default:
		log.Fatal("unable to determine seeding")

	}

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.APIRoutes(app)

	err = app.Listen(PORT)
	if err != nil {
		log.Fatal(err)
	}
}
