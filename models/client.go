package models

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	Slug       string `json:"slug" gorm:"unique"`
	ClientName string `json:"client_name"`
	Address    string `json:"address,omitempty"`
	Address2   string `json:"address_2,omitempty" gorm:"null"`
	// todo: add foreign items
	// devices
	// phone
	// contact
	PrimaryEmail   string `json:"primary_email" gorm:"null"`
	SecondaryEmail string `json:"secondary_email" gorm:"null"`
	Active         bool   `json:"active" gorm:"default:true"`
}
