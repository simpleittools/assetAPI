package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email     string `json:"email" gorm:"unique"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Username  string `json:"username" gorm:"unique"`
	Password  []byte `json:"-"`
	Active    bool   `json:"active" gorm:"default:true"`
	//Token     Token     `json:"token"`
	Transactions []TransactionLog
}
