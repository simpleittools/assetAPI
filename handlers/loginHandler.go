package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"os"
	"time"
)

// Secret Key for JWT signing
// todo: replace the secret with a real secert
var secret = []byte(os.Getenv("JWTSECRET"))

type Creds struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(c *fiber.Ctx) error {
	var creds Creds

	err := c.BodyParser(&creds)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid credentials"})
	}

	// todo: replace with actual user data
	if creds.Username != "admin" || creds.Password != "password" {
		return c.Status(401).JSON(fiber.Map{"message": "Invalid credentials"})
	}

	// Create the JWT token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set the claims and expiration time
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = creds.Username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// sign the token
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error creating token"})
	}

	// make the JWT a cookie
	cookie := fiber.Cookie{
		// todo: give the cookie a real name
		Name:     os.Getenv("SITENAME"),
		Value:    tokenString,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{"message": "success"})
}
