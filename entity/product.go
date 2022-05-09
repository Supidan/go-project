package entity

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint `gorm:"primaryKey"`
	SellerID    uint
	Name        string
	Description string `gorm:"type=text"`
	Price       uint64
	Seller      Seller
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
