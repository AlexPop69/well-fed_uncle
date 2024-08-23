package service

import (
	"time"

	"github.com/AlexPop69/well-fed_uncle/models"
	"github.com/AlexPop69/well-fed_uncle/pkg/storage"
)

// OrderService предоставляет методы для работы с заказами.
type OrderService struct {
	storage *storage.Storage // Слой доступа к данным
}

// NewOrderService создает новый экземпляр OrderService.
func NewOrderService(storage *storage.Storage) *OrderService {
	return &OrderService{storage: storage}
}

// CreateOrder создает новый заказ для клиента.
func (s *OrderService) CreateOrder(clientID, pickupPointID int, items []models.MenuItem) (*models.Order, error) {
	order := &models.Order{
		ClientID:      clientID,
		PickupPointID: pickupPointID,
		Items:         items,
		CreatedAt:     time.Now(),
	}

	err := s.storage.CreateOrder(order)
	if err != nil {
		return nil, err
	}

	return order, nil
}

// GetOrder возвращает заказ по его идентификатору.
func (s *OrderService) GetOrder(orderID int) (*models.Order, error) {
	return s.storage.GetOrderById(orderID)
}

// ListOrders возвращает все заказы клиента.
func (s *OrderService) ListOrders(clientID int) ([]models.Order, error) {
	return s.storage.GetOrdersByClientId(clientID)
}
