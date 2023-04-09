package database

import (
	"fmt"
	"github.com/simpleittools/assetapi/models"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func Conn() {
	dbType := os.Getenv("DB_ENGINE")
	dsn := ""
	switch dbType {
	case "POSTGRES":
		dsn = os.Getenv("POSTGRESDSN")
		conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("Could not connect to the POSTGRES DB")
		} else {
			fmt.Println("connected to POSTGRES DB")
		}
		DB = conn
		conn.AutoMigrate(&models.User{})
	case "MYSQL":
		dsn = "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
		conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("Could not connect to the MYSQL DB")
		} else {
			fmt.Println("connected to MYSQL")
		}
		DB = conn
		conn.AutoMigrate(&models.User{})
	case "SQLITE":
		dbName := os.Getenv("SQLITEDBNAME")
		deleteDB := os.Getenv("DELETEDB")
		if deleteDB == "TRUE" {
			err := os.Remove(dbName)
			if err != nil {
				fmt.Println(err)
				return
			}
		}

		folderPath := "./db-data"
		_, err := os.Stat(folderPath)
		if os.IsNotExist(err) {
			err := os.MkdirAll(folderPath, os.ModePerm)
			if err != nil {
				fmt.Printf("Error creating folder: %v\n", err)
			} else {
				fmt.Println("Database Folder Created Successfully")
			}
		} else if err != nil {
			fmt.Printf("Error checking if Database folder exists: %v\n", err)
		} else {
			fmt.Println("Database folder exists")
		}
		conn, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
		if err != nil {
			log.Fatal("Could not connect to the SQLITE DB")
		} else {
			fmt.Println("connected to SQLITE")
		}
		DB = conn
		conn.AutoMigrate(&models.User{}, &models.TransactionLog{})
	default:
		panic("invalid database definition")
	}

}
