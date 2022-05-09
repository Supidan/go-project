package entity

import (
	"time"

	"gorm.io/gorm"
)

type Seller struct {
	ID            uint `gorm:"primaryKey"`
	Email         string
	Name          string
	Password      string
	PickupAddress string    `gorm:"type=text"`
	Orders        []Order   `gorm:"foreignKey=OrderID"`
	Products      []Product `gorm:"foreignKey=SellerID"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
