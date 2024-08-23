package service

import (
	"errors"

	"github.com/AlexPop69/well-fed_uncle/models"
	"github.com/AlexPop69/well-fed_uncle/pkg/storage"
)

// ClientService предоставляет методы для работы с клиентами.
type ClientService struct {
	storage *storage.Storage // Слой доступа к данным
}

// NewClientService создает новый экземпляр ClientService.
func NewClientService(storage *storage.Storage) *ClientService {
	return &ClientService{storage: storage}
}

// StartInteraction проверяет, существует ли клиент, и создает его, если нет.
func (s *ClientService) StartInteraction(id int64, username string) (*models.Client, error) {
	// Проверяем, существует ли клиент с данным идентификатором чата.
	client, err := s.storage.GetClientById(id)
	if err != nil {
		if errors.Is(err, storage.ErrClientNotFound) {
			// Если клиент не найден, создаем нового.
			client = &models.Client{
				Name:  username,
				Phone: "", // Телефон можно запросить позже
			}
			err = s.storage.CreateClient(client)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	return client, nil
}

// GetClient возвращает клиента по идентификатору чата.
func (s *ClientService) GetClient(id int64) (*models.Client, error) {
	return s.storage.GetClientById(id)
}
