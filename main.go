package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/simpleittools/assetapi/database"
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
	//seedDb := true
	//if seedDb == true {
	//	handlers.Seed(),
	//}

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.APIRoutes(app)

	app.Listen(PORT)
}
