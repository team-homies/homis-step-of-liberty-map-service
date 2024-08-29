package entity

import "gorm.io/gorm"

type Map struct {
	gorm.Model
	EventId     uint    `gorm:"column:event_id;not null"`
	Name        string  `gorm:"column:name;not null"`
	Address     string  `gorm:"column:address;not null"`
	AddressRoad string  `gorm:"column:address_road;not null"`
	Longitude   float64 `gorm:"column:longitude;not null"`
	Latitude    float64 `gorm:"column:latitude;not null"`
}

func (Map) TableName() string {
	return "map"
}
