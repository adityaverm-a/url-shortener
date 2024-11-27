package services

import (
	"testing"
	"url-shortener/order/data/models"
	"url-shortener/order/domain/repositories/mocks"

	"github.com/golang/mock/gomock"
)

func TestGetOrderByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMockOrderRepository(ctrl)

	t.Run("return no error while fetching order, with ID: 1", func(t *testing.T) {
		orderID := int64(1)
		var want *models.Order

		createOrderForTesting(t, orderID, want)

		mockRepo.EXPECT().GetByID(orderID).Return(want, nil)

		testService := NewOrderService(mockRepo)

		got, err := testService.GetOrderByID(orderID)
		if err != nil {
			t.Fatal("Got error while fetching order by ID", err)
		}

		if got != want {
			t.Errorf("Expected no error Got : %v", got)
		}
	})
}

func createOrderForTesting(t testing.TB, orderID int64, order *models.Order) {
	t.Helper()

	price := float32(14.55)
	var orderItems []models.OrderItem

	orderItem := models.OrderItem{
		ID:          orderID + int64(3),
		OrderID:     orderID,
		Description: "Test Product",
		Price:       price,
		Quantity:    1,
	}

	orderItems = append(orderItems, orderItem)

	orderModel := &models.Order{
		ID:           orderID,
		Status:       "PENDING",
		Total:        price,
		CurrencyUnit: "USD",
		OrderItem:    orderItems,
	}

	order = orderModel
}
