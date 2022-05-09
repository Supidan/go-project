package entity

import (
	"time"

	"gorm.io/gorm"
)

type Buyer struct {
	ID              uint `gorm:"primaryKey"`
	Email           string
	Name            string
	Password        string
	ShippingAddress string  `gorm:"type=text"`
	Orders          []Order `gorm:"foreignKey=BuyerID"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}
