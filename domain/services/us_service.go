package services

import (
	"url-shortener/domain/repositories"
)

// UrlShortenerService is...
type UrlShortenerService interface {
	// GetOrderByID(id int64) (*models.Order, error)
	// GetOrdersByFilters(filters entities.OrderFiltersInput) (*[]models.Order, error)
	// CreateOrder(input entities.CreateOrderInput) (*models.Order, error)
	// UpdateOrderStatus(input entities.UpdateOrderStatusInput) (*models.Order, error)
}

// The NewUrlShortenerService function is a factory function that returns a new instance of the urlShortenerService
func NewUrlShortenerService(repo repositories.UrlShortenerRepository) UrlShortenerService {
	return &urlShortenerService{repo: repo}
}

type urlShortenerService struct {
	repo repositories.UrlShortenerRepository
}

// // The GetOrdersByFilters method of the orderService struct utilizes the OrderRepository instance, by calling the GetByFilters method on it, and returns all orders from the given filters and errors if received.
// func (service *orderService) GetOrdersByFilters(filters entities.OrderFiltersInput) (*[]models.Order, error) {
// 	return service.repo.GetByFilters(filters)
// }

// // The GetOrderByID method of the orderService struct utilizes the OrderRepository instance, by calling the GetByID method on it, and returns any order and errors if received.
// func (service *orderService) GetOrderByID(id int64) (*models.Order, error) {
// 	return service.repo.GetByID(id)
// }

// // The CreateOrder method of the orderService struct utilizes the OrderRepository instance, by calling the Create method on it, and returns the order created and errors if received.
// func (service *orderService) CreateOrder(input entities.CreateOrderInput) (*models.Order, error) {
// 	return service.repo.Create(input)
// }

// // The UpdateOrderStatus method of the orderService struct utilizes the OrderRepository instance, by calling the Update method on it, and returns the order with updated status and errors if received.
// func (service *orderService) UpdateOrderStatus(input entities.UpdateOrderStatusInput) (*models.Order, error) {
// 	return service.repo.Update(input)
// }
