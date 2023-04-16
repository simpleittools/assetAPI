package models

import "gorm.io/gorm"

type DeviceType struct {
	gorm.Model
	Slug       string `json:"slug" gorm:"unique"`
	DeviceType string `json:"device_type" gorm:"unique"`
	Devices    []Device
}
