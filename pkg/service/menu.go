package service

import (
	"github.com/AlexPop69/well-fed_uncle/models"
	"github.com/AlexPop69/well-fed_uncle/pkg/storage"
)

// MenuService предоставляет методы для работы с меню.
type MenuService struct {
	storage *storage.Storage // Слой доступа к данным
}

// NewMenuService создает новый экземпляр MenuService.
func NewMenuService(storage *storage.Storage) *MenuService {
	return &MenuService{storage: storage}
}

// GetMenuItems возвращает все элементы меню.
func (s *MenuService) GetMenuItems() ([]models.MenuItem, error) {
	return s.storage.GetAllMenuItems()
}

// GetMenuItem возвращает элемент меню по его идентификатору.
func (s *MenuService) GetMenuItem(id int) (*models.MenuItem, error) {
	return s.storage.GetMenuItemById(id)
}

// AddMenuItem добавляет новый элемент меню.
func (s *MenuService) AddMenuItem(name string, price float64) (*models.MenuItem, error) {
	item := &models.MenuItem{
		Name:  name,
		Price: price,
	}

	err := s.storage.CreateMenuItem(item)
	if err != nil {
		return nil, err
	}

	return item, nil
}

// UpdateMenuItem обновляет существующий элемент меню.
func (s *MenuService) UpdateMenuItem(id int, name string, price float64) error {
	return s.storage.UpdateMenuItem(id, name, price)
}

// DeleteMenuItem удаляет элемент меню.
func (s *MenuService) DeleteMenuItem(id int) error {
	return s.storage.DeleteMenuItem(id)
}
