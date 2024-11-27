package repositories

// UrlShortenerRepository is...
//
//go:generate mockgen -destination=mocks/mock_us_repo.go -package=mocks url-shortener/domain/repositories UrlShortenerRepository
type UrlShortenerRepository interface {
	// GetByID(id int64) (*models.Order, error)
	// GetByFilters(filters entities.OrderFiltersInput) (*[]models.Order, error)
	// Create(input entities.CreateOrderInput) (*models.Order, error)
	// Update(input entities.UpdateOrderStatusInput) (*models.Order, error)
}
