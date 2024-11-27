package models

import "url-shortener/order/data/constants"

type OrderItem struct {
	ID          int64   `json:"id" gorm:"primaryKey;column:id" `
	OrderID     int64   `json:"order_id" gorm:"column:order_id" `
	Description string  `json:"description" gorm:"column:description" `
	Price       float32 `json:"price" gorm:"column:price" `
	Quantity    int64   `json:"quantity" gorm:"column:quantity" `
}

func (OrderItem) TableName() string {
	return constants.TABLE_NAME_ORDER_ITEMS
}
