package models

import "gorm.io/gorm"

type Contacts struct {
	gorm.Model
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email" gorm:"unique"`
	// todo: add foreign items
	// client
	// devices -m2m
	// desk_phone
	CellPhone string `json:"cell_phone" gorm:"null"`
}
