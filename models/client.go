package models

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	Slug           string `json:"slug" gorm:"unique"`
	ClientName     string `json:"client_name"`
	Address        string `json:"address,omitempty"`
	Address2       string `json:"address_2,omitempty" gorm:"null"`
	Phone          string `json:"phone" gorm:"null"`
	PrimaryEmail   string `json:"primary_email" gorm:"null"`
	SecondaryEmail string `json:"secondary_email" gorm:"null"`
	// todo: the ClientActive variable will set to true, but never reverts to false
	ClientActive bool `json:"client_active" gorm:"default:true"`
	Contacts     []Contact
	Devices      []Device
}
