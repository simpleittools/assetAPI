package models

import "time"

type TransactionLog struct {
	ID              uint      `json:"id" gorm:"primaryKey"`
	UserID          uint      `json:"user_id" gorm:"not null"`
	TransactionType string    `json:"transaction_type" gorm:"not null"`
	TransactionTime time.Time `json:"transaction_time" gorm:"autoUpdateTime"`
}
