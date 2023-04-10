package models

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	Slug      string `json:"slug" gorm:"unique"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email" gorm:"unique"`
	// todo: add foreign items
	ClientID uint `json:"client_id"`
	// devices -m2m
	Phone []PhoneNumber
	// positional_email
	CellPhone string `json:"cell_phone" gorm:"null"`
}
