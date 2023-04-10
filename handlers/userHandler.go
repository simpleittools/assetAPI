package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/simpleittools/assetapi/database"
	"github.com/simpleittools/assetapi/models"
	"golang.org/x/crypto/bcrypt"
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

// LoginHandler will log a user in and create a transaction log entry
func LoginHandler(c *fiber.Ctx) error {
	var creds map[string]string
	err := c.BodyParser(&creds)
	if err != nil {
		return err
	}

	// todo: create a transaction log envelope to store this information
	loginFail := models.TransactionLog{
		TransactionType: "Login Failed",
		Username:        creds["username"],
		IPAddress:       c.IP(),
	}

	var user models.User
	// Search the Database for the user by email
	database.DB.Where("username = ?", creds["username"]).First(&user)

	// todo: create a transaction log envelope to store this information
	loginSuccess := models.TransactionLog{
		TransactionType: "Login Success",
		UserID:          user.ID,
		IPAddress:       c.IP(),
	}

	// todo: change the message
	if user.ID == 0 {
		c.Status(404)
		database.DB.Create(&loginFail)
		return c.JSON(fiber.Map{
			"message": "User not found!",
		})
	}
	// todo: change the message
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(creds["password"])); err != nil {
		c.Status(400)
		database.DB.Create(&loginFail)
		return c.JSON(fiber.Map{
			"message": "Incorrect Password",
		})
	}

	database.DB.Create(&loginSuccess)

	// Create the JWT token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set the claims and expiration time
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = creds["username"]
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// sign the token
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error creating token"})
	}

	// make the JWT a cookie
	cookie := fiber.Cookie{
		Name:     os.Getenv("SITENAME"),
		Value:    tokenString,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{"message": "success"})
}

// Register will register a user
func Register(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)
	if err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Passwords do not match",
		})
	}

	// hash the password
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 12)

	user := models.User{
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Username:  data["username"],
		Email:     data["email"],
		Password:  hashPassword,
	}

	database.DB.Create(&user)

	// todo: the UserID is collecting the newly created userID. Need to set this to get the logged in userID
	registerSuccess := models.TransactionLog{
		TransactionType: "User Created",
		UserID:          user.ID,
		Username:        data["username"],
		IPAddress:       c.IP(),
	}

	database.DB.Create(&registerSuccess)

	return c.JSON(user)
}
