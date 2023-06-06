package models

import "gorm.io/gorm"

type DevicePropertyType struct {
	gorm.Model
	Slug               string       `json:"slug" gorm:"unique"`
	DevicePropertyType string       `json:"device_property_type" gorm:"unique"`
	DeviceType         []DeviceType `gorm:"many2many:device_property;"`
}
