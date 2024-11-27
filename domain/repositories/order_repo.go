package repositories

import (
	"url-shortener/order/data/models"
	"url-shortener/order/domain/entities"
)

// OrderRepository is...
//
//go:generate mockgen -destination=mocks/mock_order_repo.go -package=mocks url-shortener/order/domain/repositories OrderRepository
type OrderRepository interface {
	GetByID(id int64) (*models.Order, error)
	GetByFilters(filters entities.OrderFiltersInput) (*[]models.Order, error)
	Create(input entities.CreateOrderInput) (*models.Order, error)
	Update(input entities.UpdateOrderStatusInput) (*models.Order, error)
}
