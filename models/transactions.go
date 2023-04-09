package models

import "time"

type TransactionLog struct {
	ID              uint      `json:"id" gorm:"primaryKey"`
	UserID          uint      `json:"user_id" gorm:"null"`
	Username        string    `json:"username" gorm:"null"`
	TransactionType string    `json:"transaction_type" gorm:"not null"`
	IPAddress       string    `json:"ip_address"`
	TransactionTime time.Time `json:"transaction_time" gorm:"autoUpdateTime"`
}
