package entity

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID                          uint `gorm:"primaryKey"`
	SellerID                    uint
	BuyerID                     uint
	Password                    string
	DeliverySourceAddress       string `gorm:"type=text"`
	DeliveryDestionationAddress string `gorm:"type=text"`
	TotalPrice                  uint64
	Status                      string `gorm:"size=10"`
	Seller                      Seller
	Buyer                       Buyer
	OrderDetails                []OrderDetail `gorm:"foreignKey=OrderID"`
	CreatedAt                   time.Time
	UpdatedAt                   time.Time
	DeletedAt                   gorm.DeletedAt `gorm:"index"`
}
