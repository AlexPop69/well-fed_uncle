package service

import (
	"errors"
	"time"

	"github.com/AlexPop69/well-fed_uncle/models"
	"github.com/AlexPop69/well-fed_uncle/pkg/storage"
)

// ShopService предоставляет методы для работы с пунктами самовывоза.
type ShopService struct {
	storage *storage.Storage // Слой доступа к данным
}

// NewShopService создает новый экземпляр ShopService.
func NewShopService(storage *storage.Storage) *ShopService {
	return &ShopService{storage: storage}
}

// CreateShop добавляет новый пункт самовывоза
func (s *ShopService) CreateShop(shop *models.Shop) error {
	return s.storage.CreateShop(shop)
}

// GetShop возвращает пункт самовывоза по его идентификатору.
func (s *ShopService) GetShopByName(shopName string) (*models.Shop, error) {
	return s.storage.GetShopByName(shopName)
}

// ListShops возвращает все пункты самовывоза.
func (s *ShopService) GetAllShops() ([]models.Shop, error) {
	return s.storage.GetAllShops()
}

// GetOpenShops возвращает все открытые пункты самовывоза.
func (s *ShopService) GetOpenShops() ([]models.Shop, error) {
	shops, _ := s.storage.GetAllShops()

	var result []models.Shop
	for _, shop := range shops {
		ok, _ := s.isOpen(shop.Name)
		if !ok {
			continue
		}

		result = append(result, shop)
	}

	return result, nil
}

// IsOpen проверяет, открыт ли пункт самовывоза в данный момент.
func (s *ShopService) isOpen(shopName string) (bool, error) {
	shop, err := s.storage.GetShopByName(shopName)
	if err != nil {
		return false, err
	}

	n := time.Now().Local()

	var close time.Time

	if shop.CloseTime.Hour() < shop.OpenTime.Hour() {
		close = time.Date(n.Year(), n.Month(), n.Day()+1, shop.CloseTime.Hour(), shop.CloseTime.Minute(), n.Second(), 0, time.Local)
	}

	open := time.Date(n.Year(), n.Month(), n.Day(), shop.OpenTime.Hour(), shop.OpenTime.Minute(), n.Second(), 0, time.Local)

	if n.Before(open) || n.After(close) {
		return false, errors.New("пункт самовывоза закрыт")
	}

	return true, nil
}
