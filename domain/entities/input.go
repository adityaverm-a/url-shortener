package entities

import (
	"reflect"
	"url-shortener/order/data/constants"
	"url-shortener/order/data/models"
)

type GetOrderByIDInput struct {
	OrderID int64 `json:"order_id" uri:"id" binding:"required"`
}

type CreateOrderInput struct {
	Status       string                 `form:"status" json:"status" binding:"required"`
	Total        float32                `form:"total" json:"total" binding:"required"`
	OrderItem    []CreateOrderItemInput `form:"items" json:"items" binding:"required"`
	CurrencyUnit string                 `form:"currency_unit" json:"currency_unit" binding:"required"`
}

type CreateOrderItemInput struct {
	Price       float32 `form:"price" json:"price" binding:"required"`
	Quantity    int64   `form:"quantity" json:"quantity" binding:"required"`
	Description string  `form:"description" json:"description" binding:"required"`
}

type UpdateOrderStatusInput struct {
	OrderID int64  `form:"id" json:"id" binding:"required"`
	Status  string `form:"status" json:"status" binding:"required"`
}

type OrderFiltersInput struct {
	OrderID      int                   `form:"id" json:"id"`
	Status       string                `form:"status" json:"status"`
	Total        float32               `form:"total" json:"total"`
	CurrencyUnit string                `form:"currency_unit" json:"currency_unit"`
	Item         OrderItemFiltersInput `form:"item" json:"item"`
	Limit        int                   `form:"limit" json:"limit"`
	Offset       int                   `form:"offset" json:"offset"`
	SortBy       string                `form:"sort_by" json:"sort_by"`
	Order        string                `form:"order" json:"order"`
}

func (input OrderFiltersInput) GetSortBy() string {
	if input.SortBy != "" {
		return input.SortBy
	}

	return constants.DEFAULT_SORT_BY
}

func (input OrderFiltersInput) GetOrder() string {
	if input.Order != "" {
		return input.Order
	}

	return constants.DEFAULT_ORDER
}

func (input OrderFiltersInput) VaidateSortBy() bool {
	field := input.getOrderFields()

	isSortByValid := stringInSlice(field, input.SortBy)

	return isSortByValid
}

func stringInSlice(strSlice []string, s string) bool {
	for _, v := range strSlice {
		if v == s {
			return true
		}
	}

	return false
}

func (input OrderFiltersInput) getOrderFields() []string {
	var field []string

	v := reflect.ValueOf(models.Order{})

	for i := 0; i < v.Type().NumField(); i++ {
		field = append(field, v.Type().Field(i).Tag.Get("json"))
	}

	return field
}

type OrderItemFiltersInput struct {
	Price       float64 `form:"price" json:"price"`
	Quantity    int     `form:"quantity" json:"quantity"`
	Description string  `form:"description" json:"description"`
}

func (input OrderItemFiltersInput) IsEmpty() bool {
	return input.Price == 0 && input.Quantity == 0 && input.Description == ""
}
