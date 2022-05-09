package entity

import "time"

type OrderDetail struct {
	ID        uint `gorm:"primaryKey"`
	OrderID   uint
	ProductID uint
	Quantity  uint16
	Order     Order
	Product   Product
	CreatedAt time.Time
	UpdatedAt time.Time
}
