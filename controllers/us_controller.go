package controllers

import (
	"url-shortener/app/controller"
	"url-shortener/url-shortener/domain/services"
)

// UrlShortenerController is...
type UrlShortenerController interface {
	// GetOrder(gCtx *gin.Context)
	// GetOrders(gCtx *gin.Context)
	// CreateOrder(gCtx *gin.Context)
	// UpdateOrderStatus(gCtx *gin.Context)
}

// NewUrlShortenerController creates a new instance of the urlShortenerController
func NewUrlShortenerController(orderService services.OrderService) UrlShortenerController {
	return &urlShortenerController{orderService: orderService}
}

type urlShortenerController struct {
	controller.Controller
	orderService services.OrderService
}

// // GetOrders handles incoming requests to retrieve orders by given filters.
// func (c *orderController) GetOrders(gCtx *gin.Context) {
// 	var input entities.OrderFiltersInput
// 	if err := gCtx.BindJSON(&input); err != nil {
// 		gCtx.JSON(http.StatusBadRequest, gin.H{"code": "failed", "msg": err.Error()})
// 		return
// 	}

// 	orders, err := c.orderService.GetOrdersByFilters(input)
// 	if err != nil {
// 		c.SendWithError(gCtx, err)
// 		return
// 	}

// 	c.Send(gCtx, orders)
// }

// // GetOrder handles incoming requests to retrieve an order by its ID.
// func (c *orderController) GetOrder(gCtx *gin.Context) {
// 	input := entities.GetOrderByIDInput{}

// 	if err := gCtx.BindUri(&input); err != nil {
// 		gCtx.JSON(http.StatusBadRequest, gin.H{"code": "failed", "msg": err.Error()})
// 		return
// 	}

// 	order, err := c.orderService.GetOrderByID(input.OrderID)
// 	if err != nil {
// 		c.SendWithError(gCtx, err)
// 		return
// 	}

// 	c.Send(gCtx, order)
// }

// // AddToCart adds an item in cart based on inputs and returns updated order or error
// func (c *orderController) CreateOrder(gCtx *gin.Context) {
// 	var input entities.CreateOrderInput

// 	if err := gCtx.ShouldBind(&input); err != nil {
// 		gCtx.JSON(http.StatusBadRequest, gin.H{"code": "failed", "msg": err.Error()})
// 		return
// 	}

// 	order, err := c.orderService.CreateOrder(input)
// 	if err != nil {
// 		c.SendWithError(gCtx, err)
// 		return
// 	}

// 	c.Send(gCtx, order)
// }

// // UpdateOrderStatus updates an order's status based on inputs and returns updated order or error
// func (c *orderController) UpdateOrderStatus(gCtx *gin.Context) {
// 	var input entities.UpdateOrderStatusInput
// 	if err := gCtx.ShouldBind(&input); err != nil {
// 		c.SendWithError(gCtx, err)
// 		return
// 	}

// 	order, err := c.orderService.UpdateOrderStatus(input)
// 	if err != nil {
// 		c.SendWithError(gCtx, err)
// 		return
// 	}

// 	c.Send(gCtx, order)
// }
