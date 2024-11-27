package repositories

import (
	"url-shortener/domain/repositories"
)

// The NewUrlShortenerRepository function is a factory function that returns a new instance of the urlShortenerRepo
func NewUrlShortenerRepository() repositories.UrlShortenerRepository {
	return &urlShortenerRepo{}
}

type urlShortenerRepo struct {
}

// // GetByFilters fetches orders with the provided filters.
// func (or *orderRepo) GetByFilters(filters entities.OrderFiltersInput) (*[]models.Order, error) {
// 	var orders []models.Order

// 	db := or.db.Table(constants.TABLE_NAME_ORDERS)
// 	db = or.setFilters(filters, db)

// 	sortOrderBy, err := or.validateAndReturnSortQuery(filters)
// 	if err != nil {
// 		return nil, err
// 	}

// 	err = db.Preload("OrderItem").Joins("inner join order_items oi on oi.order_id = orders.id ").Limit(filters.Limit).Offset(filters.Offset).Order(sortOrderBy).Find(&orders).Debug().Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	if len(orders) == 0 {
// 		return nil, ErrorOrdersNotFound
// 	}

// 	return &orders, nil
// }

// func (or *orderRepo) setFilters(filters entities.OrderFiltersInput, db *gorm.DB) *gorm.DB {
// 	if filters.OrderID != 0 {
// 		db = db.Where("orders.id = ?", filters.OrderID)
// 	}

// 	if filters.CurrencyUnit != "" {
// 		db = db.Where("orders.currency_unit like ?", filters.CurrencyUnit)
// 	}

// 	if filters.Total != 0 {
// 		db = db.Where("orders.total = ?", filters.Total)
// 	}

// 	if filters.Status != "" {
// 		db = db.Where("orders.status = ?", filters.Status)
// 	}

// 	if !filters.Item.IsEmpty() {
// 		db = or.setOrderItemFilters(filters.Item, db)
// 	}

// 	return db
// }

// func (or *orderRepo) setOrderItemFilters(orderItemFilter entities.OrderItemFiltersInput, db *gorm.DB) *gorm.DB {
// 	if orderItemFilter.Price != 0 {
// 		db = db.Where("oi.price = ?", orderItemFilter.Price)
// 	}

// 	if orderItemFilter.Quantity != 0 {
// 		db = db.Where("oi.quantity = ?", orderItemFilter.Quantity)
// 	}

// 	if orderItemFilter.Description != "" {
// 		db = db.Where("oi.description like ?", "%"+orderItemFilter.Description+"%")
// 	}

// 	return db
// }

// func (or *orderRepo) validateAndReturnSortQuery(filters entities.OrderFiltersInput) (string, error) {
// 	order := filters.GetOrder()
// 	if order != "desc" && order != "asc" {
// 		return "", ErrorOrderByInvalid

// 	}

// 	if !filters.VaidateSortBy() && filters.SortBy != "" {
// 		return "", ErrorInvalidSortBy
// 	}

// 	sortBy := filters.GetSortBy()

// 	return fmt.Sprintf("%s %s", sortBy, strings.ToUpper(order)), nil
// }

// // GetByID fetches order with the provided id.
// func (or *orderRepo) GetByID(id int64) (*models.Order, error) {
// 	var order models.Order

// 	db := or.db.Table(constants.TABLE_NAME_ORDERS)
// 	err := db.Preload("OrderItem").Where("orders.id = ?", id).Find(&order).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	if order.ID == 0 {
// 		return nil, ErrorOrderNotFound
// 	}

// 	return &order, nil
// }

// // Create adds an order to the database.
// func (or *orderRepo) Create(input entities.CreateOrderInput) (*models.Order, error) {
// 	// create a new order from input parameters
// 	order := models.Order{
// 		Total:        input.Total,
// 		Status:       input.Status,
// 		CurrencyUnit: input.CurrencyUnit,
// 	}

// 	// insert order into the database
// 	orderResult := or.db.Create(&order)
// 	if orderResult.Error != nil || orderResult.RowsAffected == 0 {
// 		return nil, orderResult.Error
// 	}

// 	var orderItems []models.OrderItem
// 	for _, item := range input.OrderItem {
// 		// creating each orderItem from input parameters
// 		orderItem := models.OrderItem{
// 			OrderID:     order.ID,
// 			Price:       item.Price,
// 			Quantity:    item.Quantity,
// 			Description: item.Description,
// 		}

// 		// insert the orderItem into the database
// 		orderItemResult := or.db.Create(&orderItem)
// 		if orderItemResult.Error != nil || orderItemResult.RowsAffected == 0 {
// 			return nil, orderItemResult.Error
// 		}

// 		orderItems = append(orderItems, orderItem)
// 	}

// 	// assign the orderItems slice to the OrderItem field of the order
// 	order.OrderItem = orderItems

// 	return &order, nil
// }

// // Update changes order statuses as of now, but can be extended to update an order!
// func (or *orderRepo) Update(input entities.UpdateOrderStatusInput) (*models.Order, error) {
// 	result := or.db.Model(&models.Order{}).Where("id = ?", input.OrderID).Updates(models.Order{Status: input.Status})
// 	if result.RowsAffected == 0 {
// 		return nil, ErrorOrderStatusAlreadyUpdated
// 	}

// 	if result.Error != nil {
// 		return nil, result.Error
// 	}

// 	order, err := or.GetByID(input.OrderID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return order, nil
// }
