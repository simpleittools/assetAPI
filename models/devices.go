package models

import "gorm.io/gorm"

type Device struct {
	gorm.Model
	Slug         string     `json:"slug" gorm:"unique,notnull"`
	DeviceName   string     `json:"device_name"`
	SerialNumber string     `json:"serial_number" gorm:"unique"`
	Make         string     `json:"make,omitempty" gorm:"null"`
	DeviceModel  string     `json:"device_model,omitempty" gorm:"null"`
	ClientID     uint       `json:"client_id"`
	Client       Client     `json:"client"`
	DeviceTypeID uint       `json:"device_type_id"`
	DeviceType   DeviceType `json:"device_type"`
	Active       *bool      `json:"active" gorm:"default:true"`
	IsLoaner     *bool      `json:"is_loaner" gorm:"default:false"`
}
