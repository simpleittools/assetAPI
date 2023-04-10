package models

import "gorm.io/gorm"

type PhoneNumber struct {
	gorm.Model
	Number    string `json:"number" gorm:"unique"`
	ClientID  uint   `json:"client_id"`
	ContactID uint   `json:"contact_id"`
}
